package DB

import (
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/settings/Words"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func EventCreate(Event models.Event, Session *mgo.Session) (err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	Event.Id = bson.NewObjectId()
	err = sessionCopy.DB(Words.DBname).C(Words.EventCollectionName).Insert(Event)
	return
}
func EventGetAddress(EvenaAddress string, Session *mgo.Session) (event models.Event, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err = sessionCopy.DB(Words.DBname).C(Words.EventCollectionName).Find(bson.M{"event_address": EvenaAddress}).One(&event)
	return
}
