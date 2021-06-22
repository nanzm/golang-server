package datasource

import (
	"dora/config"
	"dora/pkg/utils/logx"
	"net/smtp"
	"net/textproto"
	"sync"

	"github.com/jordan-wright/email"
)

var mailPool *email.Pool
var once sync.Once

func GetMailPool() *email.Pool {
	once.Do(func() {
		conf := config.GetMail()
		mailPool = setUp(conf)
	})
	return mailPool
}

func setUp(conf config.MailConfig) *email.Pool {
	addr := conf.Host + ":" + conf.Port

	p, err := email.NewPool(addr, 4,
		smtp.PlainAuth(conf.Password, conf.Username, conf.Password, conf.Host))

	if err != nil {
		panic(err)
	}
	return p
}

func BuilderEmail(to, from, subject, body string) *email.Email {
	m := &email.Email{
		To:      []string{to},
		From:    from,
		Subject: subject,
		HTML:    []byte(body),
		Headers: textproto.MIMEHeader{},
	}
	return m
}

func StopMailPool() {
	logx.Println("stop mail pool")
	GetMailPool().Close()
}
