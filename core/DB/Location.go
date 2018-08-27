package DB

import (
	"github.com/pkg/errors"
	. "gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/ConstKey"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func CreateLocation(Location Location, Session *mgo.Session) (err error) {
	err, exist := CheckLocationExist(Location.Name, Session)
	if exist {
		return
	}
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	var LocationDB = LocationInDB{}
	LocationDB.Id = bson.NewObjectId()
	LocationDB.Description = Location.Description
	LocationDB.Name = Location.Name
	for i := 0; i < len(Location.Devices); i++ {
		_, deviceFound := FindDeviceByName(Location.Devices[i], sessionCopy)
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
	var Location LocationInDB
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
func GetLocationByName(name string, Session *mgo.Session) (location LocationInDB, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err = sessionCopy.DB(ConstKey.DBname).C(ConstKey.LocationCollectionName).Find(bson.M{"locationname": name}).One(&location)
	return
}
func AddDeviceToLocation() {

}
func GetLocationPath(locationName string, Session *mgo.Session) (path string, err error) {
	name := locationName
	var ParentArray []string
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	for {
		Location, err := GetLocationByName(name, sessionCopy)
		if err != nil {
			break
		}

		if Location.Parent == "" {
			break
		}
		ParentArray = append(ParentArray, Location.Parent)

		name = Location.Parent
	}
	if len(ParentArray) > 0 {
		for i := len(ParentArray) - 1; i >= 0; i-- {
			path = addPath(path, ParentArray[i])
		}
		path = addPath(path, locationName)
	} else {
		path = addPath(path, locationName)

	}

	return
}
