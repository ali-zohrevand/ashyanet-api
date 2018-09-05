package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Device struct {
	Id          string    `json:"id" bson:"_id"`
	Name        string    `json:"devicename" bson:"devicename" valid:"required~Device Name Could not be empty,runelength(1|30),blacklist~Bad Char"`
	Description string    `json:"description" bson:"description"`
	Type        string    `json:"type" bson:"type" valid:"required~Description Could not be empty,runelength(1|30),blacklist~Bad Char"`
	Key         string    `json:"key" bson:"key" valid:"required~Key Could not be empty,runelength(1|30),blacklist~Bad Char"`
	Owners      []string  `json:"owner" bson:"description"`
	Location    string    `json:"location" bson:"location" valid:"blacklist~Bad Char"`
	Publish     []string  `json:"publish" bson:"publish" valid:"runelength(1|30),blacklist~Bad Char"`
	Subscribe   []string  `json:"subscribe" bson:"subscribe" valid:"runelength(1|30),blacklist~Bad Char"`
	Pubsub      []string  `json:"pubsub" bson:"pubsub" valid:"runelength(1|30),blacklist~Bad Char"`
	Data        []Data    `json:"data" bson:"data" valid:"runelength(1|30),blacklist~Bad Char"`
	Command     []Command `json:"command" bson:"command" valid:"runelength(1|30),blacklist~Bad Char"`
}
type DeviceInDB struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"devicename" bson:"devicename"`
	Description string        `json:"description" bson:"description"`
	Type        string        `json:"type" bson:"type"`
	Key         string        `json:"key" bson:"key"`
	Owners      []UserInDB    `json:"owner" bson:"owner"`
	Location    string        `json:"location" bson:"location"`
	Publish     []string      `json:"publish" bson:"publish" valid:"runelength(1|30),blacklist~Bad Char"`
	Subscribe   []string      `json:"subscribe" bson:"subscribe" valid:"runelength(1|30),blacklist~Bad Char"`
	Pubsub      []string      `json:"pubsub" bson:"pubsub" valid:"runelength(1|30),blacklist~Bad Char"`
	Data        []Data        `json:"data" bson:"data" valid:"runelength(1|30),blacklist~Bad Char"`
	Command     []Command     `json:"command" bson:"command" valid:"runelength(1|30),blacklist~Bad Char"`
}

type Command struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Dsc   string `json:"dsc"`
	Topic string `json:"topic"`
}
type Data struct {
	Name      string `json:"name"`
	ValueType string `json:"value_type"`
	Dsc       string `json:"dsc"`
	Topic     string `json:"topic"`
}

/*
{
  "id": "",
  "devicename": "lamp",
  "description": "لامپ داخل اتاقل ",
  "type": "light",
  "key": "pFLsgUHtnFG65WhO2nDD",
  "owner": null,
  "location": "room",
  "publish": [
    "/Nf4KsSiAYJf2D7HReikI/sdsd/sdsad/545465/sdfsdf",
    "/Nf4KsSiAYJf2D7HReikI/sddfdfsd/sdsad",
    "/Nf4KsSiAYJf2D7HReikI/sdsd/sdsad/sdfsdf"
  ],
  "subscribe": [
    "/Nf4KsSiAYJf2D7HReikI/dfsdfsdsd/sdsad/545465/sdfsdf",
    "/Nf4KsSiAYJf2D7HReikI/sddfdfdfdfdf5456456465dfsd/sdsad",
    "/Nf4KsSiAYJf2D7HReikI/",
    "/Nf4KsSiAYJf2D7HReikI/"
  ],
  "pubsub": [
    "/Nf4KsSiAYJf2D7HReikI/d",
    "/Nf4KsSiAYJf2D7HReikI/sd656456465dfsd/sdsad"
  ],
  "data": [
    {
      "name": "Status",
      "value_type": "int",
      "dsc": "status",
      "topic": "/Nf4KsSiAYJf2D7HReikI/sds"
    }
  ],
  "command": [
    {
      "name": "On",
      "value": "on",
      "dsc": "Turn Light On",
      "topic": "/Nf4KsSiAYJf2D7HReikI/home/root"
    },
    {
      "name": "Of",
      "value": "off",
      "dsc": "Turn Light On",
      "topic": "/Nf4KsSiAYJf2D7HReikI/home/dffd/dfsdf"
    }
  ]
}
*/
