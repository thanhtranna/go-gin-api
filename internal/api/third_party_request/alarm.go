package third_party_request

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/pkg/httpclient"
	"github.com/xinliangnote/go-gin-api/pkg/mail"

	"github.com/pkg/errors"
)

// Implement AlarmObject alarm
var _ httpclient.AlarmObject = (*AlarmEmail)(nil)

type AlarmEmail struct{}

// Email alert method
func (a *AlarmEmail) Send(subject, body string) error {
	cfg := configs.Get().Mail
	if cfg.Host == "" || cfg.Port == 0 || cfg.User == "" || cfg.Pass == "" || cfg.To == "" {
		return errors.New("mail config error")
	}

	options := &mail.Options{
		MailHost: cfg.Host,
		MailPort: cfg.Port,
		MailUser: cfg.User,
		MailPass: cfg.Pass,
		MailTo:   cfg.To,
		Subject:  subject,
		Body:     body,
	}

	return mail.Send(options)
}
