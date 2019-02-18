package DB

import (
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/settings/Words"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func TypesCreate(typeObj models.Types, Session *mgo.Session) (err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err, exist := CheckExist("name", typeObj.Name, models.Types{}, Words.DBname, Words.TypesCollectionName, Words.TypeExits, sessionCopy)
	if exist {
		return
	}
	typeObj.Id = bson.NewObjectId()
	err = sessionCopy.DB(Words.DBname).C(Words.TypesCollectionName).Insert(typeObj)
	return
}
func TypeGetAll(Session *mgo.Session) (types []models.Types, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err = sessionCopy.DB(Words.DBname).C(Words.TypesCollectionName).Find(bson.M{}).All(&types)
	return
}
func TypesGetAllTypesOfUser(userName string, Session *mgo.Session) (types []models.Types, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	_, errGeUser := UserGetByUsername(userName, sessionCopy)
	if errGeUser != nil {
		return nil, errGeUser
	}
	err = sessionCopy.DB(Words.DBname).C(Words.TypesCollectionName).Find(bson.M{"owner": userName}).All(&types)
	return
}
func TypeIsTypeExist(typeName string, Session *mgo.Session) (is bool) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err, exist := CheckExist("name", typeName, models.Types{}, Words.DBname, Words.TypesCollectionName, Words.TypeExits, sessionCopy)
	if exist && err.Error() == Words.TypeExits {
		return true
	}
	return false

}
func TypeGetTypeByName(typeName string, Session *mgo.Session) (typeObj models.Types, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err = sessionCopy.DB(Words.DBname).C(Words.TypesCollectionName).Find(bson.M{"name": typeName}).One(&typeObj)
	return
}
func TypeGetTypeByID(id string, Session *mgo.Session) (typeObj models.Types, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	idBson := bson.ObjectIdHex(id)
	err = sessionCopy.DB(Words.DBname).C(Words.TypesCollectionName).FindId(idBson).One(&typeObj)
	return
}
func TypeDeleteByName(typeName string, Session *mgo.Session) (err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	typeObj, errGetTypes := TypeGetTypeByName(typeName, sessionCopy)
	if errGetTypes != nil {
		return errGetTypes
	}
	err = sessionCopy.DB(Words.DBname).C(Words.TypesCollectionName).RemoveId(typeObj.Id)
	return
}
