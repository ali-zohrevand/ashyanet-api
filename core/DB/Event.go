package DB

import (
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
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
