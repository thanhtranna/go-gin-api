package mail

import (
	"testing"
)

func TestSend(t *testing.T) {
	options := &Options{
		MailHost: "smtp.gmail.com",
		MailPort: 465,
		MailUser: "xxx@gmail.com",
		MailPass: "", // Password or authorization code
		MailTo:   "",
		Subject:  "subject",
		Body:     "body",
	}
	err := Send(options)
	if err != nil {
		t.Error("Mail Send error", err)
		return
	}
	t.Log("success")
}
