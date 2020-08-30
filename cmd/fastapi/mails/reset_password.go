package mails

import (
	"bytes"
	"html/template"
	"log"
	"path/filepath"
)

type ResetPasswordEmail struct {
	To string
	Name string
	Code string
}

func (s *ResetPasswordEmail) Send() {
	path,err:=filepath.Abs("mails/templates/reset_password.gohtml")
	if err!=nil{
		log.Println(err)
		return
	}
	tpl, err := template.ParseFiles(path)
	if err != nil {
		log.Println(err)
		return
	}
	var body bytes.Buffer
	var content []byte
	tpl.Execute(&body, s)
	content=body.Bytes()
	email:=Email{
		To:      []string{s.To},
		Subject: "Reset Password",
		Body:    content,
	}
	go email.Send()
}
