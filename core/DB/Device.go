package DB

import (
	"github.com/pkg/errors"
	"gitlab.com/hooshyar/ChiChiNi-API/OutputAPI"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	. "gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func CreateDevice(device models.Device, user models.UserInDB, Session *mgo.Session) (err error) {
	//err,exist:=FindDeviceByName(device.Name,Session)
	err, exist := CheckExist("devicename", device.Name, models.DeviceInDB{}, DBname, DeviceCollectionName, DeviceExist, Session)
	if exist {
		return
	}
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err, Is := CheckLocationExist(device.Location, sessionCopy)
	if !Is {
		return errors.New(LocationNotFound)
	}
	var DeviceDB = models.DeviceInDB{}
	DeviceDB.Id = bson.NewObjectId()
	DeviceDB.Location = device.Location
	DeviceDB.Name = device.Name
	DeviceDB.Description = device.Description

	if device.Key != "" {
		IsValid := CheckKeyIsValid(device.Key, sessionCopy)
		if !IsValid {
			errKeyIsnotValid := errors.New(KeyIsNotValid)
			return errKeyIsnotValid
		}
	}
	err = ActiveKey(device.Key, sessionCopy)
	if err != nil {
		return
	}
	DeviceDB.Key = device.Key
	if len(device.Owners) > 0 {
		for _, user := range device.Owners {
			userFetchedFromDB, err := UserGetByUsername(user, sessionCopy)
			if err != nil {
				err = errors.New(user + ": " + UserNotExist)
				return err
			}
			DeviceDB.Owners = append(DeviceDB.Owners, userFetchedFromDB.UserName)
		}
	}
	userFetchedFromDB, err := UserGetByUsername(DeafualtAdmminUserName, sessionCopy)
	DeviceDB.Owners = append(DeviceDB.Owners, userFetchedFromDB.UserName)
	DeviceDB.Type = device.Type
	for _, p := range device.Pubsub {
		DeviceDB.Pubsub = append(DeviceDB.Pubsub, p)
	}
	for _, p := range device.Publish {
		DeviceDB.Publish = append(DeviceDB.Publish, p)
	}
	for _, p := range device.Subscribe {
		DeviceDB.Subscribe = append(DeviceDB.Subscribe, p)
	}
	for _, p := range device.Command {
		DeviceDB.Command = append(DeviceDB.Command, p)
	}
	for _, p := range device.Data {
		DeviceDB.Data = append(DeviceDB.Data, p)
	}
	DeviceDB.Pubsub, _ = DeleteRepetedCell(DeviceDB.Pubsub)
	DeviceDB.Publish, _ = DeleteRepetedCell(DeviceDB.Publish)
	DeviceDB.Subscribe, _ = DeleteRepetedCell(DeviceDB.Subscribe)

	//..........................................ADD mqtt user ...............................
	var mqttUser models.MqttUser
	mqttUser.Username = device.Name
	mqttUser.Password = device.MqttPassword
	mqttUser.Is_superuser = false
	mqttUser.Created = time.Now().String()
	errCreateMqttUser := CreateMqttUser(mqttUser, sessionCopy)
	if errCreateMqttUser != nil && errCreateMqttUser.Error() != UserExist {
		return errors.New("INTERNAL ERROR")
	}
	//..........................................ADD mqtt acl ...............................
	var acl models.MqttAcl
	acl.Username = device.Name
	acl = addTopicInArraToMqttACL(device.Subscribe, acl, "s")
	acl = addTopicInArraToMqttACL(device.Publish, acl, "p")
	acl = addTopicInArraToMqttACL(device.Pubsub, acl, "ps")
	//Delete Repeated
	for _, c := range device.Command {
		acl.Subscribe = append(acl.Subscribe, c.Topic)
	}
	for _, c := range device.Command {
		acl.Publish = append(acl.Publish, c.Topic)
	}
	acl.Subscribe, _ = DeleteRepetedCell(acl.Subscribe)
	acl.Publish, _ = DeleteRepetedCell(acl.Publish)
	acl.Pubsub, _ = DeleteRepetedCell(acl.Pubsub)
	errCreatACL := CreateMqttAcl(acl, sessionCopy)
	if errCreatACL != nil {
		return errors.New("INTERNAL ERROR")
	}
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
			userFetchedFromDB, err := UserGetByUsername(user, sessionCopy)
			if err != nil {
				err = errors.New(user + ": " + UserNotExist)
				return err
			}
			DeviceDB.Owners = append(DeviceDB.Owners, userFetchedFromDB.UserName)
		}
	}
	userFetchedFromDB, err := UserGetByUsername(DeafualtAdmminUserName, sessionCopy)
	DeviceDB.Owners = append(DeviceDB.Owners, userFetchedFromDB.UserName)
	DeviceDB.Type = device.Type
	err = sessionCopy.DB(DBname).C(DeviceCollectionName).Insert(DeviceDB)
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
func addTopicInArraToMqttACL(array []string, acl models.MqttAcl, TopicType string) models.MqttAcl {

	switch TopicType {
	case "p":
		for _, a := range array {
			acl.Publish = append(acl.Publish, a)
		}
	case "ps":
		for _, a := range array {
			acl.Pubsub = append(acl.Pubsub, a)
		}
	case "s":
		for _, a := range array {
			acl.Subscribe = append(acl.Subscribe, a)
		}

	default:
		for _, a := range array {
			acl.Publish = append(acl.Publish, a)
		}
	}
	return acl
}
