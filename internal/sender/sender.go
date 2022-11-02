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

	message := []byte(dataMail)

	auth := smtp.PlainAuth("", s.From, s.Password, s.SmtpHost)

	err := smtp.SendMail(s.SmtpHost+":"+s.SmtpPort, auth, s.From, to, message)
	if err != nil {
		return fmt.Errorf("send mail by smtp error: %s", err)
	}
	fmt.Println("Почта отправлена!")

	return err
}
