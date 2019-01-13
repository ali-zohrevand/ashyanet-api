package DB

import (
	"errors"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type UserDataStore struct {
}

func (ds *UserDataStore) CheckUserPassCorrect(user models.User, Session *mgo.Session) (IsUserPassValid bool) {
	IsUserPassValid = false
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	var Userdb = models.UserInDB{}
	err := sessionCopy.DB(Words.DBname).C(Words.UserCollectionName).Find(bson.M{"username": user.UserName}).One(&Userdb)
	if err != nil {
		return
	}
	errInComparePassword := bcrypt.CompareHashAndPassword(Userdb.Password, []byte(user.Password))
	if errInComparePassword == nil {
		IsUserPassValid = true
	}
	return
}
func (ds *UserDataStore) CreateUser(userToCreate models.User, Session *mgo.Session) (err error) {
	settings, errLoadSettings := LoadSettings(Session)
	if errLoadSettings != nil  || settings.Type==""{
		return errors.New("setting Not Exist")
	}
	userBack, err := FindUserByEmail(userToCreate.Email, Session)
	if userBack.UserName != "" {
		err = errors.New(Words.UserExist)
		return
	}
	userBack, err = UserGetByUsername(userToCreate.UserName, Session)
	if userBack.Email != "" {
		err = errors.New(Words.UserExist)
		return
	}
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	var Userdb = models.UserInDB{}
	Userdb.Id = bson.NewObjectId()
	Userdb.UserName = userToCreate.UserName
	Userdb.Email = userToCreate.Email
	Userdb.FirstName = userToCreate.FirstName
	Userdb.LastName = userToCreate.LastName
	Userdb.Role = "user"
	if settings.Type=="server"{
		Userdb.Active =false
	}else {
		Userdb.Active= true
	}
	Userdb.TempKeyGenreated = GenerateRandomString(10)
	Userdb.TimeTempKeyGenreated = time.Now().Unix()
	Userdb.Password, err = bcrypt.GenerateFromPassword([]byte(userToCreate.Password), bcrypt.MinCost)
	if err != nil {
		return
	}
	err = sessionCopy.DB(Words.DBname).C(Words.UserCollectionName).Insert(Userdb)
	return
}
func FindUserByEmail(email string, Session *mgo.Session) (userToCreate models.UserInDB, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err = sessionCopy.DB(Words.DBname).C(Words.UserCollectionName).Find(bson.M{"email": email}).One(&userToCreate)
	return
}
func UserGetByUsername(username string, Session *mgo.Session) (userToCreate models.UserInDB, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err = sessionCopy.DB(Words.DBname).C(Words.UserCollectionName).Find(bson.M{"username": username}).One(&userToCreate)
	return
}
func UserActiveBuUsername(username string, Session *mgo.Session)(success bool,err error)  {
	var user models.UserInDB
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err = sessionCopy.DB(Words.DBname).C(Words.UserCollectionName).Find(bson.M{"username": username}).One(&user)
	if err!=nil{
		return false,err
	}
	user.Active = true
	err = sessionCopy.DB(Words.DBname).C(Words.UserCollectionName).UpdateId(user.Id,user)
	if err!=nil{
		return false,err
	}
	return true,err
}
func UserGetAllTopic(username string, Type string, Session *mgo.Session) (TopicList []string, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	user, err := UserGetByUsername(username, sessionCopy)
	if err != nil {
		return
	}
	for _, device := range user.Devices {
		TopicListTemp, errGetTopic := DeviceGetAllTopic(device, Type, sessionCopy)
		if errGetTopic != nil {
			return nil, errGetTopic
		}
		TopicList = append(TopicList, TopicListTemp...)
	}
	return
}
func UserGetAllMqttCommand(username string, Session *mgo.Session) (CommandList []models.Command, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	user, err := UserGetByUsername(username, sessionCopy)
	if err != nil {
		return
	}
	for _, deviceName := range user.Devices {
		errGetDevice, device := DeviceGetByName(deviceName, sessionCopy)
		if errGetDevice != nil {
			return nil, errGetDevice
		}
		CommandList = append(CommandList, device.MqttCommand...)
	}
	return
}
func UserGetAllMqttData(username string, Session *mgo.Session) (DataList []models.Data, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	user, err := UserGetByUsername(username, sessionCopy)
	if err != nil {
		return
	}
	for _, deviceName := range user.Devices {
		errGetDevice, device := DeviceGetByName(deviceName, sessionCopy)
		if errGetDevice != nil {
			return nil, errGetDevice
		}
		DataList = append(DataList, device.MqttData...)
	}
	return
}
func UserHasMqttCommand(username string, Command models.Command, Session *mgo.Session) (Has bool, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	Has = false
	AllCommand, err := UserGetAllMqttCommand(username, sessionCopy)
	if err != nil {
		return false, err
	}
	for _, cm := range AllCommand {
		if cm.Topic == Command.Topic && cm.Value == Command.Value {
			Has = true
		}
	}
	return
}
func UserGetMqttCommandByName(username string, CommandName string, Session *mgo.Session) (data models.Command, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	AllCommand, err := UserGetAllMqttCommand(username, sessionCopy)
	if err != nil {
		return data, err
	}
	for _, cm := range AllCommand {
		if cm.Name == CommandName {
			return cm, nil
		}
	}
	return
}
func UserGetMqttDataByName(username string, DataName string, Session *mgo.Session) (data models.Data, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	AllData, err := UserGetAllMqttData(username, sessionCopy)
	if err != nil {
		return data, err
	}
	for _, cm := range AllData {
		if cm.Name == DataName {
			return cm, nil
		}
	}
	return
}
func UserGetAllDevice(username string, Session *mgo.Session) (Devices []models.Device, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	UserIndb, err := UserGetByUsername(username, sessionCopy)
	if err != nil {
		return
	}
	for _, deviceName := range UserIndb.Devices {
		errGetdevice, deviceIndb := DeviceGetByName(deviceName, sessionCopy)
		if errGetdevice != nil {
			return nil, errGetdevice
		}
		Devices = append(Devices, deviceIndb)
	}
	return
}
