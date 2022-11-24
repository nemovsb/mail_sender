package storage

import (
	"log"
	"mail_sender/internal/app"
	"sync"
	"time"
)

type MockStorage struct {
	mu           sync.Mutex
	recipients   []app.Recipient
	templates    []string
	mailingTasks []app.MailingTask
}

func NewMockStorage() *MockStorage {

	recp := []app.Recipient{}
	recp = append(recp, app.Recipient{
		MailAddr: "testtesttest9@mailtest.ru",
		Name:     "Ivan",
		Surname:  "Ivanov",
		Birthday: "01.01.2000",
	})
	recp = append(recp, app.Recipient{
		MailAddr: "testtesttest8@mailtest.ru",
		Name:     "Petr",
		Surname:  "Petrov",
		Birthday: "01.01.2001",
	})
	recp = append(recp, app.Recipient{
		MailAddr: "nemoff.serega@mail.ru",
		Name:     "Ser",
		Surname:  "Serg",
		Birthday: "01.01.2001",
	})

	templs := []string{
		`<!DOCTYPE html>
		<html>
			<head>
				<meta charset="UTF-8">
				<title>{{ .Name}}</title>
			</head>
			<body>
				<h1>{{ .Name}}</h1>
				<p>{{ .Birthday}}</p>
			</body>
		</html>`,

		`<!DOCTYPE html>
		<html>
			<head>
				<meta charset="UTF-8">
				<title>{{ .Surname}}</title>
			</head>
			<body>
				<h1>{{ .Surname}}</h1>
				<p>{{ .Birthday}}</p>
			</body>
		</html>`,

		`Тестовое сообщение через golang.`,
	}
	return &MockStorage{
		recipients: recp,
		templates:  templs,
	}
}

func (ms *MockStorage) GetRecipients(mailAddrs []string) ([]app.Recipient, error) {

	recipients := []app.Recipient{}

	for _, AddrSearch := range mailAddrs {

		for _, recipient := range ms.recipients {

			if AddrSearch == recipient.MailAddr {
				recipients = append(recipients, recipient)
			}
		}
	}
	return recipients, nil
}

func (ms *MockStorage) GetTemplate(id uint) (string, error) {
	return ms.templates[id], nil
}

func (ms *MockStorage) CreateRecipients(recipients []app.Recipient) uint {

	var check bool = true
	numElem := 0

	ms.mu.Lock()

	for _, recp := range recipients {

		for _, searchRecp := range ms.recipients {

			check = (recp.MailAddr == searchRecp.MailAddr)
			if check {
				break
			}

		}
		if !check {
			ms.recipients = append(ms.recipients, recp)
			numElem++
		}
	}

	ms.mu.Unlock()

	return uint(numElem)

}

func (ms *MockStorage) GetAllRecipients() []app.Recipient {
	return ms.recipients
}

func (ms *MockStorage) CreateTemplate(template string) (id uint, err error) {

	ms.mu.Lock()

	ms.templates = append(ms.templates, template)

	id = uint(len(ms.templates)) - 1

	ms.mu.Unlock()

	return id, nil
}

func (ms *MockStorage) GetAllTemplates() ([]string, error) {
	return ms.templates, nil
}

func (ms *MockStorage) AddMailingTask(task app.MailingTask) (SendingId string) {

	ms.mu.Lock()

	ms.mailingTasks = append(ms.mailingTasks, task)

	ms.mu.Unlock()

	log.Printf("----AddMailingTask Tasks Added: --------\n%+v\n", ms.mailingTasks)

	return task.MailingSendId
}

func (ms *MockStorage) GetMailingTasks() []app.MailingTask {

	tasks := []app.MailingTask{}
	for i, task := range ms.mailingTasks {

		if time.Until(task.ExecTime) < time.Minute {

			tasks = append(tasks, task)

			ms.mu.Lock()

			ms.mailingTasks[i] = ms.mailingTasks[len(ms.mailingTasks)-1]
			ms.mailingTasks[len(ms.mailingTasks)-1] = app.MailingTask{}
			ms.mailingTasks = ms.mailingTasks[:len(ms.mailingTasks)-1]

			ms.mu.Unlock()

		}
	}

	return tasks
}
