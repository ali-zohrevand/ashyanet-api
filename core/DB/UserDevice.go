package DB

import (
	"github.com/pkg/errors"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	. "gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func AddUserDevice(userToDevice models.UserDevice, Session *mgo.Session) (err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	var DeviceHaseUser = false
	var UserHasDevice = false
	var user models.UserInDB
	err = sessionCopy.DB(DBname).C(UserCollectionName).Find(bson.M{"username": userToDevice.UserName}).One(&user)
	if err != nil {
		err = errors.New(UserNotExist)
		return
	}
	var deviceToAdd models.Device
	err = sessionCopy.DB(DBname).C(DeviceCollectionName).Find(bson.M{"devicename": userToDevice.DeviceName}).One(&deviceToAdd)
	if err != nil {
		err = errors.New(DeviceNotExist)

		return
	}
	for _, u := range deviceToAdd.Owners {
		if u == userToDevice.UserName {
			DeviceHaseUser = true
		}
	}
	if !DeviceHaseUser {
		deviceToAdd.Owners = append(deviceToAdd.Owners, user.UserName)
		err = sessionCopy.DB(DBname).C(DeviceCollectionName).UpdateId(deviceToAdd.Id, deviceToAdd)
	}
	for _, device := range user.Devices {
		if device == userToDevice.DeviceName {
			UserHasDevice = true
		}
	}
	if !UserHasDevice {
		user.Devices = append(user.Devices, userToDevice.DeviceName)
		err = sessionCopy.DB(DBname).C(UserCollectionName).UpdateId(user.Id, user)

	}
	return
}
