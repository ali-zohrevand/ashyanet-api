package DB

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/pkg/errors"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	. "gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func DeviceCreate(device models.Device, user models.UserInDB, Session *mgo.Session) (err error) {
	err, exist := CheckExist("devicename", device.Name, models.Device{}, DBname, DeviceCollectionName, DeviceExist, Session)
	if exist {
		return
	}
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err, Is := CheckLocationExist(device.Location, sessionCopy)
	if !Is {
		return errors.New(LocationNotFound)
	}

	device.Id = bson.NewObjectId()

	IsValid := CheckKeyIsValid(device.Key, sessionCopy)
	if !IsValid {
		errKeyIsnotValid := errors.New(KeyIsNotValid)
		return errKeyIsnotValid
	}
	err = ActiveKey(device.Key, sessionCopy)
	if err != nil {
		return
	}
	if len(device.Owners) > 0 {
		for _, user := range device.Owners {
			_, err := UserGetByUsername(user, sessionCopy)
			if err != nil {
				err = errors.New(user + ": " + UserNotExist)
				return err
			}
		}
	}
	userFetchedFromDB, err := UserGetByUsername(user.UserName, sessionCopy)
	if err != nil {
		err = errors.New(UserNotExist)

	}
	device.Owners = append(device.Owners, userFetchedFromDB.UserName)
	device.Owners = append(device.Owners, user.UserName)
	device.Owners, _ = DeleteRepetedCell(device.Owners)
	device.Pubsub, _ = DeleteRepetedCell(device.Pubsub)
	device.Publish, _ = DeleteRepetedCell(device.Publish)
	device.Subscribe, _ = DeleteRepetedCell(device.Subscribe)
	userDevice :=models.UserDevice{}
	userDevice.UserName=user.UserName
	userDevice.DeviceName=device.Name
	//.......................................................................................
	commands, data, err := UserGetAllCommandData(user.UserName, sessionCopy)

	for _, data := range device.MqttData {
		data.Name = device.Name + "-"+data.Name
	}
	for _, command := range device.MqttCommand {
		command.Name = device.Name + "-"+command.Name
	}
	for _, command := range device.MqttCommand {
		for _, c := range commands {
			if c == command.Name{
				return errors.New(CommandExist)
			}
		}
	}
	for _, dataObj := range device.MqttData {
		for _, c := range data {
			if dataObj.Name==c{
				return errors.New(DataExist)
			}
		}
	}
	//..........................................ADD mqtt user ...............................
	var mqttUser models.MqttUser
	mqttUser.Username = device.Name
	mqttUser.Password = device.MqttPassword
	sha := sha256.New()
	sha.Write([]byte(device.MqttPassword))
	passByte := sha.Sum(nil)
	passStr := hex.EncodeToString(passByte)
	device.MqttPassword = passStr
	mqttUser.Is_superuser = false
	mqttUser.Created = time.Now().String()
	errCreateMqttUser := EmqttCreateUser(mqttUser, sessionCopy)
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
	for _, c := range device.MqttCommand {
		acl.Subscribe = append(acl.Subscribe, c.Topic)
	}
	for _, c := range device.MqttData {
		acl.Publish = append(acl.Publish, c.Topic)
	}
	acl.Subscribe, _ = DeleteRepetedCell(acl.Subscribe)
	acl.Publish, _ = DeleteRepetedCell(acl.Publish)
	acl.Pubsub, _ = DeleteRepetedCell(acl.Pubsub)
	errCreatACL := EmqttCreateAcl(acl, sessionCopy)
	if errCreatACL != nil {
		return errors.New("INTERNAL ERROR")
	}
	err = sessionCopy.DB(DBname).C(DeviceCollectionName).Insert(device)
	if err!=nil{
		return err
	}
	erraddUserDevice:=AddUserDevice(userDevice,sessionCopy)
	if erraddUserDevice!=nil{
		return erraddUserDevice
	}
	return
}

func IsOwnerOfDevice(username string, deviceName string, Session *mgo.Session) (Is bool, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	Is = false
	err, device := DeviceGetByName(deviceName, sessionCopy)
	for _, user := range device.Owners {
		if user == username {
			Is = true
		}
	}
	return
}

func CreateDeviceWithOutUser(device models.Device, Session *mgo.Session) (err error) {
	//err,exist:=DeviceGetByName(device.Name,Session)
	err, exist := CheckExist("devicename", device.Name, models.Device{}, DBname, DeviceCollectionName, DeviceExist, Session)
	if exist {
		return
	}
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	var DeviceDB = models.Device{}
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

func DeviceGetByName(name string, Session *mgo.Session) (err error, Device models.Device) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err = sessionCopy.DB(DBname).C(DeviceCollectionName).Find(bson.M{"devicename": name}).One(&Device)
	return
}
func DeviceGetById(id string, Session *mgo.Session) ( Device models.Device,err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	deviceId:=bson.ObjectIdHex(id)
	err = sessionCopy.DB(DBname).C(DeviceCollectionName).FindId(deviceId).One(&Device)
	return
}
func DevicesGetAll(Session *mgo.Session) (err error, Device []models.Device) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err = sessionCopy.DB(DBname).C(DeviceCollectionName).Find(bson.M{}).All(&Device)
	if err != nil {
		err = errors.New(DeviceNotExist)
		return
	}
	return
}
func DevicesGetAllByUsername(username string,Session *mgo.Session) ( Device []models.Device,err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err = sessionCopy.DB(DBname).C(DeviceCollectionName).Find(bson.M{"owner":username}).All(&Device)
	if err != nil {
		err = errors.New(DeviceNotExist)
		return
	}

	return
}

func DeviceGetAllTopic(deviceName string, Type string, Session *mgo.Session) (TopicList []string, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err, device := DeviceGetByName(deviceName, sessionCopy)
	if err != nil {
		return
	}
	for _, topic := range device.Pubsub {
		TopicList = append(TopicList, topic)
	}
	switch Type {
	case "sub":
		for _, topic := range device.Subscribe {
			TopicList = append(TopicList, topic)
		}
	case "pub":
		for _, topic := range device.Publish {
			TopicList = append(TopicList, topic)
		}
	case "pubsub":
		for _, topic := range device.Subscribe {
			TopicList = append(TopicList, topic)
		}
		for _, topic := range device.Publish {
			TopicList = append(TopicList, topic)
		}
	default:
		for _, topic := range device.Subscribe {
			TopicList = append(TopicList, topic)
		}
		for _, topic := range device.Publish {
			TopicList = append(TopicList, topic)
		}
	}
	return
}
func DeviceAddCommandByDeviceName(deviceName string,command models.Command,Session *mgo.Session)(err error)  {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err, device := DeviceGetByName(deviceName, sessionCopy)
	if err != nil {
		return
	}
	for _, c := range device.MqttCommand {
		if c.Name== command.Name{
			return errors.New(CommandExist)
		}
	}
	device.MqttCommand = append(device.MqttCommand, command)
	err = sessionCopy.DB(DBname).C(DeviceCollectionName).UpdateId(device.Id,device)
	return


}
func DeviceAddDataByDeviceName(deviceName string,data models.Data,Session *mgo.Session)(err error)  {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err, device := DeviceGetByName(deviceName, sessionCopy)
	if err != nil {
		return
	}
	for _, c := range device.MqttData {
		if c.Name== data.Name{
			return errors.New(DataExist)
		}
	}
	device.MqttData = append(device.MqttData, data)
	err = sessionCopy.DB(DBname).C(DeviceCollectionName).UpdateId(device.Id,device)
	return
}
func DeviceUpdateCommandByDeviceName(deviceName string,commmand models.Command,Session *mgo.Session)(err error)  {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err, device := DeviceGetByName(deviceName, sessionCopy)
	commandExist:=false
	if err != nil {
		return
	}
	for _, v := range device.MqttCommand {
		if v.Name == commmand.Name{
			commandExist=true
			v.Dsc=commmand.Dsc
			v.Topic = commmand.Topic
			v.Value = commmand.Value
		}
	}
	if !commandExist{
		return errors.New(CommandDataNotFOUND)
	}
	err = sessionCopy.DB(DBname).C(DeviceCollectionName).UpdateId(device.Id,device)
	return
}
func DeviceDeleteByName(deviceName string,Session *mgo.Session)(err error)  {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err, device := DeviceGetByName(deviceName, sessionCopy)
	if err != nil {
		return
	}
	err = sessionCopy.DB(DBname).C(DeviceCollectionName).RemoveId(device.Id)
	return
}

func DeviceDeleteById(id string,Session *mgo.Session)(err error)  {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	deviceId:=bson.ObjectIdHex(id)
	err = sessionCopy.DB(DBname).C(DeviceCollectionName).RemoveId(deviceId)
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
