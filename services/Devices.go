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
	"strings"
)

func CreateDevice(device *models.Device, user models.UserInDB) (int, []byte) {
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
	//.......................We add locattionPath/DeviceType to pubsub slice..........................
	LocationPath, err := DB.GetLocationPath(device.Location, session)
	if err != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")
	}
	LocationAndType := LocationPath + "/" + device.Type
	LocationAndType = strings.Replace(LocationAndType, "//", "/", -1)
	device.Pubsub = append(device.Pubsub, LocationAndType)
	//.................................................
	deviceWithCorrectTopic, err := CheckMqttTopic(device, user)
	if err != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")
	}
	errCreateUser := DB.CreateDevice(*deviceWithCorrectTopic, user, session)
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
func ListDevices() (int, []byte) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")
	}
	defer session.Close()

	return http.StatusInternalServerError, []byte("")

}
