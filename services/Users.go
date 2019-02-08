package services

import (
	"encoding/json"
	"fmt"
	"github.com/ali-zohrevand/ashyanet-api/OutputAPI"
	"github.com/ali-zohrevand/ashyanet-api/core/DB"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/services/Tools"
	"github.com/ali-zohrevand/ashyanet-api/services/log"
	"github.com/ali-zohrevand/ashyanet-api/services/validation"
	"github.com/ali-zohrevand/ashyanet-api/settings/Words"
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
		DefaultAdmin := models.User{"", Words.DeafualtAdmminUserName, Words.DeafualtAdmminFirstName, Words.DeafualtAdmminLastName, Words.DeafualtAdmminEmail, Words.DeafualtAdmminPassword, Words.DeafualtAdmminRole, nil, nil, true, "", time.Now().Unix(), nil}
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
		settings, errLoadSettings := DB.LoadSettings(session)
		if errLoadSettings != nil {
			return http.StatusInternalServerError, []byte("")
		}
		if settings.Type != "server" {
			message := OutputAPI.Message{}
			message.Info = Words.UserCreated
			json, _ := json.Marshal(message)
			return http.StatusCreated, json

		}
		// TODO: ممکنه کاربری ساخته بشه ولی میل فرستاده نشده. میبایست امکانی اندیشیده شود.
		userIndb, errFetchUser := DB.UserGetByUsername(requestUser.UserName, session)
		if errFetchUser != nil {
			message := OutputAPI.Message{}
			message.Info = Words.UserCreated + " But " + Words.UserVerifyMailProblem
			json, _ := json.Marshal(message)
			return http.StatusCreated, json
		}
		errSendMail := Verify(userIndb, session)
		if errSendMail != nil {
			message := OutputAPI.Message{}
			message.Info = Words.UserCreated + " But " + Words.UserVerifyMailProblem
			json, _ := json.Marshal(message)
			return http.StatusCreated, json
		} else {
			message := OutputAPI.Message{}
			message.Info = Words.VerifyMailSent + "Your mail is: " + requestUser.Email
			json, _ := json.Marshal(message)
			return http.StatusCreated, json
		}

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
	userInDB, errGetUser := DB.UserGetByUsername(requestUser.UserName, session)
	if errGetUser != nil {
		return http.StatusUnauthorized, []byte("")
	}
	if !userInDB.Active {
		message := OutputAPI.Message{}
		message.Error = Words.UserNotActive
		json, _ := json.Marshal(message)
		return http.StatusNonAuthoritativeInfo, json
	}
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
				Ip:            Tools.GetIpOfRequest(request),
			}
			var jwtSessionDataStore = DB.JwtSessionDataStore{}
			errAddToSessionDB := jwtSessionDataStore.JwtCreate(sessionObj, session)
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
