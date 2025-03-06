package cmd

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"new-gitlab.adesk.com/public_project/utility_go/service"
)

//func MiddlewareAnonymous(r *ghttp.Request) {
//	// 中间件处理逻辑
//	r.Response.CORSDefault()
//
//	ip := r.GetClientIp()
//	r.SetCtxVar("ip", ip)
//
//	//各种回调的日志返回
//	//get, _ := r.GetJson()
//	get := r.GetRequestMapStrStr()
//	delete(get, "r")
//	delete(get, "s")
//	delete(get, "t")
//	delete(get, "data")
//	getJson, _ := gjson.EncodeString(get)
//	g.Log("cmd").Debugf(r.GetCtx(), "from|%v|%v|%v", 0, r.RequestURI, getJson)
//
//	r.Middleware.Next()
//
//	//中间件后置
//	err := r.GetError()
//	if err != nil {
//		code, err2 := strconv.Atoi(err.Error())
//		if err2 != nil {
//			return
//		}
//		if _, ok := consts.ErrCodeList[code]; ok {
//			//g.Dump("===error:", gerror.Code(r.GetError()).Code())
//			msg := g.Map{
//				"code":    code,
//				"message": consts.ErrCodeList[code],
//				//"data":    r.GetHandlerResponse(),
//			}
//			r.Response.WriteJson(msg)
//			//错误码置空
//			r.SetError(nil)
//			//g.Log("cmd").Debugf(r.GetCtx(), "to|%v|%v|%v", uid, r.RequestURI, msg)
//			return
//		}
//	}
//
//	//回复
//	res, _ := gjson.EncodeString(r.GetHandlerResponse())
//	g.Log("cmd").Debugf(r.GetCtx(), "to|%v|%v|%v", 0, r.RequestURI, res)
//}

func MiddlewareAdmin(r *ghttp.Request) {
	// 中间件处理逻辑
	r.Response.CORSDefault()

	ip := r.GetClientIp()
	r.SetCtxVar("ip", ip)

	get := r.Cookie.Get("uid")

	if get == nil {
		//调试模式允许不验证用户名
		debug, _ := g.Cfg().GetWithEnv(nil, "debug")
		if !debug.Bool() {

			msg := g.Map{
				"code":    403,
				"message": "登录失败",
				//"data":    r.GetHandlerResponse(),
			}
			//r.SetError(http.Error(r,"403",http.StatusForbidden))
			r.Response.WriteJson(msg)
			gerror.NewCode(gcode.CodeNil, "登录失败")
			return
		}

	}

	uid := get.Int()

	r.Middleware.Next()

	//后置，所有post都写入日志
	if r.Method == "POST" {
		//黑名单列表
		LogUrl := []string{
			"/system/chatgpt",
		}
		if !gstr.InArray(LogUrl, r.RequestURI) {
			//写入日志
			service.SystemLog().AddLog(uid, r.RequestURI, ip, r.GetFormMap())
		}
	} else {
		//需要写入的get
		LogUrl := []string{
			"/admin/config/mall/del",
			"/admin/config/shop/del",
			"/admin/group/del",
			"/admin/user/del",
			"/admin/community/posts/del",
			"/admin/community/posts/limit",
			"/admin/community/posts/limit/del",
			"/admin/community/reply/del",
			"/admin/community/recommend",
		}
		for _, item := range LogUrl {
			if item == r.RequestURI {
				service.SystemLog().AddLog(uid, r.RequestURI, ip, r.GetFormMap())
			}
		}
	}

}

//// 中间件
//func Middleware(r *ghttp.Request) {
//	// 中间件处理逻辑
//	r.Response.CORSDefault()
//
//	//获取玩家的guid
//	guid := r.Header.Get("guid")
//
//	//获取所有请求的信息
//	get := r.GetRequestMapStrStr()
//
//	//cacheKey := fmt.Sprintf("sign:%s", guid)
//
//	////进入debug模式
//	//debugBool := g.Cfg().MustGetWithCmd(nil, "debug")
//	////如果收到签名，开始验证签名
//	//if sign, _ := get["s"]; !debugBool.Bool() && sign != "" {
//	//
//	//	//如果连续两次使用相同sign，直接抛出
//	//	getSign, _ := aycache.New().Get(nil, cacheKey)
//	//	if getSign.String() == sign {
//	//		//中间件授权错误
//	//		msg := g.Map{
//	//			"code":    11000,
//	//			"message": consts.ErrCodeList[11000],
//	//		}
//	//		r.Response.WriteJson(msg)
//	//		return
//	//	}
//	//	aycache.New().Set(nil, cacheKey, sign, time.Minute*10)
//	//
//	//	secretKey := "asdkjqwhiasdoplmwofjk/aws"
//	//	nonce := get["r"]
//	//	timestamp := get["t"]
//	//
//	//	message := timestamp + nonce
//	//
//	//	timeUnix := time.Now().Unix() - gconv.Int64(timestamp)
//	//	if timeUnix > 600 || timeUnix < -600 {
//	//		//中间件授权错误
//	//		msg := g.Map{
//	//			"code":    11000,
//	//			"message": consts.ErrCodeList[11000],
//	//		}
//	//		r.Response.WriteJson(msg)
//	//		return
//	//	}
//	//
//	//	// 创建 HMAC 对象
//	//	h := hmac.New(sha256.New, []byte(secretKey))
//	//
//	//	// 更新 HMAC 对象的数据
//	//	h.Write([]byte(message))
//	//
//	//	// 获取 HMAC 的十六进制表示
//	//	signature := hex.EncodeToString(h.Sum(nil))
//	//
//	//	//如果加密算法不一致
//	//	if signature != sign {
//	//		//中间件授权错误
//	//		msg := g.Map{
//	//			"code":    11000,
//	//			"message": consts.ErrCodeList[11000],
//	//		}
//	//		r.Response.WriteJson(msg)
//	//		return
//	//	}
//	//}
//
//	uid, _ := service.MemberUser().Guid2uid(guid)
//	if uid == 0 {
//		//中间件授权错误
//		msg := g.Map{
//			"code":    11000,
//			"message": consts.ErrCodeList[11000],
//			//"data":    r.GetHandlerResponse(),
//		}
//		r.Response.WriteJson(msg)
//
//		return
//	}
//	r.SetCtxVar("guid", guid)
//	r.SetCtxVar("uid", uid)
//
//	ip := r.GetClientIp()
//	r.SetCtxVar("ip", ip)
//
//	delete(get, "r")
//	delete(get, "s")
//	delete(get, "t")
//	delete(get, "data")
//	//前置输出服务器收到信息
//	getJson, _ := gjson.EncodeString(get)
//	//后置输出服务器返回信息
//	if r.GetCtxVar("not_log").IsEmpty() {
//		g.Log("cmd").Debugf(r.GetCtx(), "from|%v|%v|%v", uid, r.RequestURI, getJson)
//	}
//
//	//运行开始时间
//	RunStartTime := gtime.Now()
//
//	//proto.Marshal()
//
//	//中间件核心
//	r.Middleware.Next()
//
//	//返回运行时
//	if getTime := gtime.Now().Sub(RunStartTime); getTime > time.Millisecond*1000 {
//		g.Log().Debugf(nil, "当前运行时间:%v,uid=%d,url=%s", getTime, uid, r.RequestURI)
//	}
//
//	//中间件后置
//	err := r.GetError()
//	if err != nil {
//		code, err2 := strconv.Atoi(err.Error())
//		if err2 != nil {
//			return
//		}
//		if _, ok := consts.ErrCodeList[code]; ok {
//			//g.Dump("===error:", gerror.Code(r.GetError()).Code())
//			msg := g.Map{
//				"code":    code,
//				"message": consts.ErrCodeList[code],
//				//"data":    r.GetHandlerResponse(),
//			}
//			msgJson, _ := gjson.EncodeString(msg)
//			r.Response.WriteJson(msgJson)
//			//错误码置空
//			r.SetError(nil)
//			g.Log("cmd").Debugf(r.GetCtx(), "to|%v|%v|%v", uid, r.RequestURI, msg)
//			return
//		}
//	}
//
//	//后置输出服务器返回信息
//	if r.GetCtxVar("not_log").IsEmpty() {
//		res, _ := gjson.EncodeString(r.GetHandlerResponse())
//		g.Log("cmd").Debugf(r.GetCtx(), "to|%v|%v|%v", uid, r.RequestURI, res)
//	}
//
//}
