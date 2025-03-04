package message

import "github.com/ayflying/utility_go/package/message/drive"

type MessageV1 interface {
	Send(value string)
}

func New(typ MessageType, host string) MessageV1 {
	switch typ {
	case DingTalk:
		return drive.Load(host)

	}
	return nil
}
