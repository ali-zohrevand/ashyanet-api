package models

import "gopkg.in/mgo.v2/bson"

type SettingsInDB struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Identifier string        `json:"identifier" bson:"identifier"`
	Key        string        `json:"key" bson:"key"`
	Type       string        `json:"type" bson:"type"`
	Username   string        `json:"username" bson:"username"`
	Password   string        `json:"password" bson:"password"`
}
