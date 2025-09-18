package casdoor

import (
	"github.com/ayflying/utility_go/service"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type sCasdoor struct {
	client *casdoorsdk.Client
	config *casdoorsdk.AuthConfig
}

func init() {
	var casdoor = &sCasdoor{}

	service.RegisterCasdoor(New(casdoor))
}

func New(s *sCasdoor) *sCasdoor {
	return s
}

func (s *sCasdoor) Load(endpoint string, clientId string, clientSecret string, certificate string, organizationName string, applicationName string) {
	s.config = &casdoorsdk.AuthConfig{
		Endpoint:         endpoint,
		ClientId:         clientId,
		ClientSecret:     clientSecret,
		Certificate:      certificate,
		OrganizationName: organizationName,
		ApplicationName:  applicationName,
	}
}

func (s *sCasdoor) New() *casdoorsdk.Client {
	if s.config == nil {
		g.Log().Errorf(gctx.New(), "未读取到配置，请先加载Load方法")
		return nil
	}
	s.client = casdoorsdk.NewClient(
		s.config.Endpoint,
		s.config.ClientId,
		s.config.ClientSecret,
		s.config.Certificate,
		s.config.OrganizationName,
		s.config.ApplicationName,
	)
	return s.client
}

//func (s *sCasdoor) EditPassword(name, oldPassword, newPassword string) (res bool, err error) {
//	res, err = s.client.SetPassword(s.config.OrganizationName, name, oldPassword, newPassword)
//	return
//}
//
//func (s *sCasdoor) Edit() {
//	s.client.GetGroups()
//
//}
