package storage

import (
	"fmt"
	"mail_sender/internal/app"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetRecipients(t *testing.T) {

	s := NewMockStorage()

	req := []string{"testtesttest9@mailtest.ru", "testtesttest8@mailtest.ru"}

	recipients, err := s.GetRecipients(req)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	expected := []app.Recipient{
		{
			MailAddr: "testtesttest9@mailtest.ru",
			Name:     "Ivan",
			Surname:  "Ivanov",
			Birthday: "01.01.2000",
		},
		{
			MailAddr: "testtesttest8@mailtest.ru",
			Name:     "Petr",
			Surname:  "Petrov",
			Birthday: "01.01.2001",
		},
	}

	require.Equal(t, recipients, expected)

}
