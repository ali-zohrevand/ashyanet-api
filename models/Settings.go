package models

import "gopkg.in/mgo.v2/bson"

type SettingsInDB struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Key      string        `json:"key" bson:"key"`
	Username string        `json:"username" bson:"username"`
	Password string        `json:"password" bson:"password"`
}
