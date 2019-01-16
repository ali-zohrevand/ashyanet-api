package services

import (
	"encoding/json"
	"github.com/ali-zohrevand/ashyanet-api/OutputAPI"
	"github.com/ali-zohrevand/ashyanet-api/core/DB"
	"github.com/ali-zohrevand/ashyanet-api/services/log"
	"github.com/ali-zohrevand/ashyanet-api/settings/Words"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

func Active(username string, activeCode string) (statusCode int, message []byte) {
	var validPeroidTempKey int64
	validPeroidTempKey = 24 * 60 * 60 //1 day
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")

	}
	user, errGetUser := DB.UserGetByUsername(username, session)
	if errGetUser != nil {
		message := OutputAPI.Message{}
		message.Error = Words.UserNotExist
		json, _ := json.Marshal(message)
		return http.StatusBadRequest, json
	}
	if user.Active {
		message := OutputAPI.Message{}
		message.Info = Words.UserActivated
		json, _ := json.Marshal(message)
		return http.StatusOK, json
	}
	if user.TempKeyGenreated != activeCode {
		log.SecurityLogHappened(errors.New("bad valid key"))
		return http.StatusUnauthorized, []byte("")
	}
	timeNow := time.Now().Unix()
	delayTime := timeNow - user.TimeTempKeyGenreated
	if delayTime > validPeroidTempKey {
		message := OutputAPI.Message{}
		message.Error = Words.TimeExpired
		json, _ := json.Marshal(message)
		return http.StatusBadRequest, json
	}
	status, errActive := DB.UserActiveBuUsername(user.UserName, session)
	if !status || errActive != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")
	}
	if status {
		message := OutputAPI.Message{}
		message.Info = Words.UserActivated
		json, _ := json.Marshal(message)
		return http.StatusOK, json
	}
	return http.StatusInternalServerError, []byte("")
}
