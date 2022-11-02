package app

import "fmt"

type App struct {
	Storage
	Aggregator
	Sender
}

type Storage interface {
	GetRecipients(mailAddrs []string) ([]Recipient, error)
	GetTemplate(id uint) (string, error)
}

type Aggregator interface {
	Aggregate(recipient Recipient, HTMLPattern *string) (dataMail string, err error)
}

type Sender interface {
	Send(mailAddr, dataMail string) error
}

func NewApp(strg Storage, agr Aggregator, sendr Sender) App {
	return App{
		Storage:    strg,
		Aggregator: agr,
		Sender:     sendr,
	}
}

func (a App) GetTemplate(id uint) (string, error) {
	return a.Storage.GetTemplate(id)
}

func (a App) GetRecipients(mailAddrs []string) ([]Recipient, error) {
	return a.Storage.GetRecipients(mailAddrs)
}

func (a App) AggregateMail(recipient Recipient, HTMLPattern *string) (dataMail string, err error) {
	return a.Aggregator.Aggregate(recipient, HTMLPattern)
}

func (a App) Send(mailAddr, dataMail string) error {
	return a.Sender.Send(mailAddr, dataMail)
}

func (a App) SendMails(mailAddrs []string, templateId uint) error {

	pattern, err := a.GetTemplate(templateId)
	if err != nil {
		return fmt.Errorf("get template error: %s", err)
	}

	recipients, err := a.GetRecipients(mailAddrs)
	if err != nil {
		return fmt.Errorf("get template error: %s", err)
	}

	for _, recipient := range recipients {

		dataMail, err := a.AggregateMail(recipient, &pattern)
		if err != nil {
			return fmt.Errorf("agregate mail for %s error: %s", recipient.MailAddr, err)
		}

		err = a.Send(recipient.MailAddr, dataMail)
		if err != nil {
			return fmt.Errorf("send mail to %s error: %s", recipient.MailAddr, err)
		}

	}

	return err
}
