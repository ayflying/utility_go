package websocket

import "context"

// 定义一个处理方法的类型
type Handler func(ctx context.Context, req any) (err error)

// 路由器的处理映射
var (
	handlers = make(map[int]Handler)
)

// 注册方法，将某个消息路由器ID和对应的处理方法关联起来
func (s *SocketV1) RegisterRouter(cmd int, handler Handler) {
	handlers[cmd] = handler
}
