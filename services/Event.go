package services

import (
	"encoding/json"
	"gitlab.com/hooshyar/ChiChiNi-API/OutputAPI"
	"gitlab.com/hooshyar/ChiChiNi-API/core/DB"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/services/log"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
	"net/http"
)

func EventMqttMessageRecived(message models.MqttMessage) (err error) {
	//TopicAddress:=message.Topic

	return
}
func EventCreate(dataBinde models.DataBind, user models.UserInDB) (int, []byte) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		message := OutputAPI.Message{}
		message.Error = Words.DBNotConnectet
		json, _ := json.Marshal(message)
		return http.StatusInternalServerError, []byte(json)
	}
	CommandName := dataBinde.CommandName
	Comamand, errGetCommand := DB.UserGetMqttCommandByName(user.UserName, CommandName, session)
	if errGetCommand != nil {
		message := OutputAPI.Message{}
		message.Error = Words.DBNotConnectet
		json, _ := json.Marshal(message)
		return http.StatusInternalServerError, []byte(json)
	}
	Data, errGetDataName := DB.UserGetMqttDataByName(user.UserName, dataBinde.DataName, session)
	if errGetDataName != nil {
		message := OutputAPI.Message{}
		message.Error = Words.DBNotConnectet
		json, _ := json.Marshal(message)
		return http.StatusInternalServerError, []byte(json)
	}
	if Comamand.Name != "" || Data.Name == "" {
		message := OutputAPI.Message{}
		message.Info = Words.CommandDataNotFOUND
		json, _ := json.Marshal(message)
		return http.StatusNotFound, []byte(json)
	}
	if !dataBinde.ConditionSet.IsValid() {
		message := OutputAPI.Message{}
		message.Error = Words.ConditionIsNotValid
		json, _ := json.Marshal(message)
		return http.StatusNotAcceptable, []byte(json)

	}
	var event models.Event
	event.EventName = user.UserName + "-" + dataBinde.DataName + "-" + CommandName + "-" + GenerateRandomString(5)
	event.EventFunction = Comamand
	event.EventCondition = dataBinde.ConditionSet
	event.EventAddress = Data.Topic
	event.EventType = dataBinde.ComandType
	errCreatedEventDB := DB.EventCreate(event, session)
	if errCreatedEventDB != nil {
		message := OutputAPI.Message{}
		message.Error = Words.DBNotConnectet
		json, _ := json.Marshal(message)
		log.ErrorHappened(errCreatedEventDB)
		return http.StatusInternalServerError, []byte(json)
	} else {
		message := OutputAPI.Message{}
		message.Info = Words.EventCreated
		json, _ := json.Marshal(message)
		return http.StatusCreated, json
	}
	return http.StatusInternalServerError, []byte("")

}
