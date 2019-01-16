package services

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/ali-zohrevand/ashyanet-api/core/DB"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/services/log"
	"github.com/ali-zohrevand/ashyanet-api/services/validation"
	"github.com/ali-zohrevand/ashyanet-api/settings/Words"
	"github.com/eclipse/paho.mqtt.golang"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

/*type mqttStruct struct {
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

*/
type mqttStruct struct {
	Ip       string
	UserName string
	Password string
	MqttId   string
	Option   mqtt.ClientOptions
	Client   mqtt.Client
}

func NewTlsConfig() *tls.Config {
	// Import trusted certificates from CAfile.pem.
	// Alternatively, manually add CA certificates to
	// default openssl CA bundle.
	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile("certs/cacert.pem")
	if err == nil {
		certpool.AppendCertsFromPEM(pemCerts)
	}

	// Import client certificate/key pair
	cert, err := tls.LoadX509KeyPair("certs/client-cert.pem", "certs/client-key.pem")
	if err != nil {
		panic(err)
	}

	// Just to print out the client certificate..
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		panic(err)
	}
	fmt.Println(cert.Leaf)

	// Create tls.Config with desired tls properties
	return &tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs: certpool,
		// ClientAuth = whether to request cert from server.
		// Since the server is set up for SSL, this happens
		// anyways.
		ClientAuth: tls.NoClientCert,
		// ClientCAs = certs used to validate client cert.
		ClientCAs: nil,
		// InsecureSkipVerify = verify that cert contents
		// match server. IP matches what is in cert etc.
		InsecureSkipVerify: true,
		// Certificates = list of certs client sends to server.
		Certificates: []tls.Certificate{cert},
	}
}
func NewMqttWithTLS(ip string, username string, passwoard string, ID string) (mqttObj *mqttStruct, err error) {
	tlsconfig := NewTlsConfig()
	opts := mqtt.NewClientOptions()
	if opts == nil {
		return nil, errors.New("ERROR")
	}
	ip = "ssl://" + ip + ":8883"
	opts = opts.AddBroker(ip)
	if opts == nil {
		return nil, errors.New("ERROR")
	}
	opts = opts.SetTLSConfig(tlsconfig)
	if opts == nil {
		return nil, errors.New("ERROR")
	}
	opts.SetUsername(username)
	opts.SetPassword(passwoard)

	client := mqtt.NewClient(opts)
	if client == nil {
		return nil, errors.New("Error in create")
	}

	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {

	}
	if token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return &mqttStruct{ip, username, passwoard, ID, *opts, client}, err

}
func NewMqtt(ip string, username string, passwoard string, ID string) (mqttObj *mqttStruct, err error) {
	//tcp://127.0.0.1:1883
	ip = "tcp://" + ip + ":1883"
	opts := mqtt.NewClientOptions().AddBroker(ip).SetClientID(ID)

	opts.SetUsername(username)
	opts.SetPassword(passwoard)

	client := mqtt.NewClient(opts)
	if client == nil {
		return nil, errors.New("Error in create")
	}

	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {

	}
	if token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return &mqttStruct{ip, username, passwoard, ID, *opts, client}, err
}

func (mqttobj *mqttStruct) Publish(topic string, retaiend bool, payload interface{}, qos byte) (err error) {
	var mutex = &sync.Mutex{}
	if mqttobj == nil || mqttobj.Client == nil {
		return errors.New("Error in create")
	}

	mutex.Lock()

	token := mqttobj.Client.Publish(topic, qos, retaiend, payload)
	mutex.Unlock()

	for !token.WaitTimeout(3 * time.Second) {

	}
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
func MqttHttpCommand(command models.Command, User models.UserInDB) (int, []byte) {
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
		return http.StatusOK, []byte("sent")

	}

	return http.StatusInternalServerError, []byte("")

}

func MqttCommandTempAdmin(command models.Command) (err error) {
	EmqttDeleteMqttDefaultAdmin()
	TempUserAdminUserName, TempAdminPassword, errCreateTempAdmin := EmqttCreateTempAdminMqttUserWithDwefaultAdmin()
	if errCreateTempAdmin != nil {
		return errCreateTempAdmin
	}
	/*	TempUserAdminUserName, TempAdminPassword, errCreateTempAdmin := EmqttCreateTempAdminMqttUser()
		if errCreateTempAdmin != nil {
			return errCreateTempAdmin
		}*/
	defer EmqttDeleteUser(TempUserAdminUserName)
	mqttObj, errCreateMqttUser := NewMqtt(Words.MqttBrokerIp, TempUserAdminUserName, TempAdminPassword, "TempAdmin")
	defer mqttObj.Client.Disconnect(50)
	if errCreateMqttUser != nil {
		return errCreateMqttUser
	}
	errPublish := mqttObj.Publish(command.Topic, false, command.Value, 2)
	mqttObj.Client.Disconnect(50)
	return errPublish
}
func MqttAddMessageToDb(message models.MqttMessage) (err error) {
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
	MessageList, err = DB.MqttGetAllMessagesByTopic(topic, session)
	return
}

func MqttSubcribeRootTopic() (err error) {
	EmqttDeleteMqttDefaultAdmin()
	TempUserAdminUserName, TempAdminPassword, errCreateTempAdmin := EmqttCreateTempAdminMqttUserWithDwefaultAdmin()
	if errCreateTempAdmin != nil {
		panic(err)
		return errCreateTempAdmin
	}
	defer EmqttDeleteUser(TempUserAdminUserName)
	done := make(chan bool)
	mqttObj, errCreateMqttUser := NewMqtt(Words.MqttBrokerIp, TempUserAdminUserName, TempAdminPassword, "TempAdmin"+GenerateRandomString(3))
	defer mqttObj.Client.Disconnect(50)
	if errCreateMqttUser != nil {
		panic(errCreateMqttUser)
		return errCreateMqttUser
	}
	var eventFunc mqtt.MessageHandler = func(client mqtt.Client, message mqtt.Message) {
		var mqttmeesage models.MqttMessage
		mqttmeesage.Message = string(message.Payload())
		mqttmeesage.Topic = message.Topic()
		mqttmeesage.Time = time.Now().String()
		mqttmeesage.Retained = message.Retained()
		mqttmeesage.Qos = message.Qos()
		mqttmeesage.MessageId = string(message.MessageID())
		errAddMessage := MqttAddMessageToDb(mqttmeesage)
		if errAddMessage != nil {
			log.ErrorHappened(errAddMessage)
		}
		errEventRegister := EventMqttMessageRecived(mqttmeesage)
		if errEventRegister != nil {
			log.ErrorHappened(errAddMessage)
		}
	}
	errSubscribe := mqttObj.Subscribe("#", 0, eventFunc)
	if errSubscribe != nil {
		panic(err)
		return errSubscribe
	}
	<-done
	return
}
