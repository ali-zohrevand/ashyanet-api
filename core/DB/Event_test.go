package DB

import (
	. "github.com/ali-zohrevand/ashyanet-api/models"
	"testing"
)

func TestEventCreate(t *testing.T) {
	var TrunOnCommand Command
	TrunOnCommand.Name = "On-" + GenerateRandomString(5)
	TrunOnCommand.Topic = "/home/-" + GenerateRandomString(8)
	TrunOnCommand.Dsc = "Turn Light On"
	TrunOnCommand.Value = "on"
	session, errConnectDB := ConnectDB()
	if errConnectDB != nil {
		t.Fail()
	}
	defer session.Close()
	var testIntCondtionLowerTHan Condition
	testIntCondtionLowerTHan.ConditionType = LowerThan
	testIntCondtionLowerTHan.GetAttr(5)
	var EventTest Event
	EventTest.EventType = MqttEvent
	EventTest.EventAddress = "/" + GenerateRandomString(8)
	EventTest.EventName = "test-" + GenerateRandomString(5)
	EventTest.EventCondition = testIntCondtionLowerTHan
	EventTest.EventFunction = TrunOnCommand
	er := EventCreate(EventTest, session)
	if er != nil {
		t.Fail()
		t.Error()
	}
}
