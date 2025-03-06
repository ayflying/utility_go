package notice

import (
	v1 "new-gitlab.adesk.com/public_project/utility_go/api/pgk/v1"
	"new-gitlab.adesk.com/public_project/utility_go/pgk/notice/drive"
)

type MessageV1 interface {
	Send(value string)
}

func New(typ v1.NoticeType, host string) MessageV1 {
	switch typ {
	case v1.NoticeType_DINGTALK:
		return drive.Load(host)

	}
	return nil
}
