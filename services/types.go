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

func TypeCreate(typesObj models.Types, user models.UserInDB) (int, []byte) {
	out, errValidation, IsValid := validation.ObjectValidation(typesObj)
	if errValidation != nil || !IsValid {
		return http.StatusBadRequest, out
	}
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")

	}
	typesObj.Owner = user.UserName
	typeName := typesObj.Name
	typeName = strings.Replace(typeName, " ", "-", -1)
	typesObj.Name = user.UserName + "-" + typeName
	err := DB.TypesCreate(typesObj, session)
	if err == nil {
		///..........................................
		errAddUserTypes := DB.UserAddTypes(user, typesObj, session)
		if errAddUserTypes != nil {
			DB.TypeDeleteByName(typesObj.Name, session)
			return http.StatusInternalServerError, []byte("")
		}
		///..........................................
		var message OutputAPI.Message
		message.Info = Words.TypeCreated
		messageJson, _ := json.Marshal(message)
		return http.StatusCreated, messageJson
	}
	if err.Error() == Words.TypeExits {
		var message OutputAPI.Message
		message.Error = Words.TypeExits
		messageJson, _ := json.Marshal(message)
		return http.StatusOK, messageJson

	}

	return http.StatusInternalServerError, []byte("")
}
func TypeGetAllByUsername(user models.UserInDB) (int, []byte) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")

	}
	types, err := DB.TypesGetAllTypesOfUser(user.UserName, session)
	if types == nil {
		var message OutputAPI.Message
		message.Error = Words.TypeNotExit
		messageJson, _ := json.Marshal(message)
		return http.StatusOK, messageJson
	}
	if err == nil {
		typesJson, errJson := json.Marshal(types)
		if errJson != nil {
			return http.StatusInternalServerError, []byte("")

		}
		return http.StatusOK, typesJson
	}

	return http.StatusInternalServerError, []byte("")
}
func TypesDeleteById(id string, user models.UserInDB) (int, []byte) {

	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")

	}
	defer session.Clone()
	typeObj, err := DB.TypeGetTypeByID(id, session)
	if err != nil {
		var message OutputAPI.Message
		message.Error = Words.TypeNotExit
		messageJson, _ := json.Marshal(message)
		return http.StatusNotFound, messageJson

	}
	hasUserTypes := false
	index := 0
	for i, typeInArray := range user.Types {
		if typeInArray == typeObj.Name {
			hasUserTypes = true
			index = i
		}
	}
	if !hasUserTypes {
		return http.StatusNotFound, []byte("")
	}
	errDelete := DB.TypeDeleteByName(typeObj.Name, session)
	if errDelete == nil {
		user.Types = append(user.Types[:index], user.Types[index+1:]...)
		errUpdateuser := DB.UserUpdateByUserObj(user, session)
		if errUpdateuser != nil {
			log.ErrorHappened(errUpdateuser)
			return http.StatusInternalServerError, []byte("")
		}
		var message OutputAPI.Message
		message.Error = Words.TypeDeleted
		messageJson, _ := json.Marshal(message)
		return http.StatusOK, messageJson
	}
	return http.StatusInternalServerError, []byte("")

}
