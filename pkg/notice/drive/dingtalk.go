package drive

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
)

type DingTalkMod struct {
	DingTalkWebHook string
}

func DingTalkLoad(webHook string) *DingTalkMod {
	return &DingTalkMod{
		DingTalkWebHook: webHook,
	}
}

func (m DingTalkMod) Send(value string) {
	ctx := gctx.New()
	// 从配置中获取发送者名称
	name, _ := g.Cfg().Get(ctx, "name")

	// 定义钉钉机器人发送消息的URL，其中access_token为固定值
	url := m.DingTalkWebHook
	url += "&timestamp=" + gtime.Now().TimestampMilliStr()
	// 使用goroutine异步发送消息

	var post = g.Map{
		"msgtype": "text",
		"text": g.Map{
			"content": "通知姬 " + name.String() + "：\n" + value + "\n" + gtime.Now().String(),
		},
	}

	// 构建发送的消息体，包含消息类型和内容
	_, err := g.Client().ContentJson().Post(ctx, url, post)
	if err != nil {
		g.Log().Info(ctx, err)
	}

	return
}
