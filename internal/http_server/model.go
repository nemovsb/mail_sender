package http_server

import (
	"mail_sender/internal/app"
	"time"
)

type SendMailRequest struct {
	MailingSendId string   `form:"mailingsendid"`
	Mails         []string `form:"mails"`
	TemplateId    uint     `form:"templateid"`
}

type DeferSendMailRequest struct {
	SendMailRequest

	ExecTime time.Time `form:"exectime"`
}

type CreateRecipientsRequest struct {
	Recipients []app.Recipient `form:"recipients"`
}

type CreateTemplateRequest struct {
	Template string
}

type TrackParam struct {
	From      string `form:"from"`
	MailingId string `form:"mailingId"`
	Event     string `form:"event"`
}
