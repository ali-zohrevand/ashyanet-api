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

func AddUserLocation(userlocation *models.UserLocation) (int, []byte) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")
	}
	defer session.Close()
	out, errValidation, IsValid := validation.ObjectValidation(*userlocation)
	if errValidation != nil || !IsValid {
		return http.StatusBadRequest, out
	}
	err := DB.AddUserLocation(*userlocation, session)
	if err != nil {
		message := OutputAPI.Message{}
		message.Error = UserOrLocationNotFound
		json, _ := json.Marshal(message)
		return http.StatusBadRequest, json
	} else {
		message := OutputAPI.Message{}
		message.Info = UserAddedToLocation
		json, _ := json.Marshal(message)
		return http.StatusOK, json
	}

	return http.StatusInternalServerError, nil
}
