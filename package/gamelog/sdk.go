package gamelog

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sendBody struct {
	Pid      string  `json:"pid"`
	Data     [][]any `json:"data"`
	SaveType int     `json:"save_type" dc:"0=文件存储, 1=kafka存储"`
}

// todo 游戏日志对象
type GameLog struct {
	Uid          string         // 唯一uid
	Event        string         // 事件名
	Property     map[string]any // 事件属性
	EventTimems  int64          // 时间戳毫秒级别
	EventTimeLoc string         // 带时区的本地时间字符串
}

type SDKConfig struct {
	// 配置变量
	Pid           string // 项目id
	BaseUrl       string // 日志服务器地址
	ReportSk      string // 上报解密key
	FlushInterval int    // 刷新间隔
	DiskBakPath   string // 磁盘备份路径
	RetryN        int    // 每N次重试
	ChanSize      int    // 信道大小, 默认1000

	reportN      int
	SendSaveType int // 发送存储类型, 默认不设置为0代表文件存储, 1代表走kafka可实同步日志
}

type SDK struct {
	// 控制变量
	wg         sync.WaitGroup
	shutdown   chan struct{}
	mu         sync.Mutex
	sdkConfig  *SDKConfig
	bufferChan chan GameLog // 日志队列
	buffer     []GameLog    // 日志队列
}

var (
	ctx           = context.Background()
	gamelogClient *gclient.Client

	// location map
	// locationMap map[string]*time.Location = map[string]*time.Location{}
	locationMap sync.Map // 声明一个线程安全的Map

)

var safePropertyRE = regexp.MustCompile(
	`["'\\\/]` +
		`|\\U[0-9a-fA-F]{8}` + // Unicode 8位转义
		`|\\u[0-9a-fA-F]{4}` + // Unicode 4位转义
		`|\\[tnrfbv\\]` + // 转义字面量
		`|[\x00-\x1F\x7F-\x9F]` + // ASCII + C1 控制字符
		`|[\\u200B-\\u200D\\uFEFF]`, // 零宽字符
)

func safeProperty(property map[string]any) {
	for k, v := range property {
		if _, ok := v.(string); ok {
			property[k] = safePropertyRE.ReplaceAllString(gconv.String(v), "*")
		}
	}
}

func getLocationMapValue(key string) *time.Location {
	// 1. 先尝试读
	value, loaded := locationMap.Load(key)
	if loaded {
		return value.(*time.Location) // 如果已经存在，直接返回
	}
	// 2. 不存在，就初始化一个该key对应的**固定的**新值
	location, err := time.LoadLocation(key)
	if err != nil {
		g.Log().Warningf(ctx, "[GameLog]load location error, try use local timezone: %v", err)
		return nil
	}
	// 3. 核心：原子性地存储，如果key已存在则返回已存在的值
	actualValue, loaded := locationMap.LoadOrStore(key, location)
	if loaded {
		// 如果loaded为true，说明其他goroutine抢先存了
		// 我们可以丢弃刚创建的newValue（如果有需要的话），返回已存在的actualValue
		return actualValue.(*time.Location)
	}
	// 如果loaded为false，说明是我们存成功的，返回我们刚创建的newValue
	return actualValue.(*time.Location)
}

func (sdk *SDK) varinit() error {
	sdk.sdkConfig = &SDKConfig{}

	_pid, err := g.Config().Get(ctx, "angergs.bisdk.pid")
	if err != nil {
		return err
	}
	sdk.sdkConfig.Pid = _pid.String()

	_baseUrl, err := g.Config().Get(ctx, "angergs.bisdk.recodeServerBaseUrl")
	if err != nil {
		return err
	}
	sdk.sdkConfig.BaseUrl = _baseUrl.String()

	_sk, err := g.Config().Get(ctx, "angergs.bisdk.reportSk")
	if err != nil {
		return err
	}
	sdk.sdkConfig.ReportSk = _sk.String()

	_flushInterval, err := g.Config().Get(ctx, "angergs.bisdk.flushInterval")
	if err != nil {
		return err
	}
	sdk.sdkConfig.FlushInterval = _flushInterval.Int()

	_diskBakPath, err := g.Config().Get(ctx, "angergs.bisdk.diskBakPath")
	if err != nil {
		return err
	}
	sdk.sdkConfig.DiskBakPath = _diskBakPath.String()

	_retryN, err := g.Config().Get(ctx, "angergs.bisdk.retryN")
	if err != nil {
		return err
	}
	sdk.sdkConfig.RetryN = _retryN.Int()

	_chanSize, err := g.Config().Get(ctx, "angergs.bisdk.chanSize")
	if err != nil {
		return err
	}
	sdk.sdkConfig.ChanSize = _chanSize.Int()

	g.Log().Infof(ctx, "[GameLog]client init success, config: %v", sdk.sdkConfig)
	return nil
}

func (sdk *SDK) checkConfig() error {
	config := sdk.sdkConfig
	if config.Pid == "" {
		return fmt.Errorf("pid is empty")
	}
	if config.BaseUrl == "" {
		return fmt.Errorf("baseUrl is empty")
	}
	if config.ReportSk == "" {
		return fmt.Errorf("reportSk is empty")
	}
	if config.FlushInterval <= 0 {
		return fmt.Errorf("flushInterval is invalid")
	}
	if config.DiskBakPath == "" {
		return fmt.Errorf("diskBakPath is empty")
	}
	if config.RetryN == 0 {
		config.RetryN = 10
	}
	if config.ChanSize == 0 {
		config.ChanSize = 1000
	}
	config.DiskBakPath = strings.TrimSuffix(config.DiskBakPath, "/")

	return nil
}

func INIT(config *SDKConfig) (*SDK, error) {
	// 加载并检查配置
	sdk := &SDK{}
	if config != nil {
		sdk.sdkConfig = config
	} else if err := sdk.varinit(); err != nil { // 可以读goframe的配置
		return nil, err
	}
	if err := sdk.checkConfig(); err != nil {
		return nil, err
	}
	gamelogClient = g.Client()

	// 初始化队列
	sdk.shutdown = make(chan struct{})
	sdk.bufferChan = make(chan GameLog, 1000)
	sdk.buffer = make([]GameLog, 0, 100)
	// 加载失败日志
	failLogs, err := sdk.loadFailLogs4disk()
	if err != nil {
		g.Log().Errorf(ctx, "[GameLog]load fail logs error: %v", err)
	} else if len(failLogs) > 0 {
		sdk.buffer = append(sdk.buffer, failLogs...)
	}

	// 开启协程进行日志发送
	sdk.wg = sync.WaitGroup{}
	sdk.wg.Add(1)
	go func() {
		defer sdk.wg.Done()
		ticker := time.NewTicker(time.Duration(sdk.sdkConfig.FlushInterval) * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-sdk.shutdown:
				// 关闭时, 上传一次并备份失败数据
				g.Log().Infof(ctx, "[GameLog]begin shutdown and flush last")
				sdk.flush()
				return
			case log := <-sdk.bufferChan:
				sdk.buffer = append(sdk.buffer, log)
			case <-ticker.C:
				sdk.flush()

			}
		}
	}()
	return sdk, nil
}

// 从磁盘加载失败日志
func (sdk *SDK) loadFailLogs4disk() (logs []GameLog, err error) {
	if !gfile.Exists(sdk.sdkConfig.DiskBakPath) {
		return
	}
	// 遍历diskBakPath下所有failBufferxxx.bak.log文件， 读取到log中
	files, err := gfile.ScanDir(sdk.sdkConfig.DiskBakPath, "failBuffer*.bak.log")
	logs = []GameLog{}
	if err != nil {
		return
	}
	// 读取每个备份文件
	for _, fp := range files {
		// 每一行都是一次失败的记录
		gfile.ReadLines(fp, func(line string) error {
			_logs := []GameLog{}
			err := json.Unmarshal([]byte(line), &_logs)
			if err != nil {
				return err
			}
			// 合并到总日志列表
			logs = append(logs, _logs...)
			return nil
		})
		g.Log().Infof(ctx, "[GameLog]load %d faillogs from %s", len(logs), fp)
		gfile.Remove(fp)
	}
	return

}

// 备份失败日志追加到磁盘
func (sdk *SDK) bakFailLogs2disk(failLogs []GameLog) {
	bakPath := fmt.Sprintf("%s/failBuffer%s.bak.log", sdk.sdkConfig.DiskBakPath, gtime.Now().Format("YmdH"))
	content, err := json.Marshal(failLogs)
	if err != nil {
		g.Log().Errorf(ctx, "[GameLog]marshal fail logs error: %v", err)
		return
	}
	gfile.PutContentsAppend(bakPath, string(content)+"\n")
	g.Log().Infof(ctx, "[GameLog]backup fail buffer to %s", bakPath)
}

// 优雅关闭
func (sdk *SDK) Shutdown() {
	close(sdk.shutdown)
	sdk.wg.Wait()
}

// 日志时间格式
const datetimeFmt = time.DateOnly + " " + time.TimeOnly

// 记录日志
func (sdk *SDK) Log(uid, event string, property map[string]any, timezone string, customEventTime ...time.Time) {
	loc := time.Local
	if _loc := getLocationMapValue(timezone); _loc != nil {
		loc = _loc
	}
	if len(property) == 0 {
		property = map[string]any{"ts": gtime.Now().Timestamp()}
	}
	safeProperty(property)
	var et *gtime.Time
	if len(customEventTime) > 0 {
		et = gtime.NewFromTime(customEventTime[0])
	} else {
		et = gtime.Now()
	}
	log := GameLog{
		Uid:          uid,
		Event:        event,
		Property:     property,
		EventTimems:  et.TimestampMilli(),
		EventTimeLoc: et.In(loc).Format(datetimeFmt),
	}
	// 线程安全
	sdk.bufferChan <- log
}

// 按服务器时区记录日志
func (sdk *SDK) LogLtz(uid, event string, property map[string]any, customEventTime ...time.Time) {
	sdk.Log(uid, event, property, time.Local.String(), customEventTime...)
}

// 用户属性初始化
func (sdk *SDK) Uinit(uid string, property map[string]any, timezone string, customEventTime ...time.Time) {
	sdk.Log(uid, "u_init", property, timezone, customEventTime...)
}
func (sdk *SDK) UinitLtz(uid string, property map[string]any, customEventTime ...time.Time) {
	sdk.Uinit(uid, property, time.Local.String(), customEventTime...)
}

// 用户属性设置
func (sdk *SDK) Uset(uid string, property map[string]any, timezone string, customEventTime ...time.Time) {
	sdk.Log(uid, "u_set", property, timezone, customEventTime...)
}
func (sdk *SDK) UsetLtz(uid string, property map[string]any, customEventTime ...time.Time) {
	sdk.Uset(uid, property, time.Local.String(), customEventTime...)
}

// 用户属性删除
func (sdk *SDK) Uunset(uid string, property map[string]any, timezone string, customEventTime ...time.Time) {
	sdk.Log(uid, "u_unset", property, timezone, customEventTime...)
}
func (sdk *SDK) UunsetLtz(uid string, property map[string]any, customEventTime ...time.Time) {
	sdk.Uunset(uid, property, time.Local.String(), customEventTime...)
}

// 用户属性自增/减
func (sdk *SDK) Uinc(uid string, property map[string]any, timezone string, customEventTime ...time.Time) {
	sdk.Log(uid, "u_inc", property, timezone, customEventTime...)
}
func (sdk *SDK) UincLtz(uid string, property map[string]any, customEventTime ...time.Time) {
	sdk.Uinc(uid, property, time.Local.String(), customEventTime...)
}

// 这个方法只会在内部协程调用
func (sdk *SDK) flush() {
	sdk.mu.Lock()
	defer sdk.mu.Unlock()
	if len(sdk.buffer) == 0 {
		return
	}

	batch := make([]GameLog, len(sdk.buffer))
	copy(batch, sdk.buffer)
	sdk.buffer = sdk.buffer[:0]

	// 第N次的时候加载失败数据进行尝试
	if sdk.sdkConfig.reportN != 0 && sdk.sdkConfig.reportN%sdk.sdkConfig.RetryN == 0 {
		faillogs, err := sdk.loadFailLogs4disk()
		if err != nil {
			g.Log().Errorf(ctx, "[GameLog]load fail logs error: %v", err)
		}
		// 如果有失败日志则加入到批量数组中
		if len(faillogs) > 0 {
			batch = append(batch, faillogs...)
		}
	}
	sdk.send(batch)
}

// 发送消息
func (sdk *SDK) send(logs []GameLog) {
	waitSecond := time.Duration(sdk.sdkConfig.FlushInterval/4) * time.Second
	timeoutCtx, cancel := context.WithTimeout(context.Background(), waitSecond)
	defer cancel()
	data := make([][]any, 0, len(logs))
	// logs 拆分成二维数组
	for _, log := range logs {
		propertyJson, err := json.Marshal(log.Property)
		if err != nil {
			g.Log().Errorf(ctx, "[GameLog]skip log parse, marshal property error: %v", err)
			continue
		}
		data = append(data, []any{
			log.Uid,
			log.Event,
			string(propertyJson),
			log.EventTimems,
			log.EventTimeLoc,
		})
	}
	// json化
	sbody := sendBody{
		Pid:      sdk.sdkConfig.Pid,
		Data:     data,
		SaveType: sdk.sdkConfig.SendSaveType,
	}
	jsonBody, err := json.Marshal(sbody)
	if err != nil {
		g.Log().Errorf(ctx, "[GameLog]marshal send body error: %v", err)
		return
	}

	// giz压缩
	gzBody := bytes.NewBuffer([]byte{})
	gz := gzip.NewWriter(gzBody)
	gz.Write(jsonBody)
	gz.Close()

	// XOR 加密
	xorBody := bytesXOR(gzBody.Bytes(), []byte(sdk.sdkConfig.ReportSk))

	sdk.sdkConfig.reportN += 1
	res, err := gamelogClient.Post(timeoutCtx, sdk.sdkConfig.BaseUrl+"/report/event", xorBody)
	// 失败重新加入缓冲区
	if err != nil {
		sdk.bakFailLogs2disk(logs)
		g.Log().Warningf(ctx, "[GameLog]send log error, bak to fail buffer(%d): %v", len(logs), err)
		return
	}
	defer func() {
		cerr := res.Close()
		if cerr != nil {
			g.Log().Errorf(ctx, "[GameLog]close response error: %v", cerr)
		}
	}()
	httpcode := res.StatusCode
	resBody := res.ReadAllString()
	// 收集器拦截, 重新加入缓冲区
	if httpcode != http.StatusOK {
		sdk.bakFailLogs2disk(logs)
		g.Log().Warningf(ctx, "[GameLog]send log error, bak to fail buffer(%d): %v", len(logs), resBody)
	}
}

// 混淆
func bytesXOR(data []byte, key []byte) []byte {
	obfuscated := make([]byte, len(data))
	keyLen := len(key)
	if keyLen == 0 {
		return data
	}

	for i := range data {
		obfuscated[i] = data[i] ^ key[i%keyLen]
	}
	return obfuscated

	// // 使用示例
	// key := []byte{0x12, 0x34, 0x56, 0x78}
	// obfuscated := multiXorObfuscate(original, key)
	// deobfuscated := multiXorObfuscate(obfuscated, key) // 解密
}
