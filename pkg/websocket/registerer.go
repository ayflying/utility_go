package websocket

// 定义一个处理方法的类型
type Handler func(conn *WebsocketData, req any) (err error)
type OnConnectHandler func(conn *WebsocketData)

// 路由器的处理映射
var (
	handlers          = make(map[int]Handler)
	OnConnectHandlers = make([]OnConnectHandler, 0)
	OnCloseHandlers   = make([]OnConnectHandler, 0)
)

// 注册方法，将某个消息路由器ID和对应的处理方法关联起来
func (s *SocketV1) RegisterRouter(cmd int, handler Handler) {
	handlers[cmd] = handler
}

//注册方法，讲长连接登陆方法进行注册
func (s *SocketV1) RegisterOnConnect(_func OnConnectHandler) {
	OnConnectHandlers = append(OnConnectHandlers, _func)
}

func (s *SocketV1) RegisterOnClose(_func OnConnectHandler) {
	OnCloseHandlers = append(OnCloseHandlers, _func)
}
