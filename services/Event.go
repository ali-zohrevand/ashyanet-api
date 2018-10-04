package services

import (
	"gitlab.com/hooshyar/ChiChiNi-API/core/DB"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/services/log"
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
		return http.StatusInternalServerError, []byte("")
	}
	CommandName := dataBinde.CommandName
	Comamand, errGetCommand := DB.UserGetMqttCommandByName(user.UserName, CommandName, session)
	if errGetCommand != nil {
		return http.StatusInternalServerError, []byte("")
	}
	Data, errGetDataName := DB.UserGetMqttDataByName(user.UserName, dataBinde.DataName, session)
	if errGetDataName != nil {
		return http.StatusInternalServerError, []byte("")
	}
	if Comamand.Name != "" || Data.Name == "" {
		return http.StatusNotFound, []byte("")
	}
	if !dataBinde.ConditionSet.IsValid() {
		return http.StatusNotAcceptable, []byte("CONDITION IS NOT VALID")

	}
	var event models.Event
	event.EventName = user.UserName + "-" + dataBinde.DataName + "-" + CommandName + "-" + GenerateRandomString(5)
	event.EventFunction = Comamand
	event.EventCondition = dataBinde.ConditionSet
	event.EventAddress = Data.Topic
	event.EventType
	return http.StatusInternalServerError, []byte("")

}
