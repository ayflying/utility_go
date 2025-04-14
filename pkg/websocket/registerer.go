package websocket

import "google.golang.org/protobuf/proto"

// 定义一个处理方法的类型
type Handler func(conn *WebsocketData, req any) (err error)
type Handler2 func(conn *WebsocketData)

type HandlerMessage func(conn *WebsocketData, req any)

type PbType func(cmd int32, data []byte, code int32, msg string) proto.Message
type PbType2 func(data []byte) (int, []byte)

// 路由器的处理映射
var (
	handlers          = make(map[int]Handler)
	OnConnectHandlers = make([]Handler2, 0)
	OnCloseHandlers   = make([]Handler2, 0)
	onMessageHandlers = make([]HandlerMessage, 0)
	Byte2Pb           = make([]PbType, 0)
	Pb2Bytes          = make([]PbType2, 0)
)

// 注册方法，将某个消息路由器ID和对应的处理方法关联起来
func (s *SocketV1) RegisterRouter(cmd int, handler Handler) {
	handlers[cmd] = handler
}

//注册方法，讲长连接登陆方法进行注册
func (s *SocketV1) RegisterOnConnect(_func Handler2) {
	OnConnectHandlers = append(OnConnectHandlers, _func)
}

func (s *SocketV1) RegisterOnClose(_func Handler2) {
	OnCloseHandlers = append(OnCloseHandlers, _func)
}

//注册方法长连接消息体
func (s *SocketV1) RegisterMessage(_func HandlerMessage) {
	onMessageHandlers = append(onMessageHandlers, _func)
}

func (s *SocketV1) RegisterByte2Pb(_func PbType) {
	Byte2Pb = append(Byte2Pb, _func)
}

func (s *SocketV1) RegisterPb2Byte(_func PbType2) {
	Pb2Bytes = append(Pb2Bytes, _func)
}
