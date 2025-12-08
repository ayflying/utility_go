package v1

import "github.com/gogf/gf/v2/frame/g"

type RobotsReq struct {
	g.Meta `path:"/robots.txt" tags:"回调响应" method:"get" summary:"禁止爬虫"`
}
type RobotsRes struct {
}
