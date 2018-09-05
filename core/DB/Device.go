package DB

import (
	"github.com/pkg/errors"
	"gitlab.com/hooshyar/ChiChiNi-API/OutputAPI"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	. "gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func CreateDevice(device models.Device, user models.UserInDB, Session *mgo.Session) (err error) {
	//err,exist:=FindDeviceByName(device.Name,Session)
	err, exist := CheckExist("devicename", device.Name, models.DeviceInDB{}, DBname, DeviceCollectionName, DeviceExist, Session)
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
			errKeyIsnotValid := errors.New(KeyIsNotValid)
			return errKeyIsnotValid
		}
	}
	DeviceDB.Key = device.Key
	if len(device.Owners) > 0 {
		for _, user := range device.Owners {
			userFetchedFromDB, err := FindUserByUsername(user, sessionCopy)
			if err != nil {
				err = errors.New(user + ": " + UserNotExist)
				return err
			}
			DeviceDB.Owners = append(DeviceDB.Owners, userFetchedFromDB)
		}
	}
	userFetchedFromDB, err := FindUserByUsername(DeafualtAdmminUserName, sessionCopy)
	DeviceDB.Owners = append(DeviceDB.Owners, userFetchedFromDB)
	DeviceDB.Type = device.Type
	err = sessionCopy.DB(DBname).C(DeviceCollectionName).Insert(DeviceDB)
	return
}
func CreateDeviceWithOutUser(device models.Device, Session *mgo.Session) (err error) {
	//err,exist:=FindDeviceByName(device.Name,Session)
	err, exist := CheckExist("devicename", device.Name, models.DeviceInDB{}, DBname, DeviceCollectionName, DeviceExist, Session)
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
			errKeyIsnotValid := errors.New(KeyIsNotValid)
			return errKeyIsnotValid
		}
	}
	DeviceDB.Key = device.Key
	if len(device.Owners) > 0 {
		for _, user := range device.Owners {
			userFetchedFromDB, err := FindUserByUsername(user, sessionCopy)
			if err != nil {
				err = errors.New(user + ": " + UserNotExist)
				return err
			}
			DeviceDB.Owners = append(DeviceDB.Owners, userFetchedFromDB)
		}
	}
	userFetchedFromDB, err := FindUserByUsername(DeafualtAdmminUserName, sessionCopy)
	DeviceDB.Owners = append(DeviceDB.Owners, userFetchedFromDB)
	DeviceDB.Type = device.Type
	err = sessionCopy.DB(DBname).C(DeviceCollectionName).Insert(DeviceDB)
	return
}

func AddUserToDevice(userToDevice models.UserDevice, Session *mgo.Session) (err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	var user models.UserInDB
	err = sessionCopy.DB(DBname).C(UserCollectionName).Find(bson.M{"username": userToDevice.UserName}).One(&user)
	if err != nil {
		err = errors.New(UserNotExist)
		return
	}
	var deviceToAdd models.DeviceInDB
	err = sessionCopy.DB(DBname).C(DeviceCollectionName).Find(bson.M{"devicename": userToDevice.DeviceName}).One(&deviceToAdd)
	if err != nil {
		err = errors.New(DeviceNotExist)

		return
	}
	for _, u := range deviceToAdd.Owners {
		if u.UserName == userToDevice.UserName {
			return
		}
	}
	deviceToAdd.Owners = append(deviceToAdd.Owners, user)
	err = sessionCopy.DB(DBname).C(DeviceCollectionName).UpdateId(deviceToAdd.Id, deviceToAdd)
	return
}

func FindDeviceByName(name string, Session *mgo.Session) (err error, Device models.DeviceInDB) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err = sessionCopy.DB(DBname).C(DeviceCollectionName).Find(bson.M{"devicename": name}).One(&Device)
	return
}
func GetAllDevices(Session *mgo.Session) (err error, Device []OutputAPI.Device) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	DevicesInDB := []models.DeviceInDB{}
	err = sessionCopy.DB(DBname).C(DeviceCollectionName).Find(bson.M{}).All(&DevicesInDB)
	if err != nil {
		err = errors.New(DeviceNotExist)
		return
	}

	for i := 0; i < len(DevicesInDB); i++ {
		tempDevice := OutputAPI.Device{}
		tempDevice.Name = DevicesInDB[i].Name
		tempDevice.Location = DevicesInDB[i].Location
		tempDevice.Description = DevicesInDB[i].Description
		tempDevice.Key = DevicesInDB[i].Key
		tempDevice.Id = DevicesInDB[i].Id.Hex()
		//Todo Complete it
	}
	return
}
