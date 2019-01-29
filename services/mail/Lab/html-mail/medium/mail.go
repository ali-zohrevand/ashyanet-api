package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
	"path/filepath"
)

var auth smtp.Auth

func main() {
	auth = smtp.PlainAuth("", "ashyanet@gmail.com", "mahdi1369QWE", "smtp.gmail.com")
	templateData := struct {
		Name string
		URL  string
	}{
		Name: "Dhanush",
		URL:  "http://geektrust.in",
	}
	r := NewRequest([]string{"alihooshyar1990@gmail.com"}, "Hello Junk!", "Hello, World!")
	cwd, _ := os.Getwd()
	path:=filepath.Join( cwd, "template.html" )
	err := r.ParseTemplate(path, templateData)
	if err!=nil{
		fmt.Print(err)
		panic(err)
	}
	if err := r.ParseTemplate(path, templateData); err == nil {
		ok, _ := r.SendEmail()
		fmt.Println(ok)
	}

}

//Request struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) SendEmail() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := "smtp.gmail.com:587"

	if err := smtp.SendMail(addr, auth, "dhanush@geektrust.in", r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}