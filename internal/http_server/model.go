package http_server

type SendMailRequest struct {
	Mails      []string `form:"mails" binding:"required"`
	TemplateId uint     `form:"templateid" binding:"required"`
}
