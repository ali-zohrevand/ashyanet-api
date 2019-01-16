package models

import "gopkg.in/mgo.v2/bson"

type DeviceKey struct {
	Id     string `json:"id" bson:"_id"`
	Key    string `json:"key" bson:"key" valid:"required~Name Could not be empty,blacklist~Bad Char"`
	Device string `json:"device" bson:"device"`
	Status string `json:"status" bson:"status"`
}
type DeviceKeyInDB struct {
	Id     bson.ObjectId `json:"id,string" bson:"_id"`
	Key    string        `json:"key" bson:"key"`
	Device Device        `json:"device" bson:"device"`
	Status string        `json:"status" bson:"status"`
}

/*{
"Id": "",
"Key": "lamp",
"Device": "lamp of hall",
"Status": "light"
}*/
