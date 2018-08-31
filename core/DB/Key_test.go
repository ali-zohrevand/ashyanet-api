package DB

import (
	"fmt"
	"gitlab.com/hooshyar/ChiChiNi-API/Lab/mosquitto/03_crateTopicBaseLocation/DB"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/ConstKey"
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
	newKey := GetValidKey(session)
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
	if len(key) != ConstKey.LengthOfDeviceKey {
		t.Error("Test Failed: key is ", key)
	}

}
func TestGetValidKey(t *testing.T) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		t.Fail()
	}
	defer session.Close()
	key := GetValidKey(session)
	if len(key.Key) != ConstKey.LengthOfDeviceKey || key.Status != ConstKey.StatusValid {
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
	key := GetValidKey(session)
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