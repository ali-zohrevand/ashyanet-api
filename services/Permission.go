package services

import (
	"gitlab.com/hooshyar/ChiChiNi-API/core/DB"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/services/log"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/ConstKey"
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
