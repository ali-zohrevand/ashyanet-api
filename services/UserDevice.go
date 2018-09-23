package services

import (
	"encoding/json"
	"gitlab.com/hooshyar/ChiChiNi-API/OutputAPI"
	"gitlab.com/hooshyar/ChiChiNi-API/core/DB"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/services/log"
	"gitlab.com/hooshyar/ChiChiNi-API/services/validation"
	. "gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
	"net/http"
)

func AddUserDevice(userdevice *models.UserDevice) (int, []byte) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")
	}
	defer session.Close()
	out, errValidation, IsValid := validation.ObjectValidation(*userdevice)
	if errValidation != nil || !IsValid {
		return http.StatusBadRequest, out
	}
	err := DB.AddUserDevice(*userdevice, session)
	if err != nil {
		message := OutputAPI.Message{}
		message.Error = DeviceOrUserNotFound
		json, _ := json.Marshal(message)
		return http.StatusBadRequest, json
	} else {
		message := OutputAPI.Message{}
		message.Info = UserAddedToDevice
		json, _ := json.Marshal(message)
		return http.StatusOK, json
	}

	return http.StatusInternalServerError, nil
}
