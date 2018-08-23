package DB

import (
	"errors"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/ConstKey"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
)

func CreateDeviceKey(Session *mgo.Session) (err error) {
	var deviceKey models.DeviceKey
	deviceKey.Key = GenerateKey()
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err, exist := CheckExist("key", deviceKey.Key, models.DeviceKey{}, ConstKey.DBname, ConstKey.DeviceKeyLocationName, ConstKey.KeyExist, sessionCopy)
	if exist {
		return
	}
	var DeviceKeyDB = models.DeviceKeyInDB{}
	DeviceKeyDB.Id = bson.NewObjectId()
	DeviceKeyDB.Key = deviceKey.Key
	TempDevice := models.DeviceInDB{}
	TempDevice.Id = bson.NewObjectId()
	TempDevice.Name = "temp"
	DeviceKeyDB.Device = TempDevice
	DeviceKeyDB.Status = ConstKey.StatusValid
	err = sessionCopy.DB(ConstKey.DBname).C(ConstKey.DeviceKeyLocationName).Insert(DeviceKeyDB)
	return
}
func GetValidKey(Session *mgo.Session) (deviceKey models.DeviceKey) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	sessionCopy.DB(ConstKey.DBname).C(ConstKey.DeviceKeyLocationName).Find(bson.M{"status": ConstKey.StatusValid}).One(&deviceKey)
	return deviceKey
}
func AddKeyToDevice(deviceKey models.DeviceKey, Session *mgo.Session) (err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	//....................................Check Key is Valid ...................................
	var keyFound models.DeviceKeyInDB
	err = sessionCopy.DB(ConstKey.DBname).C(ConstKey.DeviceKeyLocationName).Find(bson.M{"key": deviceKey.Key}).One(&keyFound)
	if err != nil {
		return
	}
	if keyFound.Status != ConstKey.StatusValid {
		errKeyIsNotValid := errors.New(ConstKey.KeyIsNotValid)
		return errKeyIsNotValid
	}
	//...........................Check if Device is available ....................................
	errDevice, deviceToAdd := FindeDeviceByName(deviceKey.Device, sessionCopy)
	if errDevice != nil {
		errDeviceNotFound := errors.New(ConstKey.DeviceNotExist)
		return errDeviceNotFound
	}
	//.................................................................................
	deviceToAdd.Key = keyFound.Key
	err = sessionCopy.DB(ConstKey.DBname).C(ConstKey.DeviceCollectionName).UpdateId(deviceToAdd.Id, deviceToAdd)
	keyFound.Device = deviceToAdd
	keyFound.Status = ConstKey.StatusActivated
	err = sessionCopy.DB(ConstKey.DBname).C(ConstKey.DeviceKeyLocationName).UpdateId(keyFound.Id, keyFound)
	CreateDeviceKey(sessionCopy)
	return err
}
func GenerateKey() (key string) {
	var letterRunes = []rune(ConstKey.RuneCharInKey)
	b := make([]rune, ConstKey.LengthOfDeviceKey)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func CheckKeyIsValid(key string, Session *mgo.Session) (IsValid bool) {
	IsValid = false
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	var keyFound models.DeviceKeyInDB
	err := sessionCopy.DB(ConstKey.DBname).C(ConstKey.DeviceKeyLocationName).Find(bson.M{"key": key}).One(&keyFound)
	if err == nil && keyFound.Status == ConstKey.StatusValid {
		IsValid = true
		return
	}
	return
}
