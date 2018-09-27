package models

import "gopkg.in/mgo.v2/bson"

type MqttMessage struct {
	Id        bson.ObjectId `json:"id,string" bson:"_id"`
	Topic     string        `json:"topic" bson:"topic"`
	Message   string        `json:"message" bson:"message"`
	MessageId string        `json:"message_id" bson:"message_id"`
	Qos       string        `json:"qos,string" bson:"qos"`
	Retained  bool          `json:"retained,string" bson:"retained"`
	Time      string        `json:"time" bson:"time"`
}
type MqttCommand struct {
	Name  string `json:"name" valid:"runelength(1|60),blacklist~Bad Char"`
	Value string `json:"value" valid:"runelength(1|60),blacklist~Bad Char"`
	Dsc   string `json:"dsc"`
	Topic string `json:"topic" valid:"runelength(1|60),blacklist~Bad Char"`
}
type MqttData struct {
	Name      string `json:"name" valid:"runelength(1|30),blacklist~Bad Char"`
	ValueType string `json:"value_type" valid:"runelength(1|60),blacklist~Bad Char"`
	Dsc       string `json:"dsc"`
	Topic     string `json:"topic" valid:"runelength(1|30),blacklist~Bad Char"`
}
