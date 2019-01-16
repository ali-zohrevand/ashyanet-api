package services

import (
	"fmt"
	"github.com/ali-zohrevand/ashyanet-api/core/DB"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/services/log"
	"github.com/ali-zohrevand/ashyanet-api/settings/Words"
	"github.com/casbin/casbin"
	"github.com/casbin/mongodb-adapter"
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
