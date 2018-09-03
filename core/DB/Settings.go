package DB

import (
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func SaveSettings(settings models.SettingsInDB, Session *mgo.Session) (err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	settings.Id = bson.NewObjectId()
	err = sessionCopy.DB(Words.DBname).C(Words.SettingsCollectiName).Insert(settings)

	return
}
func LoadSettings(Session *mgo.Session) (settings models.SettingsInDB, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	sessionCopy.DB(Words.DBname).C(Words.SettingsCollectiName).Find(bson.M{}).One(&settings)
	return
}
