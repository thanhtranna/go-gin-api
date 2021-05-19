package mail

import (
	"strings"

	"gopkg.in/gomail.v2"
)

type Options struct {
	MailHost string
	MailPort int
	MailUser string // sender
	MailPass string // sender password
	MailTo   string // multiple recipients, split
	Subject  string // Email subject
	Body     string // Email content
}

func Send(o *Options) error {

	m := gomail.NewMessage()

	// Set sender
	m.SetHeader("From", o.MailUser)

	// Set to send to multiple users
	mailArrTo := strings.Split(o.MailTo, ",")
	m.SetHeader("To", mailArrTo...)

	// Set email subject
	m.SetHeader("Subject", o.Subject)

	// Set the message body
	m.SetBody("text/html", o.Body)

	d := gomail.NewDialer(o.MailHost, o.MailPort, o.MailUser, o.MailPass)

	return d.DialAndSend(m)
}
