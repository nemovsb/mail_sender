package http_server

import "mail_sender/internal/app"

type SendMailRequest struct {
	Mails      []string `form:"mails"`
	TemplateId uint     `form:"templateid"`
}

type CreateRecipientsRequest struct {
	Recipients []app.Recipient `form:"recipients"`
}

type CreateTemplateRequest struct {
	Template string
}
