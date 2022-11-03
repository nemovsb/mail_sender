package sender

import (
	"fmt"
	"net/smtp"
)

type Sender struct {
	From     string
	Password string

	SmtpHost string
	SmtpPort string
}

func NewSender(from, pswrd, host, port string) Sender {
	return Sender{
		From:     from,
		Password: pswrd,
		SmtpHost: host,
		SmtpPort: port,
	}
}

func (s Sender) Send(mailAddr, dataMail string) error {

	to := []string{mailAddr}

	message := []byte(
		"From: MailApp <" + s.From + ">\r\n" +
			"To: " + to[0] + "\r\n" +
			"Subject: Test MailApp with golang\r\n" +
			"MIME: MIME-version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
			"\r\n" + dataMail)

	auth := smtp.PlainAuth("", s.From, s.Password, s.SmtpHost)

	err := smtp.SendMail(s.SmtpHost+":"+s.SmtpPort, auth, s.From, to, message)
	if err != nil {
		return fmt.Errorf("send mail by smtp error: %s", err)
	}

	return err
}
