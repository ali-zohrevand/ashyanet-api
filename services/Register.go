package services

import (
	"encoding/json"
	"github.com/ali-zohrevand/ashyanet-api/OutputAPI"
	"github.com/ali-zohrevand/ashyanet-api/core/DB"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/services/log"
	"github.com/ali-zohrevand/ashyanet-api/services/validation"
	"github.com/ali-zohrevand/ashyanet-api/settings/ConstKey"
	"net/http"
)

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
		if errCreateUser.Error() == ConstKey.UserExist {
			//User Exist
			message := OutputAPI.Message{}
			message.Error = ConstKey.UserExist
			json, _ := json.Marshal(message)
			return http.StatusOK, json
		}
		log.SystemErrorHappened(errCreateUser)
		return http.StatusInternalServerError, []byte("")

	} else {

		message := OutputAPI.Message{}
		message.Info = ConstKey.UserCreated
		json, _ := json.Marshal(message)
		return http.StatusCreated, json
	}
	//......................................................................................
	return http.StatusInternalServerError, []byte("")
}
