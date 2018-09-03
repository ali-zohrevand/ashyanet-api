package services

import (
	"fmt"
	"gitlab.com/hooshyar/ChiChiNi-API/core/DB"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/services/log"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
)

func AddInitUser() {
	UserDatastore := DB.UserDataStore{}
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		panic(errConnectDB)
	}
	// چک میکنیم ببینیم آیا قبلا کاربر ادمین ثبت نام شده یا خیر
	user, _ := DB.FindUserByUsername(Words.DeafualtAdmminUserName, session)
	if user.Role != Words.DeafualtAdmminRole {
		//کاربر ادمین را ایجاد می نماییم.
		DefaultAdmin := models.User{"", Words.DeafualtAdmminUserName, Words.DeafualtAdmminFirstName, Words.DeafualtAdmminLastName, Words.DeafualtAdmminEmail, Words.DeafualtAdmminPassword, Words.DeafualtAdmminRole}
		// کاربر ادمیت را به سمت پایگاه داده ارسال میکنیم.
		errCreateUser := UserDatastore.CreateUser(DefaultAdmin, session)
		if errCreateUser != nil && errCreateUser.Error() != Words.UserExist {
			log.SystemErrorHappened(errCreateUser)
			panic(errCreateUser)
		}
		//..................................
		fmt.Println("Default Admin Ok.")
	}

}
func addUser() {}
