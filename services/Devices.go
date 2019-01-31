package services

import (
	"encoding/json"
	"github.com/ali-zohrevand/ashyanet-api/OutputAPI"
	"github.com/ali-zohrevand/ashyanet-api/core/DB"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/services/log"
	"github.com/ali-zohrevand/ashyanet-api/services/validation"
	"github.com/ali-zohrevand/ashyanet-api/settings/Words"
	"net/http"
	"strings"
)

func DeviceCreate(device *models.Device, user models.UserInDB) (int, []byte) {
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
	errCreateUser := DB.DeviceCreate(*deviceWithCorrectTopic, user, session)
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
		if errCreateUser.Error() == Words.DataExist || errCreateUser.Error() == Words.CommandExist {
			message := OutputAPI.Message{}
			message.Error = errCreateUser.Error()
			json, _ := json.Marshal(message)
			return http.StatusBadRequest, json

		}
		log.SystemErrorHappened(errCreateUser)
		return http.StatusInternalServerError, []byte("")

	} else {

		message := OutputAPI.Message{}
		message.Info = Words.DeviceCreated
		json, _ := json.Marshal(device)
		return http.StatusCreated, json
	}
	return http.StatusInternalServerError, nil
}
func DevicesGetAll(user models.UserInDB) (int, []byte) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")
	}
	defer session.Close()
	devices, errGetDevice := DB.DevicesGetAllByUsername(user.UserName, session)
	if errGetDevice == nil {
		jsonOut, errjson := json.Marshal(devices)
		if errjson != nil {
			message := OutputAPI.Message{}
			message.Error = "BAD JSON"
			json, _ := json.Marshal(message)
			return http.StatusInternalServerError, json
		}
		return http.StatusOK, jsonOut
	}
	if errGetDevice != nil {
		message := OutputAPI.Message{}
		message.Error = Words.DeviceNotExist
		json, _ := json.Marshal(message)
		return http.StatusBadRequest, json
	}

	return http.StatusInternalServerError, []byte("")

}
func DeviceGetId(user models.UserInDB, id string) (int, []byte) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")
	}
	defer session.Close()

	devices, errGetDevice := DB.DevicesGetAllByUsernameAndId(user.UserName, id, session)
	if errGetDevice == nil {
		jsonOut, errjson := json.Marshal(devices)
		if errjson != nil {
			message := OutputAPI.Message{}
			message.Error = "BAD JSON"
			json, _ := json.Marshal(message)
			return http.StatusInternalServerError, json
		}
		return http.StatusOK, jsonOut

	}
	if errGetDevice != nil {
		message := OutputAPI.Message{}
		message.Error = Words.DeviceNotExist
		json, _ := json.Marshal(message)
		return http.StatusBadRequest, json
	}

	return http.StatusInternalServerError, []byte("")

}
func DeviceDeleteID(user models.UserInDB, id string) (int, []byte) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")
	}
	defer session.Close()
	errGetDevice := DB.DevicesDeletelByUsernameAndId(user.UserName, id, session)
	if errGetDevice == nil {
		message := OutputAPI.Message{}
		message.Info = Words.DeviceDeleted
		json, _ := json.Marshal(message)
		return http.StatusOK, json

	}
	if errGetDevice != nil {
		message := OutputAPI.Message{}
		message.Error = Words.DeviceNotExist
		json, _ := json.Marshal(message)
		return http.StatusBadRequest, json
	}
	return http.StatusInternalServerError, []byte("")

}
