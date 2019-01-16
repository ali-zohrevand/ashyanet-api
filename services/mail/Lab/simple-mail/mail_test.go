package simple_mail

import (
	"crypto/tls"
	"github.com/ali-zohrevand/ashyanet-api/settings/Words"
	"log"
	"net/smtp"
	"testing"
)

func TestMAIL(t *testing.T) {
	mail := Mail{}
	mail.Sender = "ashyanet@gmail.com"
	mail.To = []string{"alihooshyar1990@gmail.com"}
	/*	mail.Cc = []string{"mnp@gmail.com"}
		mail.Bcc = []string{"a69@outlook.com"}*/
	mail.Subject = "تایید میل"
	mail.Body = Words.VerifyMail
	messageBody := mail.BuildMessage()

	smtpServer := SmtpServer{Host: "smtp.gmail.com", Port: "465"}
	smtpServer.TlsConfig = &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.Host,
	}

	auth := smtp.PlainAuth("", mail.Sender, "mahdi1369QWE", smtpServer.Host)

	conn, err := tls.Dial("tcp", smtpServer.ServerName(), smtpServer.TlsConfig)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	client, err := smtp.NewClient(conn, smtpServer.Host)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		t.Error(err)
		t.Fail()
	}

	// step 2: add all from and to
	if err = client.Mail(mail.Sender); err != nil {
		t.Error(err)
		t.Fail()
	}
	receivers := append(mail.To, mail.Cc...)
	receivers = append(receivers, mail.Bcc...)
	for _, k := range receivers {
		log.Println("sending to: ", k)
		if err = client.Rcpt(k); err != nil {
			t.Error(err)
			t.Fail()
		}
	}

	// Data
	w, err := client.Data()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	client.Quit()

	log.Println("Mail sent successfully")

}
