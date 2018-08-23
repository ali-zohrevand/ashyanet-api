package DB

import (
	"github.com/pkg/errors"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/ConstKey"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func CreateLocation(Location models.Location, Session *mgo.Session) (err error) {
	err, exist := CheckLocationExist(Location.Name, Session)
	if exist {
		return
	}
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	var LocationDB = models.LocationInDB{}
	LocationDB.Id = bson.NewObjectId()
	LocationDB.Description = Location.Description
	LocationDB.Name = Location.Name
	for i := 0; i < len(Location.Devices); i++ {
		_, deviceFound := FindeDeviceByName(Location.Devices[i], sessionCopy)
		if deviceFound.Name == Location.Devices[i] {
			LocationDB.Devices = append(LocationDB.Devices, Location.Devices[i])
		} else {
			err = errors.New(ConstKey.DeviceNotExist)
			return
		}
	}
	if Location.Parent != "" {
		_, exist := CheckLocationExist(Location.Parent, sessionCopy)
		if exist {
			LocationDB.Parent = Location.Parent

		} else {
			err = errors.New(ConstKey.LocationNotFound)
			return
		}
	}
	err = sessionCopy.DB(ConstKey.DBname).C(ConstKey.LocationCollectionName).Insert(LocationDB)
	return
}
func CheckLocationExist(name string, Session *mgo.Session) (err error, Exist bool) {
	var Location models.LocationInDB
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err = sessionCopy.DB(ConstKey.DBname).C(ConstKey.LocationCollectionName).Find(bson.M{"locationname": name}).One(&Location)
	if Location.Name != "" && err == nil {
		err = errors.New(ConstKey.LocationExist)
		Exist = true
		return
	}
	Exist = false
	return
}
func AddDeviceToLocation() {

}
