package services

import (
	"fmt"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"strings"
)

func CheckMqttTopic(device *models.Device, user models.UserInDB) (err error) {
	settings, err := GetServerSettings()
	if err != nil {
		return
	}
	switch settings.Type {
	case "gateway":

	default:

	}
	return
}
func AddRootTopic(rootPath string, TopicPath string) (UpdatedPath string) {
	PathArray := strings.Split(TopicPath, "/")
	fmt.Println(len(PathArray))
	if len(PathArray) == 0 {
		return
	}
	if PathArray[0] == rootPath {
		return TopicPath
	} else if len(PathArray) >= 2 && PathArray[1] == rootPath {
		return TopicPath
	} else {
		var NewArrPath []string
		NewArrPath = append(NewArrPath, rootPath)
		for _, v := range PathArray {
			NewArrPath = append(NewArrPath, v)
		}
		for _, v := range NewArrPath {
			UpdatedPath = UpdatedPath + "/" + v
		}
	}
	fmt.Println(UpdatedPath)
	UpdatedPath = strings.Replace(UpdatedPath, "//", "/", -1)
	return
}

/*
type Device struct {
	Id          string   `json:"id" bson:"_id"`
	Name        string   `json:"devicename" bson:"devicename" valid:"required~Device Name Could not be empty,runelength(1|30),blacklist~Bad Char"`
	Description string   `json:"description" bson:"description"`
	Type        string   `json:"type" bson:"type" valid:"required~Description Could not be empty,runelength(1|30),blacklist~Bad Char"`
	Key         string   `json:"key" bson:"key" valid:"required~Key Could not be empty,runelength(1|30),blacklist~Bad Char"`
	Owners      []string `json:"owner" bson:"description"`
	Location    string   `json:"location" bson:"location" valid:"blacklist~Bad Char"`
	Publish     []string `json:"publish" bson:"publish" valid:"runelength(1|30),blacklist~Bad Char"`
	Subscribe   []string `json:"subscribe" bson:"subscribe" valid:"runelength(1|30),blacklist~Bad Char"`
	Pubsub      []string `json:"pubsub" bson:"pubsub" valid:"runelength(1|30),blacklist~Bad Char"`
	Data        []Data	 `json:"data" bson:"data" valid:"runelength(1|30),blacklist~Bad Char"`
	Command     []Command `json:"command" bson:"command" valid:"runelength(1|30),blacklist~Bad Char"`
}

*/
