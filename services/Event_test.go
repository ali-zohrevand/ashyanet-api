package services

import (
	"encoding/json"
	"fmt"
	"github.com/ali-zohrevand/ashyanet-api/core/DB"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"testing"
)

func TestEventCreate(t *testing.T) {
	session, err := DB.ConnectDB()
	if err != nil {
		t.Fail()
		t.Error(err)
	}
	user, err := DB.UserGetByUsername("user6", session)
	if err != nil {
		t.Fail()
		t.Error(err)
	}
	condition := models.Condition{}
	condition.Attr = append(condition.Attr, 5)
	condition.ConditionType = models.GraterThan

	var dataBinde models.DataBindCommand
	dataBinde.ComandType = models.MqttEvent
	dataBinde.CommandName = "On"
	dataBinde.DataName = "Status"
	dataBinde.ConditionSet = condition
	//...............................................
	dataBindINJson, err := json.Marshal(dataBinde)
	if err != nil {
		t.Fail()
		t.Error(err)
	}
	fmt.Println(string(dataBindINJson), user)
}
