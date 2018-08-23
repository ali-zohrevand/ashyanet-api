package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type JwtSession struct {
	Id            bson.ObjectId `json:"id" bson:"_id"`
	JwtToken      string        `json:"token" bson:"token"`
	OwnerUsername string        `json:"username" bson:"username"`
	TimeCreated   time.Time     `json:"time_created"bson:"time"`
	Ip            string        `json:"ip" bson:"ip"`
}
