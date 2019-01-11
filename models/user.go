package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id        string   `json:"id" bson:"_id"`
	UserName  string   `json:"username" bson:"username" valid:"required~Username Could not be empty,runelength(1|30),blacklist~Bad Char"`
	FirstName string   `json:"firstname"  bson:"firstname" valid:"required~First name Could not be empty,runelength(1|30),blacklist~Bad Char"`
	LastName  string   `json:"lastname" bson:"lastname" valid:"required~Last name Could not be empty,runelength(1|30),blacklist~Bad Char"`
	Email     string   `json:"email" bson:"email" valid:"required~Email Could not be empty,runelength(1|30),email~Email is not valid,blacklist~Bad Char"`
	Password  string   `json:"password" bson:"password" valid:"required~Password Could not be empty,runelength(5|30)~Password length Must be between 5 and 30 "`
	Role      string   `json:"role" bson:"role"`
	Locations []string `json:"locations" bson:"locations"`
	Devices   []string `json:"devices" bson:"devices"`
	Active 	  bool      `json:"active" bson:"active"`
	TempKeyGenreated string `json:"-" bson:"temp_key_genreated"`
}
type UserInDB struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	UserName  string        `json:"username" bson:"username" `
	FirstName string        `json:"firstname" bson:"firstname"`
	LastName  string        `json:"lastname" bson:"lastname"`
	Email     string        `json:"email" bson:"email"`
	Password  []byte        `json:"password" bson:"password"`
	Role      string        `json:"role" bson:"role"`
	Locations []string      `json:"locations" bson:"locations"`
	Devices   []string      `json:"devices" bson:"devices"`
	Active 	  bool     	    `json:"active" bson:"active"`
	TempKeyGenreated string `json:"-" bson:"temp_key_genreated"`

}

/*
{
  "Id": "",
  "UserName": "ali",
  "FirstName": "adfdfli",
  "LastName": "zohrevand",
  "Email": "ali@ali.ir",
  "Password": "mahdi"
}
*/
