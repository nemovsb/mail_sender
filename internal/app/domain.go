package app

import "time"

type Recipient struct {
	MailAddr string `form:"mailaddr"`
	Name     string `form:"name"`
	Surname  string `form:"surname"`
	Birthday string `form:"birthday"`
}

type TrackMailParam struct {
	From      string
	MailingId string
	Event     string
}

type MailingTask struct {
	ExecTime time.Time

	MailingSendId string
	MailAddrs     []string
	TemplateId    uint
}
