package DB

import (
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/settings/ConstKey"
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
func LoadSettings(Session *mgo.Session) (settings models.SettingsInDB, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	sessionCopy.DB(ConstKey.DBname).C(ConstKey.SettingsCollectiName).Find(bson.M{}).One(&settings)
	return
}
