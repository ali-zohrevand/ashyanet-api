package services

import (
	"fmt"
	"gitlab.com/hooshyar/ChiChiNi-API/core/DB"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"testing"
)

func TestAddRootTopic(t *testing.T) {
	p := AddRootTopic("sadasdasd", "/")
	fmt.Println(p)
}
func TestCheckMqttTopic(t *testing.T) {
	device := models.Device{}
	s, _ := DB.ConnectDB()
	user, _ := DB.FindUserByUsername("admin", s)
	fmt.Println(user)
	CheckMqttTopic(&device, user)
}
