package models

import "gopkg.in/mgo.v2/bson"

type SettingsInDB struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
	Key  string        `json:"key" bson:"key"`
}
