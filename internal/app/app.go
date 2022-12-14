package app

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type App struct {
	Storage
	Aggregator
	Sender
	Tracker
}

type Storage interface {

	//Get recipients by email-adresses
	GetRecipients(mailAddrs []string) ([]Recipient, error)

	//Get all recipients from storage
	GetAllRecipients() []Recipient

	//Get template by id
	GetTemplate(id uint) (string, error)

	//Get all templates from storage
	GetAllTemplates() []string

	//Create recipient
	CreateRecipients(recipients []Recipient) uint

	//Create template
	CreateTemplate(string) (id uint)

	AddMailingTask(MailingTask) string

	GetMailingTasks() []MailingTask
}

type Aggregator interface {
	Aggregate(recipient Recipient, HTMLPattern *string) (dataMail string, err error)
}

type Sender interface {
	Send(mailingSendId, mailAddr, dataMail string) error
}

type Tracker interface {
	Track(TrackMailParam)
}

func NewApp(strg Storage, agr Aggregator, sendr Sender, tracker Tracker) App {
	return App{
		Storage:    strg,
		Aggregator: agr,
		Sender:     sendr,
		Tracker:    tracker,
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

func (a App) Send(mailingSendId, mailAddr, dataMail string) error {
	return a.Sender.Send(mailingSendId, mailAddr, dataMail)
}

func (a App) SendMails(mailingSendId string, mailAddrs []string, templateId uint) error {

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

		err = a.Send(mailingSendId, recipient.MailAddr, dataMail)
		if err != nil {
			return fmt.Errorf("send mail to %s error: %s", recipient.MailAddr, err)
		}

	}

	return err
}

func (a App) AddRecipients(recipients []Recipient) uint {
	return a.Storage.CreateRecipients(recipients)
}

func (a App) GetAllRecipients() []Recipient {
	return a.Storage.GetAllRecipients()
}

func (a App) GetAllTemplates() []string {
	return a.Storage.GetAllTemplates()
}

func (a App) CreateTemplate(template string) (id uint) {
	return a.Storage.CreateTemplate(template)
}

func (a App) Track(param TrackMailParam) {
	a.Tracker.Track(param)
}

func (a App) AddMailingTask(task MailingTask) (mailingId string) {
	return a.Storage.AddMailingTask(task)
}

func GetCecker(a *App) func() {
	return func() {
		for {
			tasks := a.Storage.GetMailingTasks()

			log.Println("----- Mailing Tasks ------- :")
			for _, task := range tasks {
				fmt.Printf("Id: %s,	Time Exec: %s\n", task.MailingSendId, task.ExecTime)
			}
			fmt.Println("-----------------------------")

			var wg sync.WaitGroup

			for _, task := range tasks {

				wg.Add(1)

				go func(timeToExec time.Time, mailingSendId string, mailAddrs []string, templateId uint) {
					defer wg.Done()

					time.Sleep(time.Until(timeToExec))

					err := a.SendMails(mailingSendId, mailAddrs, templateId)
					if err != nil {
						log.Printf("mail send error: %s", err)
					}

				}(task.ExecTime, task.MailingSendId, task.MailAddrs, task.TemplateId)

				wg.Wait()

			}

			time.Sleep(time.Minute)
		}
	}
}
