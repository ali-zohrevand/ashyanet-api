package DB

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func CheckExist(fieldToSearch string, ToSearch string, Object interface{}, DbNameString string, ColloctionName string, errorString string, Session *mgo.Session) (err error, Exist bool) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err = sessionCopy.DB(DbNameString).C(ColloctionName).Find(bson.M{fieldToSearch: ToSearch}).One(&Object)
	if Object != nil && err == nil {
		err = errors.New(errorString)
		Exist = true
		return
	}
	Exist = false
	return
}
