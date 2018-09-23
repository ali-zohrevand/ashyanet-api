package DB

import (
	"gitlab.com/hooshyar/ChiChiNi-API/models"
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
