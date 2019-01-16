package DB

import (
	"SimpleAPIBasePlatform/SimpleAPI/core/DB"
	"fmt"
	"testing"
)

func TestUserGetAllTopic(t *testing.T) {

	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		t.Fail()
	}
	defer session.Close()
	Topic, _ := UserGetAllTopic("ali", "pub", session)
	fmt.Println(Topic, "   ", len(Topic))
	Topic, _ = UserGetAllTopic("ali", "sub", session)

	fmt.Println(Topic, "   ", len(Topic))
	Topic, _ = UserGetAllTopic("ali", "pubsub", session)

	fmt.Println(Topic, "   ", len(Topic))

}
func TestUserGetAllCommand(t *testing.T) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		t.Fail()
	}
	defer session.Close()
	command, _ := UserGetAllMqttCommand("ali", session)
	fmt.Println(command)
}
func TestUserGetAllDevice(t *testing.T) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		t.Fail()
	}
	defer session.Close()
	a, _ := UserGetAllDevice("ali", session)
	fmt.Println(a)
}
func TestUserGetAllCommandDat(t *testing.T) {
	session, errConnectDB := ConnectDB()
	if errConnectDB != nil {
		t.Fail()
	}
	commands, data, err := UserGetAllCommandData("user6", session)
	if err!=nil{
		t.Fail()

	}
	fmt.Println(commands)
	fmt.Println(data)
}