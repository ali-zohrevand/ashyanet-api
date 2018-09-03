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

func CreateDefaultKey() {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		panic(errConnectDB)
	}
	defer session.Close()
	for i := 0; i < 10; i++ {
		DB.CreateDeviceKey(session)
	}
}
func ListKey() (key []string) {
	return
}
func IsKeyValid(key string) (isValid bool) {
	session, errConnectDB := DB.ConnectDB()
	defer session.Close()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
	}
	isValid = DB.CheckKeyIsValid(key, session)
	return isValid
}
func CreatValidKey() (int, []byte) {
	session, errConnectDB := DB.ConnectDB()
	defer session.Close()
	if errConnectDB != nil {
		return http.StatusInternalServerError, []byte("")
		log.SystemErrorHappened(errConnectDB)
	}
	key := DB.GetValidKey(session)
	if len(key.Key) != Words.LengthOfDeviceKey {
		return http.StatusInternalServerError, []byte("")

	} else {
		message := OutputAPI.Key{}
		message.Key = key.Key
		json, _ := json.Marshal(message)
		return http.StatusOK, json
	}
	return http.StatusInternalServerError, []byte("")

}
func AddKeyToDevice(deviceKey *models.DeviceKey) (int, []byte) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")
	}
	defer session.Close()
	out, errValidation, IsValid := validation.ObjectValidation(*deviceKey)
	if errValidation != nil || !IsValid {
		return http.StatusBadRequest, out
	}
	err := DB.AddKeyToDevice(*deviceKey, session)
	if err != nil {
		message := OutputAPI.Message{}
		message.Error = Words.KeyIsNotValid
		json, _ := json.Marshal(message)
		return http.StatusBadRequest, json
	} else {
		message := OutputAPI.Message{}
		message.Info = Words.KeyAddedTodevice
		json, _ := json.Marshal(message)
		return http.StatusOK, json
	}

	return http.StatusInternalServerError, nil
}
