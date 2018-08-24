package models

import "gopkg.in/mgo.v2/bson"

type MqttUser struct {
	Id           string `json:"id" bson:"_id"`
	Username     string `json:"username" bson:"username"`
	Password     string `json:"password" bson:"password"`
	Is_superuser bool   `json:"is_superuser" bson:"is_superuser"`
	Created      string `json:"created" bson:"created"`
}
type MqttUserInDB struct {
	Id           bson.ObjectId `json:"id" bson:"_id"`
	Username     string        `json:"username" bson:"username"`
	Password     string        `json:"password" bson:"password"`
	Is_superuser bool          `json:"is_superuser" bson:"is_superuser"`
	Created      string        `json:"created" bson:"created"`
}

/*
{
    username: "user",
    password: "password hash",
    is_superuser: boolean (true, false),
    created: "datetime"
}
*/
