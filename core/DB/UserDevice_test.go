package DB

import (
	"github.com/ali-zohrevand/ashyanet-api/models"
	"testing"
)

func TestAddUserToDevice(t *testing.T) {
	session, errConnectDB := ConnectDB()
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
		output := AddUserDevice(test.input, session)
		if output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
			//t.Fail()
		}
	}
}
