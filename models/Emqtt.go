package models

import "gopkg.in/mgo.v2/bson"

type MqttUser struct {
	Id           string `json:"id" bson:"_id" `
	Username     string `json:"username" bson:"username" valid:"required~Username Could not be empty,runelength(1|30),blacklist~Bad Char"`
	Password     string `json:"password" bson:"password" valid:"required~password Could not be empty,runelength(6|30)~Password must be between 6 and 30 char,blacklist~Bad Char"`
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
    "username" :"root",
    "password" : "123456",
    "is_superuser" : true,
    "created" : ""
}
*/

type MqttAcl struct {
	Id        string   `json:"id" bson:"_id"`
	Username  string   `json:"username" bson:"username"`
	Clientid  string   `json:"clientid" bson:"clientid"`
	Publish   []string `json:"publish" bson:"publish"`
	Subscribe []string `json:"subscribe" bson:"subscribe"`
	Pubsub    []string `json:"pubsub" bson:"pubsub" `
}
type MqttAclIndb struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Username  string        `json:"username" bson:"username"`
	Clientid  string        `json:"clientid" bson:"clientid"`
	Publish   []string      `json:"publish" bson:"publish"`
	Subscribe []string      `json:"subscribe" bson:"subscribe"`
	Pubsub    []string      `json:"pubsub" bson:"pubsub" `
}

/*
{
    username: "username",
    clientid: "clientid",
    publish: ["topic1", "topic2", ...],
    subscribe: ["subtop1", "subtop2", ...],
    pubsub: ["topic/#", "topic1", ...]
}
*/
