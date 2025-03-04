package message

type MessageType int32

const (
	DingTalk MessageType = iota
	Wechat
	Email
	Sms
	Voice
)
