package DB

import (
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/ConstKey"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func SaveSettings(settings models.SettingsInDB, Session *mgo.Session) (err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	settings.Id = bson.NewObjectId()
	err = sessionCopy.DB(ConstKey.DBname).C(ConstKey.SettingsCollectiName).Insert(settings)

	return
}
