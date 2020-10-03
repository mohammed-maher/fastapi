package mails

import (
	"bytes"
	"fmt"
	"github.com/mohammed-maher/fastapi/config"
	"log"
	"mime/quotedprintable"
	"net/smtp"
	"strings"
)

type Email struct {
	To      []string
	Subject string
	Body    []byte
}

func (e *Email) Send() {
	settings := config.Config.SMTP
	auth := smtp.PlainAuth("", settings.Username, settings.Password, settings.Host)
	header := map[string]string{
		"From":                      settings.Username,
		"Subject":                   e.Subject,
		"MIME-Version":              "1.0",
		"Content-Type":              fmt.Sprintf("%s; charset=\"utf-8\"", "text/html"),
		"Content-Disposition":       "inline",
		"Content-Transfer-Encoding": "quoted-printable",
	}
	var messageHeader strings.Builder
	for key, value := range header {
		messageHeader.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}
	var messageBody bytes.Buffer
	temp := quotedprintable.NewWriter(&messageBody)
	temp.Write(e.Body)
	temp.Close()
	finalMessage := messageHeader.String() + "\r\n" + messageBody.String()
	err := smtp.SendMail(fmt.Sprintf("%s:%d", settings.Host, settings.Port), auth, settings.Username, e.To, []byte(finalMessage))
	if err != nil {
		log.Println(err)
	}
}
