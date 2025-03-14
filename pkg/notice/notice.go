package notice

import (
	v1 "github.com/ayflying/utility_go/api/pgk/v1"
	"github.com/ayflying/utility_go/pkg/notice/drive"
)

type MessageV1 interface {
	Send(value string)
}

func New(typ v1.NoticeType, host string, value ...interface{}) MessageV1 {
	switch typ {
	case v1.NoticeType_DINGTALK:
		return drive.DingTalkLoad(host)
	case v1.NoticeType_EMAIL:
		return drive.MailLoad(host, value[0].(int), value[1].(string), value[2].(string))
	}
	return nil
}
