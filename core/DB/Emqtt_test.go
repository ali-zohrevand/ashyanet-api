package DB

import (
	"fmt"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"testing"
	"time"
)

func TestGetEmqttUserByUserName(t *testing.T) {
	session, errConnectDB := ConnectDB()
	if errConnectDB != nil {
		t.Fail()
	}
	user := models.MqttUser{}
	user.Username = "admin"
	user.Password = "123456789"
	user.Is_superuser = true
	user.Created = time.Now().String()
	err := EmqttCreateUser(user, session)
	if err != nil {
		t.Fail()
		t.Error(err)
	}
	fmt.Println("Create EmqttUser Checked!", user)
	u, errGetUser := EmqttGetUserByUserName(user.Username, session)
	if errGetUser != nil {
		t.Fail()
		t.Error(errGetUser)
	}
	fmt.Println("Get EmqttUser Checked! ", u)

	/*	errDele := EmqttDeleteUser(user.Username, session)
		if errDele != nil {
			t.Fail()
			t.Error(errDele)
		}
		fmt.Println("Delete EmqttUser Checked!")
	*/
}
