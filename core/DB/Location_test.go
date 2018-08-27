package DB

import (
	"fmt"
	"testing"
)

func TestGetLocationPath(t *testing.T) {
	session, errConnectDB := ConnectDB()
	if errConnectDB != nil {
		t.Fail()
	}
	defer session.Close()
	path, er := GetLocationPath("room", session)
	fmt.Println(path)
	fmt.Println(er)
}
