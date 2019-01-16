package services

import (
	"github.com/ali-zohrevand/ashyanet-api/core/DB"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/services/log"
	"github.com/ali-zohrevand/ashyanet-api/settings/ConstKey"
	"gopkg.in/mgo.v2/bson"
)

func AddPermissionModelToDB() {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		panic(errConnectDB)

	}
	defer session.Close()

	perm := models.CasbinPermision{bson.NewObjectId(), ConstKey.PermissionModel, ConstKey.PermissionPolicy}
	PerMDB := DB.PermissionDataStore{}
	err := PerMDB.CreatePermissionInDB(perm, session)
	if err != nil {
		log.SystemErrorHappened(err)
		panic(err)
		return
	}

}
