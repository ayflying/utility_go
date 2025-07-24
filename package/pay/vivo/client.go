package vivo

type Pay struct {
	AppId  string
	AppKey string
	//AppSecret string
}

func New(cfg *Pay) (client *Pay) {
	return &Pay{
		AppId:  cfg.AppId,
		AppKey: cfg.AppKey,
		//AppSecret: cfg.AppSecret,
	}
}
