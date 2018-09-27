package services

import (
	"github.com/eclipse/paho.mqtt.golang"
	"gitlab.com/hooshyar/ChiChiNi-API/core/DB"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/services/log"
	"gitlab.com/hooshyar/ChiChiNi-API/services/validation"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
	"net/http"
)

type mqttStruct struct {
	Ip       string
	UserName string
	Password string
	MqttId   string
	Option   mqtt.ClientOptions
	Client   mqtt.Client
}

func NewMqtt(ip string, username string, passwoard string, ID string) (mqttObj *mqttStruct, err error) {
	//tcp://127.0.0.1:1883
	opts := mqtt.NewClientOptions().AddBroker(ip).SetClientID(ID)
	opts.SetUsername(username)
	opts.SetPassword(passwoard)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return &mqttStruct{ip, username, passwoard, ID, *opts, client}, err
}

func (mqttobj *mqttStruct) Publish(topic string, retaiend bool, payload interface{}, qos byte) (err error) {
	token := mqttobj.Client.Publish(topic, qos, retaiend, payload)

	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return
}

func (mqttobj *mqttStruct) Subscribe(topic string, qos byte, callbackFunction mqtt.MessageHandler) (err error) {
	token := mqttobj.Client.Subscribe(topic, qos, callbackFunction)

	if token.Wait() && token.Error() != nil {
		return
	}
	return
}
func MqttHttpCommand(command models.MqttCommand, User models.UserInDB) (int, []byte) {
	out, errValidation, IsValid := validation.ObjectValidation(command)
	if errValidation != nil || !IsValid {
		return http.StatusBadRequest, out
	}
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")
	}
	defer session.Close()
	Is, errCHeckuser := DB.UserHasMqttCommand(User.UserName, command, session)
	if errCHeckuser != nil {
		return http.StatusInternalServerError, []byte("")
	}
	if !Is {
		return http.StatusUnauthorized, []byte("")
	}
	errPublish := MqttCommandTempAdmin(command)
	if errPublish != nil {
		return http.StatusInternalServerError, []byte("")
	} else {
		return http.StatusOK, []byte("")

	}

	return http.StatusInternalServerError, []byte("")

}

func MqttCommandTempAdmin(command models.MqttCommand) (err error) {
	TempUserAdminUserName, TempAdminPassword, errCreateTempAdmin := EmqttCreateTempAdminMqttUser()
	if errCreateTempAdmin != nil {
		return errCreateTempAdmin
	}
	defer EmqttDeleteUser(TempUserAdminUserName)
	mqttObj, errCreateMqttUser := NewMqtt(Words.MqttBrokerIp, TempUserAdminUserName, TempAdminPassword, "TempAdmin")
	if errCreateMqttUser != nil {
		return errCreateMqttUser
	}
	errPublish := mqttObj.Publish(command.Topic, false, command.Value, 2)
	return errPublish
}
func MqttAddMessage(message models.MqttMessage) (err error) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return
	}
	defer session.Close()
	err = DB.MqttAddMessage(message, session)
	return
}
func MqttGetAllMessageByTopicName(topic string) (MessageList []models.MqttMessage, err error) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return
	}
	MessageList, err = DB.MqttGetMessagesByTopic(topic, session)
	return
}
