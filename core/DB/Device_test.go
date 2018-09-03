package DB

import (
	"gitlab.com/hooshyar/ChiChiNi-API/Lab/mosquitto/03_crateTopicBaseLocation/DB"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"testing"
)

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
	if output := CreateDevice(test.input,session); output != test.expected {
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
	errExisted:=errors.New(ConstKey.DeviceExist)
	var tests = []struct {
		input    models.Device
		expected error
	}{
	{ValidUser,errExisted},

	}
	for _, test := range tests {
		if output := CreateDevice(test.input,session); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
			//t.Fail()
		}
	}
}*/

func TestAddUserToDevice(t *testing.T) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		t.Fail()
	}
	defer session.Close()
	ValidaUserTodevice := models.UserDevice{"ali", "lamp"}
	var tests = []struct {
		input    models.UserDevice
		expected error
	}{
		{ValidaUserTodevice, nil},
	}
	for _, test := range tests {
		output := AddUserToDevice(test.input, session)
		if output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
			//t.Fail()
		}
	}
}
