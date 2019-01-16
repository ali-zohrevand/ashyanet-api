package services

import (
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/settings/Words"
	"math/rand"
	"strings"
	"time"
)

func CheckMqttTopic(device *models.Device, user models.UserInDB) (OutputDevice *models.Device, err error) {
	root := ""
	settings, err := GetServerSettings()
	if err != nil {
		return
	}

	/*if settings.Type == "server" {
		root = user.UserName
	} else if settings.Type == "gateway" {
		root = settings.Identifier

	} else {
		root = settings.Identifier
	}*/
	root = settings.Identifier + "/" + user.UserName
	for i, topicPath := range device.Publish {
		topicPath = AddRootTopic(root, topicPath)
		device.Publish[i] = topicPath
	}
	for i, topicPath := range device.Subscribe {
		topicPath = AddRootTopic(root, topicPath)
		device.Subscribe[i] = topicPath
	}
	for i, topicPath := range device.Pubsub {
		topicPath = AddRootTopic(root, topicPath)
		device.Pubsub[i] = topicPath
	}
	for i, command := range device.MqttCommand {
		command.Topic = AddRootTopic(root, command.Topic)
		device.MqttCommand[i].Topic = command.Topic
	}
	for i, data := range device.MqttData {
		data.Topic = AddRootTopic(root, data.Topic)
		device.MqttData[i].Topic = data.Topic
	}
	return device, nil
}
func GenerateRandomString(length int) string {
	var charset string
	charset = Words.RuneCharInKey
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func AddRootTopic(rootPath string, TopicPath string) (UpdatedPath string) {
	PathArray := strings.Split(TopicPath, "/")
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
	MqttHttpCommand     []MqttHttpCommand `json:"command" bson:"command" valid:"runelength(1|30),blacklist~Bad Char"`
}

*/
