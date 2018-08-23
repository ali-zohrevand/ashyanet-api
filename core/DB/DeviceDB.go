package DB

import (
	"github.com/pkg/errors"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/ConstKey"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func CreateDevice(device models.Device, Session *mgo.Session) (err error) {
	//err,exist:=FindeDeviceByName(device.Name,Session)
	err, exist := CheckExist("devicename", device.Name, models.DeviceInDB{}, ConstKey.DBname, ConstKey.DeviceCollectionName, ConstKey.DeviceExist, Session)
	if exist {
		return
	}
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	var DeviceDB = models.DeviceInDB{}
	DeviceDB.Id = bson.NewObjectId()
	DeviceDB.Name = device.Name
	DeviceDB.Description = device.Description
	if device.Key != "" {
		IsValid := CheckKeyIsValid(device.Key, sessionCopy)
		if !IsValid {
			errKeyIsnotValid := errors.New(ConstKey.KeyIsNotValid)
			return errKeyIsnotValid
		}
	}
	DeviceDB.Key = device.Key
	if len(device.Owners) > 0 {
		for _, user := range device.Owners {
			userFetchedFromDB, err := FindUserByUsername(user, sessionCopy)
			if err != nil {
				err = errors.New(user + ": " + ConstKey.UserNotExist)
				return err
			}
			DeviceDB.Owners = append(DeviceDB.Owners, userFetchedFromDB)
		}
	}
	userFetchedFromDB, err := FindUserByUsername(ConstKey.DeafualtAdmminUserName, sessionCopy)
	DeviceDB.Owners = append(DeviceDB.Owners, userFetchedFromDB)
	DeviceDB.Type = device.Type
	err = sessionCopy.DB(ConstKey.DBname).C(ConstKey.DeviceCollectionName).Insert(DeviceDB)
	return
}
func AddUserToDevice(userToDevice models.UserDevice, Session *mgo.Session) (err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	var user models.UserInDB
	err = sessionCopy.DB(ConstKey.DBname).C(ConstKey.UserCollectionName).Find(bson.M{"username": userToDevice.UserName}).One(&user)
	if err != nil {
		err = errors.New(ConstKey.UserNotExist)
		return
	}
	var deviceToAdd models.DeviceInDB
	err = sessionCopy.DB(ConstKey.DBname).C(ConstKey.DeviceCollectionName).Find(bson.M{"devicename": userToDevice.DeviceName}).One(&deviceToAdd)
	if err != nil {
		err = errors.New(ConstKey.DeviceNotExist)

		return
	}
	for _, u := range deviceToAdd.Owners {
		if u.UserName == userToDevice.UserName {
			return
		}
	}
	deviceToAdd.Owners = append(deviceToAdd.Owners, user)
	err = sessionCopy.DB(ConstKey.DBname).C(ConstKey.DeviceCollectionName).UpdateId(deviceToAdd.Id, deviceToAdd)
	return
}

func FindeDeviceByName(name string, Session *mgo.Session) (err error, Device models.DeviceInDB) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err = sessionCopy.DB(ConstKey.DBname).C(ConstKey.DeviceCollectionName).Find(bson.M{"devicename": name}).One(&Device)
	return
}
