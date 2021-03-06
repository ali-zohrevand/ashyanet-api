package DB

import (
	"encoding/json"
	"fmt"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"testing"
)

func TestDeviceCreate(t *testing.T) {
	session, errConnectDB := ConnectDB()
	if errConnectDB != nil {
		t.Fail()
	}
	defer session.Close()
	var D models.Device
	var MqttCommand models.Command
	var MqttData models.Data
	MqttCommand.Topic = "/test"
	MqttCommand.Dsc = "Test"
	MqttCommand.Name = "on"
	MqttCommand.Value = "ON"
	MqttData.Name = "OF"
	MqttData.Dsc = "tst"
	MqttData.ValueType = "int"
	D.MqttCommand = append(D.MqttCommand, MqttCommand)
	D.MqttData = append(D.MqttData, MqttData)
	D.MqttPassword = "mahdi1369"
	D.Name = "alialiali"
	D.Description = "test device"
	D.Location = "room"
	D.Owners = append(D.Owners, "admin")
	D.Publish = append(D.Publish, "test")
	D.Subscribe = append(D.Publish, "test")
	D.Pubsub = append(D.Publish, "test")
	D.Type = "light"
	k, er := GetValidKey(session)
	if er != nil {
		t.Error(er)
		t.Fail()
	}
	D.Key = k.Key
	var u models.UserInDB
	e := DeviceCreate(D, u, session)
	fmt.Print(e)
	if e != nil {
		t.Error(e)
		t.Fail()
	}

}

/*func TestCreateDevice(t *testing.T) {
session, errConnectDB := DB.ConnectDB()
if errConnectDB != nil {
t.Fail()
}
defer session.Close()
ValidUser:=models.Device{"","test_"+string(rand.Intn(100)),"dsc","light","jhjdhfjskdfhjksdf",nil}
var tests = []struct {
	input    models.Device
	expected error
}{
	{ValidUser,nil },

}
for _, test := range tests {
	if output := DeviceCreate(test.input,session); output != test.expected {
		t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		//t.Fail()
	}
}

}*/
/*func TestCheckExist(t *testing.T) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		t.Fail()
	}
	defer session.Close()
	ValidUser:=models.Device{"","test_"+string(rand.Intn(100)),"dsc","light","jhjdhfjskdfhjksdf",nil}
	errExisted:=errors.New(Words.DeviceExist)
	var tests = []struct {
		input    models.Device
		expected error
	}{
	{ValidUser,errExisted},

	}
	for _, test := range tests {
		if output := DeviceCreate(test.input,session); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
			//t.Fail()
		}
	}
}*/
func TestDeviceGetAllTopic(t *testing.T) {
}
func TestDevicesGetAllByUsername(t *testing.T) {
	session, errConnectDB := ConnectDB()
	if errConnectDB != nil {
		t.Fail()
	}
	username := "user6"
	Devices, err := DevicesGetAllByUsername(username, session)
	if err != nil {
		t.Fail()
	}
	DeviceJson, _ := json.Marshal(Devices)
	fmt.Println(len(Devices))
	fmt.Println(string(DeviceJson))
}
func TestDeviceGetById(t *testing.T) {
	session, errConnectDB := ConnectDB()
	if errConnectDB != nil {
		t.Fail()
	}
	device, _ := DeviceGetById("5c3c3da31d41c82efb985cef", session)
	fmt.Println(device)
}
