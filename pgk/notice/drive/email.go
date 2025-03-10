package drive

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"gopkg.in/gomail.v2"
)

type MailMod struct {
	Host    string `json:"host" dc:"邮件服务器地址"`
	Port    int    `json:"port" dc:"邮件服务器端口"`
	User    string `json:"user" dc:"邮件服务器用户名"`
	Pass    string `json:"pass" dc:"邮件服务器密码"`
	From    string `json:"from" dc:"邮件发送者"`
	To      string `json:"to" dc:"邮件接收者"`
	Subject string `json:"subject" dc:"邮件主题"`
}

func MailLoad(Host string, port int, to string, subject string) *MailMod {
	return &MailMod{
		Host:    Host,
		Port:    port,
		User:    "root",
		Pass:    "root",
		From:    "root",
		To:      to,
		Subject: subject,
	}
}

func (m MailMod) Send(value string) {
	// 创建一个新的消息
	obj := gomail.NewMessage()
	// 设置发件人
	obj.SetHeader("From", m.From)
	// 设置收件人
	obj.SetHeader("To", m.To)
	// 设置邮件主题
	obj.SetHeader("Subject", m.Subject)
	// 设置邮件正文
	obj.SetBody("text/plain", value)

	// 创建 SMTP 拨号器，这里需要提供 SMTP 服务器地址、端口、发件人邮箱和密码
	d := gomail.NewDialer(m.Host, m.Port, m.User, m.Pass)

	// 发送邮件
	if err := d.DialAndSend(obj); err != nil {
		g.Log().Error(gctx.New(), err)
	}
	return
}
