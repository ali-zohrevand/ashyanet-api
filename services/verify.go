package services

import (
	"gitlab.com/hooshyar/ChiChiNi-API/core/DB"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	. "gitlab.com/hooshyar/ChiChiNi-API/services/mail"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
	"gopkg.in/mgo.v2"
	"strings"
)

func Verify(user models.UserInDB,Session *mgo.Session) (err error)  {
	settings, errLoadSettings := DB.LoadSettings(Session)

	if errLoadSettings!=nil{
		return
	}
	verifyUrl:= settings.Url+"/"+user.UserName+"/"+user.TempKeyGenreated
	MailMeesageInHtmlTHeme:=Words.VerifyMail
	MailMeesageInHtmlTHeme=strings.Replace(MailMeesageInHtmlTHeme,"***url***",verifyUrl,-1)
	verifyMail:= NewMail(settings.MailHost,settings.MailPort,settings.MailVerifyUsername,settings.MailVerifyPassword,[]string{user.Email},nil,nil,"Verfy mail",MailMeesageInHtmlTHeme,true)
	_,err=verifyMail.SendMail()

	return err
}
