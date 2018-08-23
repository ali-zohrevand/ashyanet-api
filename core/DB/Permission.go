package DB

import (
	"fmt"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/services/log"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/ConstKey"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PermissionDataStore struct {
}

func (per *PermissionDataStore) CreatePermissionInDB(casbin models.CasbinPermision, dbSession *mgo.Session) (err error) {
	sessionCopy := dbSession.Copy()
	defer sessionCopy.Close()
	if !per.IsThereAnyModelSet(dbSession) {
		err = sessionCopy.DB(ConstKey.DBname).C(ConstKey.PermissionCollectionName).Insert(casbin)
		return err
	} else {
		return
	}

}
func (per *PermissionDataStore) IsThereAnyModelSet(Session *mgo.Session) (Is bool) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	var CasbinModel []models.CasbinPermision
	err := sessionCopy.DB(ConstKey.DBname).C(ConstKey.PermissionCollectionName).Find(bson.M{}).All(&CasbinModel)

	if len(CasbinModel) == 0 {
		return false
	}
	log.SystemErrorHappened(err)
	fmt.Println(err)
	return true

}
func GetPermision(Session *mgo.Session) (CasbinModel models.CasbinPermision) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()

	sessionCopy.DB(ConstKey.DBname).C(ConstKey.PermissionCollectionName).Find(bson.M{}).One(&CasbinModel)
	return

}
