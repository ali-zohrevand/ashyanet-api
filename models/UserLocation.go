package models

type UserLocation struct {
	UserName     string `json:"username"  valid:"required~username Could not be empty,runelength(1|30),blacklist~Bad Char"`
	LocationName string `json:"locationname"  valid:"required~locationname Could not be empty,runelength(1|30),blacklist~Bad Char"`
}

/*
{
  "UserName": "ali",
  "LocationName": "room"

}
*/
