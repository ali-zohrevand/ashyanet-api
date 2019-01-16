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

func CreateMqttAcl(acl *models.MqttAcl) (int, []byte) {

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
	errCreateAcl := DB.CreateMqattAcl(*acl, session)
	if errCreateAcl != nil {
		switch errCreateAcl.Error() {
		case ConstKey.UserExist:
			message := OutputAPI.Message{}
			message.Error = ConstKey.UserExist
			json, _ := json.Marshal(message)
			return http.StatusOK, json
		case ConstKey.MqttUserNotFound:
			message := OutputAPI.Message{}
			message.Error = ConstKey.MqttUserNotFound
			json, _ := json.Marshal(message)
			return http.StatusOK, json
		default:
			log.SystemErrorHappened(errCreateAcl)
			return http.StatusInternalServerError, []byte("")
		}

	} else {
		message := OutputAPI.Message{}
		message.Info = ConstKey.AclCreated
		json, _ := json.Marshal(message)
		return http.StatusCreated, json
	}

	return http.StatusInternalServerError, []byte("")
}
