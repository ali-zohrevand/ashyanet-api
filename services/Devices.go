package services

import (
	"encoding/json"
	"gitlab.com/hooshyar/ChiChiNi-API/OutputAPI"
	"gitlab.com/hooshyar/ChiChiNi-API/core/DB"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/services/log"
	"gitlab.com/hooshyar/ChiChiNi-API/services/validation"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
	"net/http"
)

func CreateDevice(device *models.Device) (int, []byte) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")
	}
	defer session.Close()
	out, errValidation, IsValid := validation.ObjectValidation(*device)
	if errValidation != nil || !IsValid {
		return http.StatusBadRequest, out
	}
	errCreateUser := DB.CreateDevice(*device, session)
	if errCreateUser != nil {
		if errCreateUser.Error() == Words.DeviceExist {
			//User Exist
			message := OutputAPI.Message{}
			message.Error = Words.DeviceExist
			json, _ := json.Marshal(message)
			return http.StatusOK, json
		}
		if errCreateUser.Error() == Words.KeyIsNotValid {
			message := OutputAPI.Message{}
			message.Error = Words.KeyIsNotValid
			json, _ := json.Marshal(message)
			return http.StatusNotFound, json
		}
		log.SystemErrorHappened(errCreateUser)
		return http.StatusInternalServerError, []byte("")

	} else {

		message := OutputAPI.Message{}
		message.Info = Words.DeviceCreated
		json, _ := json.Marshal(message)
		return http.StatusCreated, json
	}
	return http.StatusInternalServerError, nil
}
func AddUserToDevice(userdevice *models.UserDevice) (int, []byte) {
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
	err := DB.AddUserToDevice(*userdevice, session)
	if err != nil {
		message := OutputAPI.Message{}
		message.Error = Words.DeviceOrUserNotFound
		json, _ := json.Marshal(message)
		return http.StatusBadRequest, json
	} else {
		message := OutputAPI.Message{}
		message.Info = Words.UserAddedToDevice
		json, _ := json.Marshal(message)
		return http.StatusOK, json
	}

	return http.StatusInternalServerError, nil
}
func ListDevices() (int, []byte) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")
	}
	defer session.Close()

	return http.StatusInternalServerError, []byte("")

}
