package models

import "gopkg.in/mgo.v2/bson"

type DataBind struct {
	DataName     string    `json:"data_name" bson:"data_name"`
	CommandName  string    `json:"command_name" bson:"command_name"`
	ConditionSet Condition `json:"condition" bson:"condition"`
}

type Event struct {
	Id             bson.ObjectId `json:"id,string" bson:"_id"`
	EventName      string        `json:"event_name" bson:"event_name"`
	EventAddress   string        `json:"event_address" bson:"event_address"`
	EventType      EventType     `json:"event_type" bson:"event_type"`
	EventCondition Condition     `json:"event_condition" bson:"event_condition"`
	EventFunction  Command       `json:"event_function" bson:"event_function"`
}
type EventType int

const (
	MqttEvent EventType = 0
	SmsEvent  EventType = 1
)
