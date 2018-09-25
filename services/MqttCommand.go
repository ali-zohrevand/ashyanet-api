package services

import (
	"gitlab.com/hooshyar/ChiChiNi-API/core/DB"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/services/log"
	"gitlab.com/hooshyar/ChiChiNi-API/services/validation"
	"net/http"
)

func MqttCommand(command models.MqttCommand, User models.UserInDB) (int, []byte) {
	out, errValidation, IsValid := validation.ObjectValidation(command)
	if errValidation != nil || !IsValid {
		return http.StatusBadRequest, out
	}
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")
	}
	defer session.Close()
	Is, errCHeckuser := DB.UserHasMqttCommand(User.UserName, command, session)
	if errCHeckuser != nil {
		return http.StatusInternalServerError, []byte("")
	}
	if !Is {
		return http.StatusUnauthorized, []byte("")
	}
	return http.StatusInternalServerError, []byte("")

}
