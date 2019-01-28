package DB

import (
	"fmt"
	"github.com/ali-zohrevand/ashyanet-api/Lab/mosquitto/03_crateTopicBaseLocation/DB"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/settings/Words"
	"testing"
)

func TestCreateDeviceKey(t *testing.T) {
	var devicekey models.DeviceKey
	devicekey.Key = GenerateKey()
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		t.Fail()
	}
	defer session.Close()
	var tests = []struct {
		input    models.DeviceKey
		expected string
	}{
		{devicekey, ""},
	}
	for _, test := range tests {
		if output := CreateDeviceKey(session); output.Error() != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
			//t.Fail()
		}
	}
}
func TestSetKeyForDevice(t *testing.T) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		t.Fail()
	}
	defer session.Close()
	keydevice := models.DeviceKey{}
	keydevice.Device = "lamp"
	newKey, _ := GetValidKey(session)
	keydevice.Key = newKey.Key
	var tests = []struct {
		input    models.DeviceKey
		expected error
	}{
		{keydevice, nil},
	}
	for _, test := range tests {
		if output := AddKeyToDevice(keydevice, session); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
			//t.Fail()
		}
	}
}
func TestGenerateKey(t *testing.T) {
	key := GenerateKey()
	if len(key) != Words.LengthOfDeviceKey {
		t.Error("Test Failed: key is ", key)
	}
	fmt.Println(key)

}
func TestGetValidKey(t *testing.T) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		t.Fail()
	}
	defer session.Close()
	key, _ := GetValidKey(session)
	if len(key.Key) != Words.LengthOfDeviceKey {
		t.Error("Test Failed: key is ", key)
	}
	fmt.Println("key is : ", key)
}
func TestCheckKeyIsValid(t *testing.T) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		t.Fail()
	}
	defer session.Close()
	CheckKeyIsValid("", session)
	key, _ := GetValidKey(session)
	var tests = []struct {
		input    string
		expected bool
	}{
		{"alaki", false}, {key.Key, true},
	}
	for _, test := range tests {
		if output := CheckKeyIsValid(test.input, session); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
			//t.Fail()
		}
	}
}
