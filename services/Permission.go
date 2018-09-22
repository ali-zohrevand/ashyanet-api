package services

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/mongodb-adapter"
	"gitlab.com/hooshyar/ChiChiNi-API/core/DB"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/services/log"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
	"gopkg.in/mgo.v2/bson"
)

func AddPermissionModelToDB() {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		panic(errConnectDB)

	}
	defer session.Close()

	perm := models.CasbinPermision{bson.NewObjectId(), Words.PermissionModel, Words.PermissionPolicy}
	PerMDB := DB.PermissionDataStore{}
	err := PerMDB.CreatePermissionInDB(perm, session)
	if err != nil {
		log.SystemErrorHappened(err)
		panic(err)
		return
	}

}
func AddDefaultModelConfToDB() {
	a := mongodbadapter.NewAdapter("127.0.0.1:27017/" + Words.DBname) // Your MongoDB URL.
	e := casbin.NewEnforcer("auth_model.conf", a)
	b := e.IsFiltered()
	fmt.Println(b)

}
