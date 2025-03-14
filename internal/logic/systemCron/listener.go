package systemCron

import (
	"encoding/json"
	"fmt"
	v1 "github.com/ayflying/utility_go/api/pgk/v1"
	"github.com/ayflying/utility_go/pkg/notice"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
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
			notice.New(v1.NoticeType_DINGTALK, DingTalkWebHook).Send(fmt.Sprintf("监控报警：服务端访问失败 (%v 服务器),err=%v", v.Name, err))
		} else if get.StatusCode != 200 {
			notice.New(v1.NoticeType_DINGTALK, DingTalkWebHook).Send(fmt.Sprintf("监控报警：服务端访问失败 (%v 服务器),code=%v,err=%v", v.Name, get.StatusCode, err))
		} else {
			var ststus Status
			err = json.Unmarshal(get.ReadAll(), &ststus)
			if ststus.Code != 0 {
				notice.New(v1.NoticeType_DINGTALK, DingTalkWebHook).Send(fmt.Sprintf("监控报警：服务端访问失败 (%v 服务器),msg=%v", v.Name, ststus.Message))
			}
		}
	}
}

// DingTalk 发送钉钉消息
//
// @Description: 向指定的钉钉机器人发送消息。
// @receiver s: 系统定时任务结构体指针。
// @param value: 要发送的消息内容。
// Deprecated: Use message.New(message.DingTalk, DingTalkWebHook).Send(value)
func (s *sSystemCron) DingTalk(DingTalkWebHook string, value string) (res *gclient.Response) {
	notice.New(v1.NoticeType_DINGTALK, DingTalkWebHook).Send(value)
	return
}
