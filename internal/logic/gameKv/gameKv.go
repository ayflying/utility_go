package gameKv

import (
	"github.com/ayflying/utility_go/service"
	"github.com/ayflying/utility_go/tools"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"strconv"
	"strings"
	"sync"
)

var (
	ctx  = gctx.New()
	Name = "game_kv"
)

type sGameKv struct {
	Lock sync.Mutex
}

func New() *sGameKv {
	return &sGameKv{}
}

func init() {
	service.RegisterGameKv(New())

	//支付钩子
	//task.Task.Trigger(tasks.TaskType_PAY, service.GameKv().HookPay)
	//task.Task.Trigger(tasks.TaskType_WARDROBE_LEVEL, service.GameKv().HookLevelRwd)
}

// SavesV1 方法
//
// @Description: 保存用户KV数据列表。
// @receiver s: sGameKv的实例。
// @return err: 错误信息，如果操作成功，则为nil。
func (s *sGameKv) SavesV1() (err error) {
	// 从Redis列表中获取所有用户KV索引的键
	//keys, err := utils.RedisScan("user:kv:*")
	err = tools.Redis.RedisScanV2("user:kv:*", func(keys []string) (err error) {
		// 定义用于存储用户数据的结构体
		type ListData struct {
			Uid int64       `json:"uid"`
			Kv  interface{} `json:"kv"`
		}
		var list []*ListData
		// 初始化列表，长度与keys数组一致
		list = make([]*ListData, 0)
		//需要删除的key
		var delKey []string
		// 遍历keys，获取每个用户的数据并填充到list中
		for _, cacheKey := range keys {
			//g.Log().Infof(ctx, "保存用户kv数据%v", v)
			//uid := v.Int64()
			//cacheKey = "user:kv:" + strconv.FormatInt(uid, 10)
			result := strings.Split(cacheKey, ":")
			var uid int64
			uid, err = strconv.ParseInt(result[2], 10, 64)
			if err != nil {
				g.Log().Error(ctx, err)
				g.Redis().Del(ctx, cacheKey)
				continue
			}

			////如果1天没有活跃，跳过
			//user, _ := service.MemberUser().Info(uid)
			//if user.UpdatedAt.Seconds < gtime.Now().Add(consts.ActSaveTime).Unix() {
			//	continue
			//}

			get, _ := g.Redis().Get(ctx, cacheKey)
			var data interface{}
			get.Scan(&data)
			list = append(list, &ListData{
				Uid: int64(uid),
				Kv:  data,
			})

			delKey = append(delKey, cacheKey)
		}

		// 将列表数据保存到数据库
		if len(list) > 0 {
			_, err2 := g.Model("game_kv").Batch(30).Data(list).Save()
			list = make([]*ListData, 0)
			if err2 != nil {
				g.Log().Error(ctx, err2)
				return
			}

			//批量删除key
			for _, v := range delKey {
				_, err2 = g.Redis().Del(ctx, v)
				if err2 != nil {
					g.Log().Errorf(ctx, "删除存档错误：%v,err=%v", v, err2)
					return
				}
			}

			delKey = make([]string, 0)

		}
		if err != nil {
			g.Log().Error(ctx, "当前kv数据入库失败: %v", err)
		}

		return
	})

	//if err != nil {
	//	return err
	//}
	////跳过
	//if len(keys) == 0 {
	//	return
	//}
	////一次最多处理10w条
	//if len(keys) > 10000 {
	//	keys = keys[:10000]
	//}

	return
}
