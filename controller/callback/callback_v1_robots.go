package callback

import (
	"context"

	"github.com/ayflying/utility_go/api/callback/v1"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Robots(ctx context.Context, req *v1.RobotsReq) (res *v1.RobotsRes, err error) {
	text := "User-agent: *\nDisallow: /"
	g.RequestFromCtx(ctx).Response.Write(text)
	return
}
