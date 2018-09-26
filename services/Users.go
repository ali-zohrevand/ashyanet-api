package services

import (
	"encoding/json"
	"fmt"
	"gitlab.com/hooshyar/ChiChiNi-API/OutputAPI"
	"gitlab.com/hooshyar/ChiChiNi-API/core/DB"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/services/log"
	"gitlab.com/hooshyar/ChiChiNi-API/services/validation"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

func AddInitUser() {
	UserDatastore := DB.UserDataStore{}
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		panic(errConnectDB)
	}
	// چک میکنیم ببینیم آیا قبلا کاربر ادمین ثبت نام شده یا خیر
	user, _ := DB.UserGetByUsername(Words.DeafualtAdmminUserName, session)
	if user.Role != Words.DeafualtAdmminRole {
		//کاربر ادمین را ایجاد می نماییم.
		DefaultAdmin := models.User{"", Words.DeafualtAdmminUserName, Words.DeafualtAdmminFirstName, Words.DeafualtAdmminLastName, Words.DeafualtAdmminEmail, Words.DeafualtAdmminPassword, Words.DeafualtAdmminRole, nil, nil}
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
func Register(requestUser *models.User) (int, []byte) {
	out, errValidation, IsValid := validation.ObjectValidation(*requestUser)
	if errValidation != nil || !IsValid {
		return http.StatusBadRequest, out
	}
	UserDatastore := DB.UserDataStore{}
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")
	}
	defer session.Close()
	//................................. Main Action is Hear .................................
	errCreateUser := UserDatastore.CreateUser(*requestUser, session)

	if errCreateUser != nil {
		if errCreateUser.Error() == Words.UserExist {
			//User Exist
			message := OutputAPI.Message{}
			message.Error = Words.UserExist
			json, _ := json.Marshal(message)
			return http.StatusOK, json
		}
		log.SystemErrorHappened(errCreateUser)
		return http.StatusInternalServerError, []byte("")

	} else {

		message := OutputAPI.Message{}
		message.Info = Words.UserCreated
		json, _ := json.Marshal(message)
		return http.StatusCreated, json
	}
	//......................................................................................
	return http.StatusInternalServerError, []byte("")
}
func Login(requestUser *models.User, request *http.Request) (int, []byte) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")
	}
	defer session.Close()
	out, errValidation, IsValid := validation.UserLoginValidation(*requestUser)
	if errValidation != nil || !IsValid {
		return http.StatusBadRequest, out
	}
	UserDatastore := DB.UserDataStore{}
	session, errConDB := DB.ConnectDB()
	if errConDB != nil {
		log.SystemErrorHappened(errConDB)
		return http.StatusUnauthorized, []byte("")
	}
	defer session.Close()
	if UserDatastore.CheckUserPassCorrect(*requestUser, session) {
		token, err := GenerateToken(requestUser.UserName)
		if err != nil {
			log.SystemErrorHappened(err)
			return http.StatusUnauthorized, []byte("")
		} else {
			//................................. Main Action is Hear .................................
			response, err := json.Marshal(OutputAPI.TokenAuthentication{token})
			if err != nil {
				log.SystemErrorHappened(err)
				return http.StatusUnauthorized, []byte("")
			}
			sessionObj := models.JwtSession{
				Id:            bson.NewObjectId(),
				JwtToken:      token,
				OwnerUsername: requestUser.UserName,
				TimeCreated:   time.Now(),
				Ip:            GetIpOfRequest(request),
			}
			var jwtSessionDataStore = DB.JwtSessionDataStore{}
			errAddToSessionDB := jwtSessionDataStore.CreateJwtSession(sessionObj, session)
			if errAddToSessionDB != nil {
				log.SystemErrorHappened(errConnectDB)
				return http.StatusInternalServerError, []byte("")
			}
			return http.StatusOK, response
			//..................................................................
		}
	} else {
		//todo: security log Happened beacuse user and password not match
	}
	return http.StatusUnauthorized, []byte("")
}
