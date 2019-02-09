package DB

import (
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/settings/Words"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func LocationCreateWithUserName(Location models.Location, userName string, Session *mgo.Session) (err error) {
	err, exist := CheckLocationExist(userName+"-"+Location.Name, Session)
	if exist {
		return
	}
	user, errGetUser := UserGetByUsername(userName, Session)
	if errGetUser != nil {
		return errors.New(Words.UserNotExist)
	}
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	var LocationDB = models.LocationInDB{}
	LocationDB.Id = bson.NewObjectId()
	LocationDB.Description = Location.Description
	LocationDB.DisplayName = Location.DisplayName
	LocationDB.Name = userName + "-" + Location.Name
	LocationDB.Users = append(LocationDB.Users, userName)
	for i := 0; i < len(Location.Devices); i++ {
		_, deviceFound := DeviceGetByName(Location.Devices[i], sessionCopy)
		if deviceFound.Name == Location.Devices[i] {
			LocationDB.Devices = append(LocationDB.Devices, Location.Devices[i])
		} else {
			err = errors.New(Words.DeviceNotExist)
			return
		}
	}
	if Location.Parent != "" {
		_, exist := CheckLocationExist(Location.Parent, sessionCopy)
		if exist {
			LocationDB.Parent = Location.Parent

		} else {
			err = errors.New(Words.LocationNotFound)
			return
		}
	}
	err = sessionCopy.DB(Words.DBname).C(Words.LocationCollectionName).Insert(LocationDB)
	if err != nil {
		return
	}
	user.Locations = append(user.Locations, LocationDB.Name)
	err = sessionCopy.DB(Words.DBname).C(Words.UserCollectionName).UpdateId(user.Id, user)
	if err != nil {
		sessionCopy.DB(Words.DBname).C(Words.LocationCollectionName).RemoveId(LocationDB.Id)
	}
	return
}
func LocationCreate(Location models.Location, Session *mgo.Session) (err error) {
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
		_, deviceFound := DeviceGetByName(Location.Devices[i], sessionCopy)
		if deviceFound.Name == Location.Devices[i] {
			LocationDB.Devices = append(LocationDB.Devices, Location.Devices[i])
		} else {
			err = errors.New(Words.DeviceNotExist)
			return
		}
	}
	if Location.Parent != "" {
		_, exist := CheckLocationExist(Location.Parent, sessionCopy)
		if exist {
			LocationDB.Parent = Location.Parent

		} else {
			err = errors.New(Words.LocationNotFound)
			return
		}
	}
	err = sessionCopy.DB(Words.DBname).C(Words.LocationCollectionName).Insert(LocationDB)

	return
}
func CheckLocationExist(name string, Session *mgo.Session) (err error, Exist bool) {
	var Location models.LocationInDB
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err = sessionCopy.DB(Words.DBname).C(Words.LocationCollectionName).Find(bson.M{"locationname": name}).One(&Location)
	if Location.Name != "" && err == nil {
		err = errors.New(Words.LocationExist)
		Exist = true
		return
	}
	Exist = false
	return
}
func LocationGetByName(name string, Session *mgo.Session) (location models.LocationInDB, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err = sessionCopy.DB(Words.DBname).C(Words.LocationCollectionName).Find(bson.M{"locationname": name}).One(&location)
	return
}

func GetLocationPath(locationName string, Session *mgo.Session) (path string, err error) {
	name := locationName
	var ParentArray []string
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	for {
		Location, err := LocationGetByName(name, sessionCopy)
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
func LocationGetByUsername(userName string, Session *mgo.Session) (location []models.LocationInDB, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	_, errGeUser := UserGetByUsername(userName, sessionCopy)
	if errGeUser != nil {
		return nil, errGeUser
	}
	err = sessionCopy.DB(Words.DBname).C(Words.LocationCollectionName).Find(bson.M{"users": userName}).All(&location)
	return
}
func LocationGetById(id string, Session *mgo.Session) (location models.LocationInDB, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	idBson := bson.ObjectIdHex(id)
	err = sessionCopy.DB(Words.DBname).C(Words.LocationCollectionName).FindId(idBson).One(&location)
	return
}
func LocationDeleteByName(locationName string, Session *mgo.Session) (err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	locationObj, errGetTypes := LocationGetByName(locationName, sessionCopy)
	if errGetTypes != nil {
		return errGetTypes
	}
	err = sessionCopy.DB(Words.DBname).C(Words.LocationCollectionName).RemoveId(locationObj.Id)
	return
}
