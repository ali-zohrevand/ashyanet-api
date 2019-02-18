package models

import "gopkg.in/mgo.v2/bson"

type MqttMessage struct {
	Id        bson.ObjectId `json:"id,string" bson:"_id"`
	Topic     string        `json:"topic" bson:"topic"`
	Message   string        `json:"message" bson:"message"`
	MessageId string        `json:"message_id" bson:"message_id"`
	Qos       byte          `json:"qos,string" bson:"qos"`
	Retained  bool          `json:"retained" bson:"retained"`
	Time      int64         `json:"time" bson:"time"`
}

type MqttDataRequest struct {
	Topic string `json:"topic"`
}
type MqttInfo struct {
	NumberOfTopics  string `json:"number_of_topic"`
	NumberOfMessage string `json:"number_of_message"`
}
