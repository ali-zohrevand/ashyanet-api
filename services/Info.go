package services

import (
	"encoding/json"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"net/http"
)

func Info(user models.UserInDB) (int, []byte) {
	settings, err := GetServerSettings()
	if err != nil {
		return http.StatusNotFound, []byte("")

	}
	var info models.Info
	info.Username = user.UserName
	info.Name = user.FirstName
	info.TopicRootPath = settings.Identifier + "/" + user.UserName
	infoJson, err := json.Marshal(info)
	if err != nil {
		return http.StatusInternalServerError, []byte("")

	}

	return http.StatusOK, infoJson

}
