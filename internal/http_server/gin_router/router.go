package gin_router

import (
	"mail_sender/internal/app"
	"mail_sender/internal/http_server"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	App app.App
}

func NewRouter(h Handler) (router *gin.Engine) {
	router = gin.Default()

	mails := router.Group("/send")
	{
		mails.POST("", h.Send)

	}

	return router
}

func (h Handler) Send(ctx *gin.Context) {

	req := new(http_server.SendMailRequest)
	if err := ctx.ShouldBind(req); err != nil {
		ctx.Status(http.StatusBadRequest)
		err = ctx.Error(err)
		return
	}

	if err := h.App.SendMails(req.Mails, req.TemplateId); err != nil {
		ctx.Status(http.StatusInternalServerError)
		err = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}
