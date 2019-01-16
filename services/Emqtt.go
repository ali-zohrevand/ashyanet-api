package services

import (
	"encoding/json"
	"errors"
	"github.com/ali-zohrevand/ashyanet-api/OutputAPI"
	"github.com/ali-zohrevand/ashyanet-api/core/DB"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/services/log"
	"github.com/ali-zohrevand/ashyanet-api/services/validation"
	"github.com/ali-zohrevand/ashyanet-api/settings/Words"
	"net/http"
	"time"
)

func EmqttHttpCreateAcl(acl *models.MqttAcl) (int, []byte) {

	out, errValidation, IsValid := validation.ObjectValidation(*acl)
	if errValidation != nil || !IsValid {
		return http.StatusBadRequest, out
	}
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")
	}
	defer session.Close()
	//..................................................................
	errCreateAcl := DB.EmqttCreateAcl(*acl, session)
	if errCreateAcl != nil {
		switch errCreateAcl.Error() {
		case Words.UserExist:
			message := OutputAPI.Message{}
			message.Error = Words.UserExist
			json, _ := json.Marshal(message)
			return http.StatusOK, json
		case Words.MqttUserNotFound:
			message := OutputAPI.Message{}
			message.Error = Words.MqttUserNotFound
			json, _ := json.Marshal(message)
			return http.StatusOK, json
		default:
			log.SystemErrorHappened(errCreateAcl)
			return http.StatusInternalServerError, []byte("")
		}

	} else {
		message := OutputAPI.Message{}
		message.Info = Words.AclCreated
		json, _ := json.Marshal(message)
		return http.StatusCreated, json
	}

	return http.StatusInternalServerError, []byte("")
}
func EmqttCreateTempAdminMqttUserWithDwefaultAdmin() (UserName string, Passwoard string, err error) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return "", "", errors.New("DB IS NOT OK")
	}
	var user models.MqttUser
	user.Username = Words.MqttDefaultAdmin
	//user.Username="admin-"+GenerateRandomString(9)
	user.Password = GenerateRandomString(16)
	user.Is_superuser = true
	user.Created = time.Now().String()
	errCreateUser := DB.EmqttCreateUser(user, session)
	if errCreateUser != nil {
		return "", "", errCreateUser

	}
	UserName = user.Username
	Passwoard = user.Password
	return
}
func EmqttDeleteMqttDefaultAdmin() (err error) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return errors.New("DB IS NOT OK")
	}
	err = DB.EmqttDeleteUser(Words.MqttDefaultAdmin, session)
	return
}
func EmqttCreateTempAdminMqttUser() (UserName string, Passwoard string, err error) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return "", "", errors.New("DB IS NOT OK")
	}
	var user models.MqttUser
	user.Username = GenerateRandomString(5)
	user.Password = GenerateRandomString(16)
	user.Is_superuser = true
	user.Created = time.Now().String()
	errCreateUser := DB.EmqttCreateUser(user, session)
	if errCreateUser != nil {
		return "", "", errCreateUser

	}
	UserName = user.Username
	Passwoard = user.Password
	return
}
func EmqttDeleteUser(userName string) (err error) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return errors.New("DB IS NOT OK")
	}
	err = DB.EmqttDeleteUser(userName, session)
	return
}
func EmqttHttpCreatUser(user *models.MqttUser) (int, []byte) {

	out, errValidation, IsValid := validation.ObjectValidation(*user)
	if errValidation != nil || !IsValid {
		return http.StatusBadRequest, out
	}
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")
	}
	defer session.Close()
	//.....................................
	errCreateUser := DB.EmqttCreateUser(*user, session)
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
	return http.StatusInternalServerError, []byte("")
}
