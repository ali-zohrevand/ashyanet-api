package services

import (
	"fmt"
	"github.com/ali-zohrevand/ashyanet-api/core/DB"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/services/log"
	"github.com/ali-zohrevand/ashyanet-api/settings/ConstKey"
)

func AddInitUser() {
	UserDatastore := DB.UserDataStore{}
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		panic(errConnectDB)
	}
	// چک میکنیم ببینیم آیا قبلا کاربر ادمین ثبت نام شده یا خیر
	user, _ := DB.FindUserByUsername(ConstKey.DeafualtAdmminUserName, session)
	if user.Role != ConstKey.DeafualtAdmminRole {
		//کاربر ادمین را ایجاد می نماییم.
		DefaultAdmin := models.User{"", ConstKey.DeafualtAdmminUserName, ConstKey.DeafualtAdmminFirstName, ConstKey.DeafualtAdmminLastName, ConstKey.DeafualtAdmminEmail, ConstKey.DeafualtAdmminPassword, ConstKey.DeafualtAdmminRole}
		// کاربر ادمیت را به سمت پایگاه داده ارسال میکنیم.
		errCreateUser := UserDatastore.CreateUser(DefaultAdmin, session)
		if errCreateUser != nil && errCreateUser.Error() != ConstKey.UserExist {
			log.SystemErrorHappened(errCreateUser)
			panic(errCreateUser)
		}
		//..................................
		fmt.Println("Default Admin Ok.")
	}

}
func addUser() {}
