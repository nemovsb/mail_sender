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
		mails.POST("/defer", h.DeferSend)

	}

	recipient := router.Group("/recipient")
	{
		recipient.GET("", h.GetRecipients)
		recipient.POST("/create", h.CreateRecipients)
	}

	template := router.Group("/template")
	{
		template.POST("/create", h.CreateTemplate)
		template.GET("", h.GetAllTemplates)
	}

	tracker := router.Group("/track")
	{
		tracker.GET("", h.Track)
	}

	return router
}

func (h Handler) Track(ctx *gin.Context) {

	param := new(http_server.TrackParam)
	if err := ctx.ShouldBind(param); err != nil {
		ctx.Status(http.StatusBadRequest)
		err = ctx.Error(err)
		return
	}

	h.App.Track(http_server.GetTrackParam(*param))

	ctx.Status(http.StatusOK)

}

func (h Handler) Send(ctx *gin.Context) {

	req := new(http_server.SendMailRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Status(http.StatusBadRequest)
		err = ctx.Error(err)
		return
	}

	if err := h.App.SendMails(req.MailingSendId, req.Mails, req.TemplateId); err != nil {
		ctx.Status(http.StatusInternalServerError)
		err = ctx.Error(err)

		log.Println(err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (h Handler) DeferSend(ctx *gin.Context) {

	req := new(http_server.DeferSendMailRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Status(http.StatusBadRequest)
		err = ctx.Error(err)
		return
	}

	task := http_server.GetDeferMailingTask(*req)

	mailingId, err := h.App.AddMailingTask(task)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		err = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"mailingSendId": mailingId,
	})

}

func (h Handler) CreateRecipients(ctx *gin.Context) {

	req := new(http_server.CreateRecipientsRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Status(http.StatusBadRequest)
		err = ctx.Error(err)
		return
	}

	res, err := h.App.CreateRecipients(req.Recipients)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		err = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"recipientsAdded": res,
	})
}

func (h Handler) GetRecipients(ctx *gin.Context) {

	res, err := h.App.GetAllRecipients()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		err = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)

}
func (h Handler) CreateTemplate(ctx *gin.Context) {

	req := new(http_server.CreateTemplateRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Status(http.StatusBadRequest)
		err = ctx.Error(err)
		return
	}

	id, err := h.App.CreateTemplate(req.Template)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"templateid": id,
	})
}

func (h Handler) GetAllTemplates(ctx *gin.Context) {

	templates, err := h.App.GetAllTemplates()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"templates": templates,
	})

}
