package DB

import (
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func MqttAddMessage(message models.MqttMessage, Session *mgo.Session) (err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	message.Id = bson.NewObjectId()

	err = sessionCopy.DB(Words.DBname).C(Words.MqttMessageCollectionName).Insert(message)

	return
}
func MqttGetAllMessagesByTopic(topic string, Session *mgo.Session) (MessageList []models.MqttMessage, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err = sessionCopy.DB(Words.DBname).C(Words.MqttMessageCollectionName).Find(bson.M{"topic": topic}).All(&MessageList)
	return
}
