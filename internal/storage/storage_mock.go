package storage

import (
	"mail_sender/internal/app"
)

type MockStorage struct {
	recipients []app.Recipient
	templates  []string
}

func NewMockStorage() MockStorage {

	recp := make([]app.Recipient, 10)
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
	}
	return MockStorage{
		recipients: recp,
		templates:  templs,
	}
}

func (ms MockStorage) GetRecipients(mailAddrs []string) ([]app.Recipient, error) {

	recipients := make([]app.Recipient, cap(mailAddrs))

	for _, AddrSearch := range mailAddrs {

		for _, recipient := range ms.recipients {

			if AddrSearch == recipient.MailAddr {
				recipients = append(recipients, recipient)
			}
		}
	}
	return recipients, nil
}

func (ms MockStorage) GetTemplate(id uint) (string, error) {
	return ms.templates[id], nil
}
