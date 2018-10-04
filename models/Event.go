package models

import "gopkg.in/mgo.v2/bson"

type Event struct {
	Id             bson.ObjectId `json:"id,string" bson:"_id"`
	EventName      string        `json:"event_name" bson:"event_name"`
	EventAddress   string        `json:"event_address" bson:"event_address"`
	EventType      EventType     `json:"event_type" bson:"event_type"`
	EventCondition Condition     `json:"event_condition" bson:"event_condition"`
}
type EventType int

const (
	MqttEvent EventType = 0
	SmsEvent  EventType = 1
)
