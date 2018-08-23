package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Device struct {
	Id          string   `json:"id" bson:"_id"`
	Name        string   `json:"devicename" bson:"devicename" valid:"required~Device Name Could not be empty,runelength(1|30),blacklist~Bad Char"`
	Description string   `json:"description" bson:"description"`
	Type        string   `json:"type" bson:"type" valid:"required~Description Could not be empty,runelength(1|30),blacklist~Bad Char"`
	Key         string   `json:"key" bson:"key" valid:"required~Key Could not be empty,runelength(1|30),blacklist~Bad Char"`
	Owners      []string `json:"owner" bson:"description"`
	Location    string   `json:"location" bson:"location" valid:"blacklist~Bad Char" '`
	Topics      []string `json:"topics" bson:"topics" valid:"blacklist~Bad Char"`
	MqttPass    string   `json:"mqtt_pass" bson:"mqtt_pass"`
}
type DeviceInDB struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"devicename" bson:"devicename"`
	Description string        `json:"description" bson:"description"`
	Type        string        `json:"type" bson:"type"`
	Key         string        `json:"key" bson:"key"`
	Owners      []UserInDB    `json:"owner" bson:"owner"`
	Location    string        `json:"location" bson:"location"`
	Topics      []string      `json:"topics" bson:"topics"`
	MqttPass    string        `json:"mqtt_pass" bson:"mqtt_pass"`
}

/*
{
  "Id": "",
  "DeviceName": "lamp",
  "Description": "lamp of hall",
  "Type": "light",
  "Key": "afhdkjfhsdkjfhksdjfhk",
  "Owners": []
}
*/
