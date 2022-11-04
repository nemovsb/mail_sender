package http_server

import (
	"mail_sender/internal/app"
)

func GetTrackParam(req TrackParam) app.TrackMailParam {

	return app.TrackMailParam{
		From:      req.From,
		MailingId: req.MailingId,
		Event:     req.Event,
	}

}

func GetDeferMailingTask(req DeferSendMailRequest) app.MailingTask {

	return app.MailingTask{
		ExecTime:      req.ExecTime,
		MailingSendId: req.MailingSendId,
		MailAddrs:     req.Mails,
		TemplateId:    req.TemplateId,
	}

}
