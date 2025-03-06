package callback

import (
	"context"
	"new-gitlab.adesk.com/public_project/utility_go/service"

	"new-gitlab.adesk.com/public_project/utility_go/api/callback/v1"
)

func (c *ControllerV1) Ip(ctx context.Context, req *v1.IpReq) (res *v1.IpRes, err error) {
	res = &v1.IpRes{}
	res.Address = service.Ip2Region().GetIp(req.Ip)
	return
}
