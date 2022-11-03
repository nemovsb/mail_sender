package gin_router

import (
	"log"
	"mail_sender/internal/app"
	"mail_sender/internal/http_server"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	App app.App
}

func NewHandler(app app.App) Handler {
	return Handler{
		App: app,
	}
}

func NewRouter(h Handler) (router *gin.Engine) {
	router = gin.Default()

	mails := router.Group("/send")
	{
		mails.POST("", h.Send)

	}

	recipient := router.Group("/recipient")
	{
		recipient.POST("/create", h.CreateRecipients)
	}

	return router
}

func (h Handler) Send(ctx *gin.Context) {

	req := new(http_server.SendMailRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Status(http.StatusBadRequest)
		err = ctx.Error(err)
		return
	}

	if err := h.App.SendMails(req.Mails, req.TemplateId); err != nil {
		ctx.Status(http.StatusInternalServerError)
		err = ctx.Error(err)

		log.Println(err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (h Handler) CreateRecipients(ctx *gin.Context) {

	//fmt.Printf("--------------CHECK ctx.req:  %+v", ctx.Request)

	req := new(http_server.CreateRecipientsRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Status(http.StatusBadRequest)
		err = ctx.Error(err)
		return
	}

	res := h.App.CreateRecipients(req.Recipients)
	ctx.JSON(http.StatusOK, gin.H{
		"recipientsAdded": res,
	})
}
