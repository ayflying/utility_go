package websocket

import (
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"strconv"
	"sync"
)

type SocketV1 struct {
	Type int `json:"type"`
}

var (
	//ctx = gctx.New()
	//Conn map[uuid.UUID]*WebsocketData
	lock sync.Mutex

	m = gmap.New(true)
)

type WebsocketData struct {
	Ws   *websocket.Conn
	Uuid uuid.UUID
	Uid  int64
	Ctx  context.Context
}

func NewV1() *SocketV1 {
	return &SocketV1{
		Type: 1,
	}
}

type SocketInterface interface {
	OnConnect(*websocket.Conn)
	OnMessage(*WebsocketData, []byte, int)
	Send(uuid.UUID, []byte) (err error)
	SendAll(data []byte)
	OnClose(conn *websocket.Conn)
}

func (s *SocketV1) Load(serv *ghttp.Server, prefix string) {
	//websocket服务启动
	serv.Group(prefix, func(group *ghttp.RouterGroup) {

		var websocketCfg = websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}
		group.Bind(
			func(r *ghttp.Request) {
				ctx := r.Context()
				ws, err := websocketCfg.Upgrade(r.Response.Writer, r.Request, nil)
				if err != nil {
					glog.Error(ctx, err)
					r.Exit()
				}

				//ws联机触发器
				NewV1().OnConnect(ctx, ws)
			},
		)

	})
}

// OnConnect
//
//	@Description:
//	@receiver s
//	@param conn
func (s *SocketV1) OnConnect(ctx context.Context, conn *websocket.Conn) {
	//lock.Lock()
	//defer lock.Unlock()

	defer conn.Close()
	id, _ := uuid.NewUUID()
	ip := conn.RemoteAddr().String()
	data := &WebsocketData{
		Uuid: id,
		Ws:   conn,
		Ctx:  ctx,
	}
	m.Set(id, data)

	//defer delete(Conn, id)

	//to := fmt.Sprintf("创建连接：%v,ip=%v", id, ip)
	//s.Send(id, []byte(to))

	//用户登录钩子执行
	for _, connect := range OnConnectHandlers {
		connect(data)
	}

	for {
		//进入当前连接线程拥堵
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			//客户端断开返回错误，断开当前连接
			break
		}
		s.OnMessage(m.Get(id).(*WebsocketData), msg, msgType)
	}
	//关闭连接触发
	s.OnClose(data)
	g.Log().Debugf(ctx, "断开连接:uuid=%v,ip=%v", id, ip)
}

// OnMessage
//
//	@Description:
//	@receiver s
//	@param msg
//	@param msgType
func (s *SocketV1) OnMessage(conn *WebsocketData, req []byte, msgType int) {
	s.Type = msgType
	//g.Log().Debugf(ctx, "收到消息：%v,type=%v,conn=%v", string(req), msgType, conn)
	//s.Send(conn.Uuid, msg)
	//s.SendAll(msg)
	//msgStr := string(req)
	msg := req[8:]
	cmd := gconv.Int(req[:8])
	//GetRouter(cmd, conn.Uid, msg)
	handler, exist := handlers[cmd]
	if exist {
		//匹配上路由器
		err := handler(conn, msg)
		g.Log().Error(conn.Ctx, err)
	} else {
		//fmt.Println("未注册的路由器ID:", cmd)
		s.Send(conn.Uuid, []byte("未注册的协议号:"+strconv.Itoa(cmd)))
		s.OnClose(conn)
		return
	}

}

// Send
//
//	@Description:
//	@receiver s
//	@param uid
//	@param data
//	@return err
func (s *SocketV1) Send(id uuid.UUID, data []byte) (err error) {
	if !m.Contains(id) {
		return
	}

	conn := m.Get(id).(*WebsocketData)
	conn.Ws.WriteMessage(s.Type, data)

	return
}

// 批量发送
func (s *SocketV1) SendAll(data []byte) {
	m.Iterator(func(k interface{}, v interface{}) bool {
		//fmt.Printf("%v:%v ", k, v)
		conn := v.(*WebsocketData)
		conn.Ws.WriteMessage(s.Type, data)

		return true
	})
}

// OnClose
//
//	@Description:
//	@receiver s
//	@param conn
func (s *SocketV1) OnClose(conn *WebsocketData) {
	// 在此处编写断开连接后的处理逻辑
	//g.Log().Debugf(gctx.New(), "WebSocket connection from %s has been closed.", conn.RemoteAddr())

	//用户登录钩子执行
	for _, connect := range OnCloseHandlers {
		connect(conn)
	}

	// 可能的后续操作：
	// 1. 更新连接状态或从连接池移除
	// 2. 发送通知或清理关联资源
	// 3. 执行特定于业务的断开处理
	m.Remove(conn.Uuid)
	conn.Ws.Close()
}
