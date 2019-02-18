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
)

func CreateLocation(Location *models.Location, user models.UserInDB) (int, []byte) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)

	}
	defer session.Close()
	out, errValidation, IsValid := validation.ObjectValidation(*Location)
	if errValidation != nil || !IsValid {
		return http.StatusBadRequest, out
	}
	errCreateLocation := DB.LocationCreateWithUserName(*Location, user.UserName, session)
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

func LocationGetAllByUsername(user models.UserInDB) (int, []byte) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")

	}
	locations, err := DB.LocationGetByUsername(user.UserName, session)
	if locations == nil {
		var message OutputAPI.Message
		message.Error = Words.LocationNotFound
		messageJson, _ := json.Marshal(message)
		return http.StatusOK, messageJson
	}
	if err == nil {
		locationJson, errJson := json.Marshal(locations)
		if errJson != nil {
			return http.StatusInternalServerError, []byte("")

		}
		return http.StatusOK, locationJson
	}

	return http.StatusInternalServerError, []byte("")

}
func LocationDeleteById(id string, user models.UserInDB) (int, []byte) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")

	}
	defer session.Clone()
	location, err := DB.LocationGetById(id, session)
	var locationBackUp models.Location
	locationBackUp.Name = location.Name
	locationBackUp.Devices = location.Devices
	locationBackUp.Description = location.Description
	locationBackUp.Latitude = location.Latitude
	locationBackUp.Longitude = location.Longitude
	locationBackUp.Parent = location.Parent
	locationBackUp.Users = location.Users
	locationBackUp.DisplayName = location.DisplayName
	if err != nil {
		var message OutputAPI.Message
		message.Error = Words.LocationNotFound
		messageJson, _ := json.Marshal(message)
		return http.StatusNotFound, messageJson

	}
	hasUserLocation := false
	index := 0
	for i, locationInArray := range user.Locations {
		if locationInArray == location.Name {
			hasUserLocation = true
			index = i
		}
	}
	if !hasUserLocation {
		return http.StatusNotFound, []byte("")
	}
	errDelete := DB.LocationDeleteByName(location.Name, session)
	if errDelete == nil {
		user.Locations = append(user.Locations[:index], user.Locations[index+1:]...)
		errUpdateuser := DB.UserUpdateByUserObj(user, session)
		if errUpdateuser != nil {
			log.ErrorHappened(errUpdateuser)
			DB.LocationCreate(locationBackUp, session)
			return http.StatusInternalServerError, []byte("")
		}
		var message OutputAPI.Message
		message.Error = Words.LocationDeleted
		messageJson, _ := json.Marshal(message)
		return http.StatusOK, messageJson
	}
	return http.StatusInternalServerError, []byte("")

}
