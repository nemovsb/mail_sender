package aggregator

import (
	"bytes"
	"fmt"
	"mail_sender/internal/app"
	"text/template"
)

type Aggregator struct {
}

func (a Aggregator) Aggregate(recipient app.Recipient, HTMLPattern *string) (dataMail string, err error) {

	tmpl, err := template.New("data").Parse(*HTMLPattern)
	if err != nil {
		return "", fmt.Errorf("parse pattern error: %s", err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, recipient)
	if err != nil {
		return "", fmt.Errorf("template filling error: %s", err)
	}

	dataMail = buf.String()

	return dataMail, nil
}
