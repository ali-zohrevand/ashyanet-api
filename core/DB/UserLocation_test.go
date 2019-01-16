package DB

import (
	"github.com/ali-zohrevand/ashyanet-api/models"
	"testing"
)

func TestAddUserLocation(t *testing.T) {
	session, errConnectDB := ConnectDB()
	if errConnectDB != nil {
		t.Fail()
	}
	defer session.Close()
	userLocation := models.UserLocation{}
	userLocation.UserName = "admin"
	userLocation.LocationName = "room"
	err := AddUserLocation(userLocation, session)
	if err != nil {
		t.Fail()
		t.Error(err)

	}
}
