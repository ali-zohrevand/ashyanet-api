package models

import "gopkg.in/mgo.v2/bson"

type SettingsInDB struct {
	Id                 bson.ObjectId `json:"id" bson:"_id"`
	Url                string        `json:"url" bson:"url"`
	Identifier         string        `json:"identifier" bson:"identifier"`
	Key                string        `json:"key" bson:"key"`
	Type               string        `json:"type" bson:"type"`
	Username           string        `json:"username" bson:"username"`
	Password           string        `json:"password" bson:"password"`
	MailHost           string        `json:"mail_host" bson:"mail_host"`
	MailPort           string        `json:"mail_port" bson:"mail_port"`
	MailVerifyUsername string        `json:"mail_verify_username" bson:"mail_verify_username"`
	MailVerifyPassword string        `json:"mail_verify_password" bson:"mail_verify_password"`
}
