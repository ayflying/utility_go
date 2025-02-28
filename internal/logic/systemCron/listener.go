package systemCron

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gtime"
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Status int `json:"status"`
	} `json:"data"`
}

func (s *sSystemCron) Guardian(DingTalkWebHook string) {
	var list []struct {
		Name    string
		Address string
	}
	cfg, _ := g.Cfg().Get(ctx, "serverList")
	cfg.Scan(&list)
	for _, v := range list {
		get, err := g.Client().Discovery(nil).Get(ctx, v.Address+"/callback/status")

		defer get.Close()
		if err != nil {
			s.DingTalk(DingTalkWebHook, fmt.Sprintf("监控报警：服务端访问失败 (%v 服务器),err=%v", v.Name, err))
		} else if get.StatusCode != 200 {
			s.DingTalk(DingTalkWebHook, fmt.Sprintf("监控报警：服务端访问失败 (%v 服务器),code=%v,err=%v", v.Name, get.StatusCode, err))
		} else {
			var ststus Status
			err = json.Unmarshal(get.ReadAll(), &ststus)
			if ststus.Code != 0 {
				s.DingTalk(DingTalkWebHook, fmt.Sprintf("监控报警：服务端访问失败 (%v 服务器),msg=%v", v.Name, ststus.Message))
			}
		}
	}
}

// DingTalk 发送钉钉消息
//
// @Description: 向指定的钉钉机器人发送消息。
// @receiver s: 系统定时任务结构体指针。
// @param value: 要发送的消息内容。
func (s *sSystemCron) DingTalk(DingTalkWebHook string, value string) (res *gclient.Response) {
	// 从配置中获取发送者名称
	name, _ := g.Cfg().Get(ctx, "name")

	// 定义钉钉机器人发送消息的URL，其中access_token为固定值
	url := DingTalkWebHook
	url += "&timestamp=" + gtime.Now().TimestampMilliStr()
	// 使用goroutine异步发送消息

	var post = g.Map{
		"msgtype": "text",
		"text": g.Map{
			"content": "通知姬 " + name.String() + "：\n" + value + "\n" + gtime.Now().String(),
		},
	}

	// 构建发送的消息体，包含消息类型和内容
	res, err := g.Client().Discovery(nil).ContentJson().Post(ctx, url, post)
	if err != nil {
		g.Log().Info(ctx, err)
	}
	return
}
