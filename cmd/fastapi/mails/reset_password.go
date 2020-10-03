package mails

import (
	"bytes"
	"github.com/mohammed-maher/fastapi/helpers"
	"html/template"
	"log"
)

type ResetPasswordEmail struct {
	To   string
	Name string
	Code string
}
var tplCache []byte

func (s *ResetPasswordEmail) Send() {
	if tplCache==nil{
		tplObj:=helpers.StorageObject{
			Key:    "reset_password.gohtml",
			Bucket: "internals",
			File:   nil,
		}
		tplData,err:=tplObj.Get()
		if err!=nil{
			log.Println(err)
			return
		}
		tplCache=tplData
	}

	tpl, err := template.New("reset").Parse(string(tplCache))
	if err != nil {
		log.Println(err)
		return
	}
	var body bytes.Buffer
	var content []byte
	tpl.Execute(&body, s)
	content = body.Bytes()
	email := Email{
		To:      []string{s.To},
		Subject: "Reset Password",
		Body:    content,
	}
	go email.Send()
}
