package models

import "gopkg.in/mgo.v2/bson"

type Location struct {
	Id          string   `json:"id" bson:"_id"`
	Name        string   `json:"locationname" bson:"locationname" valid:"required~Location Could not be empty,runelength(1|30),blacklist~Bad Char"`
	Parent      string   `json:"parent" bson:"parent" valid:"runelength(1|30),blacklist~Bad Char"`
	Devices     []string `json:"devices" bson:"devices"`
	Description string   `json:"dsc" bson:"dsc"`
	Latitude    string   `json:"latitude" bson:"lat" valid:"IsLatitude"`
	Longitude   string   `json:"longitude" bson:"long" valid:"IsLongitude"`
	Users       []string `json:"users" bson:"users"`
}
type LocationInDB struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"locationname" bson:"locationname"`
	Parent      string        `json:"parent" bson:"parent"`
	Devices     []string      `json:"devices" bson:"devices"`
	Description string        `json:"dsc" bson:"dsc"`
	Latitude    string        `json:"latitude" bson:"lat" valid:"IsLatitude"`
	Longitude   string        `json:"longitude" bson:"long" valid:"IsLongitude"`
	Users       []string      `json:"users" bson:"users"`
}

/*
{
  "Id": "",
  "locationname": "kisdthddffcen",
  "parent": "home",
  "devices": [],
  "dsc": "آشپزخانه",
  "latitude": "",
  "longitude": ""
}
*/
