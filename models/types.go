package models

import "gopkg.in/mgo.v2/bson"

type Types struct {
	Id       bson.ObjectId `json:"id,string" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Dsc      string        `json:"dsc" bson:"dsc"`
	IconName string        `json:"-"  bson:"icon"`
	Owner    string        `json:"owner"`
}
