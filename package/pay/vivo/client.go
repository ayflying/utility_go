package vivo

type Pay struct {
	AppId     string
	AppKey    string
	AppSecret string
}

func New(appId, appKey, appSecret string) (client *Pay) {
	return &Pay{
		AppId:     appId,
		AppKey:    appKey,
		AppSecret: appSecret,
	}
}
