package DB

import (
	"errors"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
)

func CreateDeviceKey(Session *mgo.Session) (err error) {
	var deviceKey models.DeviceKey
	deviceKey.Key = GenerateKey()
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err, exist := CheckExist("key", deviceKey.Key, models.DeviceKey{}, Words.DBname, Words.DeviceKeyLocationName, Words.KeyExist, sessionCopy)
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
	DeviceKeyDB.Status = Words.StatusValid
	err = sessionCopy.DB(Words.DBname).C(Words.DeviceKeyLocationName).Insert(DeviceKeyDB)
	return
}
func GetValidKey(Session *mgo.Session) (deviceKey models.DeviceKey) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	sessionCopy.DB(Words.DBname).C(Words.DeviceKeyLocationName).Find(bson.M{"status": Words.StatusValid}).One(&deviceKey)
	return deviceKey
}
func AddKeyToDevice(deviceKey models.DeviceKey, Session *mgo.Session) (err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	//....................................Check Key is Valid ...................................
	var keyFound models.DeviceKeyInDB
	err = sessionCopy.DB(Words.DBname).C(Words.DeviceKeyLocationName).Find(bson.M{"key": deviceKey.Key}).One(&keyFound)
	if err != nil {
		return
	}
	if keyFound.Status != Words.StatusValid {
		errKeyIsNotValid := errors.New(Words.KeyIsNotValid)
		return errKeyIsNotValid
	}
	//...........................Check if Device is available ....................................
	errDevice, deviceToAdd := FindDeviceByName(deviceKey.Device, sessionCopy)
	if errDevice != nil {
		errDeviceNotFound := errors.New(Words.DeviceNotExist)
		return errDeviceNotFound
	}
	//.................................................................................
	deviceToAdd.Key = keyFound.Key
	err = sessionCopy.DB(Words.DBname).C(Words.DeviceCollectionName).UpdateId(deviceToAdd.Id, deviceToAdd)
	keyFound.Device = deviceToAdd
	keyFound.Status = Words.StatusActivated
	err = sessionCopy.DB(Words.DBname).C(Words.DeviceKeyLocationName).UpdateId(keyFound.Id, keyFound)
	CreateDeviceKey(sessionCopy)
	return err
}
func GenerateKey() (key string) {
	var letterRunes = []rune(Words.RuneCharInKey)
	b := make([]rune, Words.LengthOfDeviceKey)
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
	err := sessionCopy.DB(Words.DBname).C(Words.DeviceKeyLocationName).Find(bson.M{"key": key}).One(&keyFound)
	if err == nil && keyFound.Status == Words.StatusValid {
		IsValid = true
		return
	}
	return
}
