package simple_mail

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"

	"strings"
)




type Mail struct {
	Host      string
	Port      string
	Sender  string
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Body    string
	HtmlTheme bool
	tlsConfig *tls.Config

}

func NewMail(host string,port string,sender string,to []string,cc []string,bcc []string,subject string,body string,htmlTheme bool)*Mail  {
	tlsConf :=&tls.Config{
		InsecureSkipVerify: true,
		ServerName: host,
	}
	return &Mail{Host:host,Port:port,Sender:sender,To:to,Cc:cc,Bcc:bcc,Subject:subject,Body:body,HtmlTheme:htmlTheme,tlsConfig:tlsConf}
}

func (s *Mail) ServerName() string {
	return s.Host + ":" + s.Port
}

func (mail *Mail) buildMessage() string {
	header := ""
	header += fmt.Sprintf("From: %s\r\n", mail.Sender)
	if len(mail.To) > 0 {
		header += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	}
	if len(mail.Cc) > 0 {
		header += fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Cc, ";"))
	}
	header += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	if mail.HtmlTheme{
		header += fmt.Sprintf("MIME-version: %s\r\n","1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n")

	}
	header += "\r\n" + mail.Body

	return header
}

func (mail *Mail) SendMail() (success bool,err error) {
	messageBody:= mail.buildMessage()
	auth := smtp.PlainAuth("", mail.Sender, "mahdi1369QWE", mail.Host)

	conn, err := tls.Dial("tcp", mail.ServerName(), mail.tlsConfig)
	if err != nil {
		return false,err
	}
	client, err := smtp.NewClient(conn, mail.Host)
	if err != nil {

	}
	if err = client.Auth(auth); err != nil {
		return false,err

	}
	if err = client.Mail(mail.Sender); err != nil {
		return false,err

	}
	receivers := append(mail.To, mail.Cc...)
	receivers = append(receivers, mail.Bcc...)
	for _, k := range receivers {
		log.Println("sending to: ", k)
		if err = client.Rcpt(k); err != nil {
			return false,err

		}
	}
	w, err := client.Data()
	if err != nil {
		return false,err

	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		return false,err

	}

	err = w.Close()
	if err != nil {
		return false,err
	}

	client.Quit()
	return true,nil

}