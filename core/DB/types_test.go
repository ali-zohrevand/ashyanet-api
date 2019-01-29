package DB

import (
	"fmt"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"testing"
)

func TestTypesCreate(t *testing.T) {
	var typeObj models.Types
	typeObj.Name = "newTypeTest"
	typeObj.Dsc = "justForTest"
	typeObj.Owner = "No One Yet!"
	typeObj.IconName = "nothing yet"

	session, errConnectDB := ConnectDB()
	if errConnectDB != nil {
		t.Fail()
	}
	errCreate := TypesCreate(typeObj, session)
	if errCreate != nil {
		t.Fail()
		t.Error(errCreate)
	}
	types, errGetAll := TypeGetAll(session)
	if errGetAll != nil {
		t.Fail()
		t.Error(errGetAll)
	}
	_, errOne := TypeGetTypeByName(typeObj.Name, session)
	if errOne != nil {

		t.Fail()
		t.Error(errOne)
	}
	is := TypeIsTypeExist(typeObj.Name, session)
	if !is {

		t.Fail()
		t.Error("not exist")
	}
	fmt.Println("all types: ", types)
	errDelete := TypeDeleteByName(typeObj.Name, session)
	if errDelete != nil {
		t.Fail()
		t.Error(errDelete)
	}
}
