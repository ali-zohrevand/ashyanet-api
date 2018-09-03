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
		if errCreateLocation.Error() == Words.LocationExist {
			//User Exist
			message := OutputAPI.Message{}
			message.Error = Words.LocationExist
			json, _ := json.Marshal(message)
			return http.StatusOK, json
		}
		if errCreateLocation.Error() == Words.LocationNotFound {
			message := OutputAPI.Message{}
			message.Error = Words.LocationNotFound
			json, _ := json.Marshal(message)
			return http.StatusOK, json
		}
		log.SystemErrorHappened(errCreateLocation)
		return http.StatusInternalServerError, []byte("")
	} else {
		message := OutputAPI.Message{}
		message.Info = Words.LocationCreated
		json, _ := json.Marshal(message)
		return http.StatusCreated, json
	}

	return http.StatusInternalServerError, nil
}
