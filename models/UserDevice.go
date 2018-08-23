package models

type UserDevice struct {
	UserName   string `json:"username"  valid:"required~username Could not be empty,runelength(1|30),blacklist~Bad Char"`
	DeviceName string `json:"devicename"  valid:"required~v Could not be empty,runelength(1|30),blacklist~Bad Char"`
}

/*
{
  "UserName": "ali",
  "DeviceName": "lamp"

}
*/
