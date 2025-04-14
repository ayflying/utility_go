package websocket

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"sync"
)

type SocketV1 struct {
	Type int `json:"type"`
}

var (
	//ctx = gctx.New()
	//Conn map[uuid.UUID]*WebsocketData
	lock sync.Mutex

	m = gmap.NewHashMap(true)
)

type WebsocketData struct {
	Ws     *websocket.Conn
	Uuid   string
	Uid    int64
	Ctx    context.Context
	RoomId int
}

func NewV1() *SocketV1 {
	return &SocketV1{
		Type: 2,
	}
}

type SocketInterface interface {
	Load(serv *ghttp.Server, prefix string)
	OnConnect(ctx context.Context, conn *websocket.Conn)
	OnMessage(conn *WebsocketData, req []byte, msgType int)
	Send(cmd int32, uid int64, req proto.Message)
	SendAll(cmd int32, req proto.Message)
	OnClose(conn *WebsocketData)
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

	defer conn.Close()
	id := guid.S()
	ip := conn.RemoteAddr().String()
	data := &WebsocketData{
		Uuid: id,
		Ws:   conn,
		Ctx:  ctx,
	}
	m.Set(id, data)

	//defer delete(Conn, id)

	to := fmt.Sprintf("创建连接：%v,ip=%v", id, ip)
	g.Log().Debugf(ctx, to)
	//s.Send(id, []byte(to))

	//用户登录钩子执行
	for _, connect := range OnConnectHandlers {
		connect(data)
	}

	for {
		//进入当前连接线程拥堵
		msgType, msg, err := conn.ReadMessage()
		s.Type = msgType
		if err != nil {
			//客户端断开返回错误，断开当前连接
			//g.Log().Error(ctx, err)
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
	s.Type = 2
	var cmd int
	var msg []byte
	//uid := conn.Uid
	for _, v := range Pb2Bytes {
		cmd, msg = v(req)
	}
	g.Log("cmd").Debugf(gctx.New(), fmt.Sprintf("from|%d|%d|%v", cmd, conn.Uid, gjson.MustEncodeString(req)))

	//msgStr := string(req)
	//cmd = gconv.Int(msgStr[:8])
	//msg = []byte(msgStr[8:])

	handler, exist := handlers[cmd]
	if exist {
		//匹配上路由器
		err := handler(conn, msg)
		if err != nil {
			g.Log().Error(conn.Ctx, err)
		}
	} else {
		//fmt.Println("未注册的路由器ID:", cmd)
		//s.Send(20000000, conn.Uid, []byte("未注册的协议号:"+strconv.Itoa(cmd)))
		s.OnClose(conn)
		return
	}

}

//绑定用户编号
func (s *SocketV1) BindUid(conn *WebsocketData, uid int64) {
	lock.Lock()
	defer lock.Unlock()

	cacheKey := fmt.Sprintf("socket:uid:%d", uid)
	g.Redis().Set(nil, cacheKey, conn.Uuid)

	if conn.Uid == 0 {
		conn.Uid = uid
	}

}

//解绑用户
func (s *SocketV1) UnBindUid(uid int64) {
	lock.Lock()
	defer lock.Unlock()

	cacheKey := fmt.Sprintf("socket:uid:%d", uid)
	g.Redis().Del(nil, cacheKey)
}

// Uid2Uuid 用户编号转uuid唯一标识
func (s *SocketV1) Uid2Uuid(uid int64) (uuid string) {
	cacheKey := fmt.Sprintf("socket:uid:%d", uid)
	get, _ := g.Redis().Get(nil, cacheKey)
	if get.IsNil() {
		return
	}

	uuid = get.String()

	//如果不在线了
	if !m.Contains(uuid) {
		// 解绑用户编号
		s.UnBindUid(uid)
		return
	}

	return
}

// SendUuid
//
//	@Description:
//	@receiver s
//	@param uid
//	@param data
func (s *SocketV1) SendUuid(cmd int32, id uuid.UUID, req proto.Message) {
	if !m.Contains(id) {
		return
	}
	var data, err = proto.Marshal(req)
	if err != nil {
		g.Log().Error(gctx.New(), err)
		return
	}

	conn := m.Get(id).(*WebsocketData)

	//前置方法

	for _, v := range Byte2Pb {
		temp := v(cmd, data, 0, "")
		data, _ = proto.Marshal(temp)
	}

	conn.Ws.WriteMessage(s.Type, data)

	return
}

// Send
//
//	@Description:
//	@receiver s
//	@param uid
//	@param data
func (s *SocketV1) Send(cmd int32, uid int64, req proto.Message) {
	g.Log("cmd").Debugf(gctx.New(), fmt.Sprintf("to|%d|%d|%v", cmd, uid, gjson.MustEncodeString(req)))

	uuid := s.Uid2Uuid(uid)
	if uuid == "" {
		return
	}
	if !m.Contains(uuid) {
		return
	}

	//格式化数据
	var data, err = proto.Marshal(req)
	if err != nil {
		g.Log().Error(gctx.New(), err)
		return
	}

	conn := m.Get(uuid).(*WebsocketData)

	//前置方法

	for _, v := range Byte2Pb {
		temp := v(cmd, data, 0, "")
		data, _ = proto.Marshal(temp)
	}
	conn.Ws.WriteMessage(s.Type, data)
	return
}

// 批量发送
func (s *SocketV1) SendAll(cmd int32, req proto.Message) {
	g.Log("cmd").Debugf(gctx.New(), fmt.Sprintf("all:%d|-1|%v", cmd, gjson.MustEncodeString(req)))

	//格式化数据
	var data, err = proto.Marshal(req)
	if err != nil {
		g.Log().Error(gctx.New(), err)
		return
	}

	for _, v := range Byte2Pb {
		temp := v(cmd, data, 0, "")
		data, _ = proto.Marshal(temp)
	}
	m.Iterator(func(k interface{}, v interface{}) bool {
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
	uid := conn.Uid
	if uid > 0 {
		s.UnBindUid(uid)
	}

	// 可能的后续操作：
	// 1. 更新连接状态或从连接池移除
	// 2. 发送通知或清理关联资源
	// 3. 执行特定于业务的断开处理
	m.Remove(conn.Uuid)
	conn.Ws.Close()
}

// 是否在线
func (s *SocketV1) IsOnline(uid int64) bool {
	uuid := s.Uid2Uuid(uid)
	if m.Contains(uuid) {
		return true
	}
	return false
}
