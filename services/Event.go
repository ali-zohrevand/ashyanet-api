package services

import (
	"encoding/json"
	"github.com/ali-zohrevand/ashyanet-api/OutputAPI"
	"github.com/ali-zohrevand/ashyanet-api/core/DB"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/services/log"
	"github.com/ali-zohrevand/ashyanet-api/settings/Words"
	"net/http"
)

func EventMqttMessageRecived(message models.MqttMessage) (err error) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return errConnectDB
	}
	TopicAddress := message.Topic
	event, errGetAdd := DB.EventGetAddress(TopicAddress, session)
	if errGetAdd != nil {
		return errGetAdd
	}
	switch event.EventType {
	case models.MqttEvent:
		IsHappened, err := event.EventCondition.Happened(message.Message)
		if IsHappened && err == nil {
			errRunCOmmand := MqttCommandTempAdmin(event.EventFunction)
			if errRunCOmmand != nil {
				log.ErrorHappened(errRunCOmmand)

				return errRunCOmmand
			}
		} else {
			log.ErrorHappened(err)
			return err
		}
	case models.SmsEvent:
	default:
	}
	return
}
func EventCreate(dataBinde models.DataBindCommand, user models.UserInDB) (int, []byte) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		message := OutputAPI.Message{}
		message.Error = Words.DBNotConnectet
		json, _ := json.Marshal(message)
		return http.StatusInternalServerError, []byte(json)
	}
	//  در این تابع تمامی کامند هایی که در تمامی دستگاه های کاربر مورد نظر وجود دارد
	//  بررسی خواهد شد و چک میشود کاربر کامندی با نام مورد نظر دارد یا خیر
	Comamand, errGetCommand := DB.UserGetMqttCommandByName(user.UserName, dataBinde.CommandName, session)
	if errGetCommand != nil {
		message := OutputAPI.Message{}
		message.Error = Words.CommandDataNotFOUND
		json, _ := json.Marshal(message)
		return http.StatusInternalServerError, []byte(json)
	}
	//ر این تابع تمامی دیتا هایی که در تمامی دستگاه های کاربر مورد نظر وجود دارد
	//  بررسی خواهد شد و چک میشود کاربر دیتایی با نام مورد نظر دارد یا خیر
	Data, errGetDataName := DB.UserGetMqttDataByName(user.UserName, dataBinde.DataName, session)
	if errGetDataName != nil {
		message := OutputAPI.Message{}
		message.Error = Words.CommandDataNotFOUND
		json, _ := json.Marshal(message)
		return http.StatusInternalServerError, []byte(json)
	}
	// برای بررسی دوباره صحت کامند و دیتا را بررسی می نماییم.
	if Comamand.Name == "" || Data.Name == "" {
		message := OutputAPI.Message{}
		message.Info = Words.CommandDataNotFOUND
		json, _ := json.Marshal(message)
		return http.StatusNotFound, []byte(json)
	}
	// در این مرحله متغیر های کلاس condition‌ را اعتبار سنجی اولیه میکنیم.
	if !dataBinde.ConditionSet.IsValid() {
		message := OutputAPI.Message{}
		message.Error = Words.ConditionIsNotValid
		json, _ := json.Marshal(message)
		return http.StatusNotAcceptable, []byte(json)

	}
	// در نهایت میخواهیم رخ داد (event) را ایجاد و در پایگاه داده ذخیره نماییم.
	var event models.Event
	event.UserOwner = user.UserName
	event.EventName = user.UserName + "-" + dataBinde.DataName + "-" + dataBinde.CommandName + "-" + GenerateRandomString(5)
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
