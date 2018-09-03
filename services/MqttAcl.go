package services

import (
	"encoding/json"
	"gitlab.com/hooshyar/ChiChiNi-API/OutputAPI"
	"gitlab.com/hooshyar/ChiChiNi-API/core/DB"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/services/log"
	"gitlab.com/hooshyar/ChiChiNi-API/services/validation"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/ConstKey"
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
