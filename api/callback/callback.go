// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package callback

import (
	"context"

	"github.com/ayflying/utility_go/api/callback/v1"
)

type ICallbackV1 interface {
	Ip(ctx context.Context, req *v1.IpReq) (res *v1.IpRes, err error)
	Robots(ctx context.Context, req *v1.RobotsReq) (res *v1.RobotsRes, err error)
}
