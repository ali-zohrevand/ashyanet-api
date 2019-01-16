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

func CreateLocation(Location *models.Location) (int, []byte) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		panic(errConnectDB)

	}
	defer session.Close()
	out, errValidation, IsValid := validation.ObjectValidation(*Location)
	if errValidation != nil || !IsValid {
		return http.StatusBadRequest, out
	}
	errCreateLocation := DB.CreateLocation(*Location, session)
	if errCreateLocation != nil {
		if errCreateLocation.Error() == ConstKey.LocationExist {
			//User Exist
			message := OutputAPI.Message{}
			message.Error = ConstKey.LocationExist
			json, _ := json.Marshal(message)
			return http.StatusOK, json
		}
		if errCreateLocation.Error() == ConstKey.LocationNotFound {
			message := OutputAPI.Message{}
			message.Error = ConstKey.LocationNotFound
			json, _ := json.Marshal(message)
			return http.StatusOK, json
		}
		log.SystemErrorHappened(errCreateLocation)
		return http.StatusInternalServerError, []byte("")
	} else {
		message := OutputAPI.Message{}
		message.Info = ConstKey.LocationCreated
		json, _ := json.Marshal(message)
		return http.StatusCreated, json
	}

	return http.StatusInternalServerError, nil
}
