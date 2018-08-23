package models

type Location struct {
	Id     string `json:"id" bson:"_id"`
	Name   string
	Parent Location
	DeviceListId
}
