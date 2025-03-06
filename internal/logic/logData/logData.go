package logData

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"new-gitlab.adesk.com/public_project/utility_go/service"
	"new-gitlab.adesk.com/public_project/utility_go/tools"
	"time"
)

var (
	ctx = gctx.New()
	//pathStr    = "runtime/log/logData"
	//te         thinkingdata.TDAnalytics
	//logChannel chan map[string]interface{}
	//wg         sync.WaitGroup
)

type sLogData struct {
}

func New() *sLogData {
	return &sLogData{}
}

func init() {
	service.RegisterLogData(New())

	//加载日志模块
	//service.LogData().Load()

}

func (s *sLogData) Load() {
	//数数科技初始化配置

	//// 创建 LogConfig 配置文件
	//config := thinkingdata.TDLogConsumerConfig{
	//	Directory: pathStr, // 事件采集的文件路径
	//	//FileSize:  99,      //单个日志文件的最大大小（MB）
	//}
	//// 初始化 logConsumer
	//consumer, _ := thinkingdata.NewLogConsumerWithConfig(config)
	//// 创建 te 对象
	//te = thinkingdata.New(consumer)

	//日志写入通道开启
	//logAppend()
}

// UserSet 方法
//
// @Description: 设置用户信息。
// @receiver s: sLogData 的实例，表示日志数据的结构体。
// @param accountId: 账户ID，用于标识账户，是字符串格式。
// @param uid: 用户的ID，是整型的唯一标识符。
// @param data: 要设置的用户信息，以键值对的形式提供，是map[string]interface{}类型，支持多种用户属性。
// @return err: 执行过程中可能出现的错误，如果执行成功，则返回nil。
func (s *sLogData) UserSet(accountId string, uid int64, data map[string]interface{}) (err error) {
	// 将用户ID转换为字符串格式的唯一标识
	//distinctId := strconv.FormatInt(uid, 10)
	// 使用accountId和distinctId以及data来设置用户信息，此处调用外部方法完成设置。
	//te.UserSet(accountId, distinctId, data)
	data["#uid"] = uid
	data["#time"] = time.Now()
	data["#type"] = "user_set"
	//data["_id"], _ = uuid.NewUUID()
	//data["#name"] = name
	g.Log("elk").Info(nil, data)

	//todo 暂时关闭update
	//err = s.Update(uid, data)
	return
}

// Track 函数记录特定事件。
//
// @Description: 用于跟踪和记录一个指定事件的发生，收集相关数据。
// @receiver s: sLogData 的实例，代表日志数据的存储或处理实体。
// @param accountId: 账户ID，用于标识事件所属的账户。
// @param uid: 用户的ID，一个整型数值，用于区分不同的用户。
// @param name: 事件名称，标识所记录的具体事件。
// @param data: 事件相关的数据映射，包含事件的详细信息。
// @return err: 错误信息，如果操作成功则为nil。
func (s *sLogData) Track(ctx context.Context, accountId string, uid int64, name string, data map[string]interface{}) {
	// 将用户ID转换为字符串格式的唯一标识
	//distinctId := strconv.FormatInt(uid, 10)
	// 调用te.Track函数来实际记录事件，传入账户ID、用户唯一标识、事件名称及事件数据
	//te.Track(accountId, distinctId, name, data)
	if data == nil {
		return
	}

	data["#uid"] = uid
	data["#event_name"] = name
	data["#time"] = time.Now()
	data["#type"] = "track"
	//data["_id"], _ = uuid.NewUUID()
	//道具类型特殊格式化
	if get, ok := data["items"]; ok {
		if get != nil {
			data["items"] = tools.Tools.Items2Map(get.([][]int64))
		}
	}

	g.Log("elk").Info(nil, data)
	//err = s.Add(data)
	//由于实时写入日志太占用资源，关闭日志写入方法
	return
}

//// 上报用不了，弃用
//func (s *sLogData) Send() (err error) {
//	consumer, err := thinkingdata.NewBatchConsumer("https://yoyatime-server-release.yoyaworld.com/callback", "dev")
//	te = thinkingdata.New(consumer)
//	err = te.Flush()
//	return
//}

//func (s *sLogData) Flush() (err error) {
//	//调用flush接口数据会立即写入文件，生产环境注意避免频繁调用flush引发IO或网络开销问题
//	err = te.Flush()
//	return
//}
//
//func (s *sLogData) Close() {
//	// 关闭通道，表示没有更多的日志条目需要写入
//	if logChannel != nil {
//		close(logChannel)
//		wg.Wait() // 等待通道监听goroutine结束
//	}
//
//	te.Close()
//}
