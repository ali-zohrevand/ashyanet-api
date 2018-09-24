package DB

import (
	"errors"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
func UserGetAllCommand(username string, Session *mgo.Session) (CommandList []models.Command, err error) {
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
		CommandList = append(CommandList, device.Command...)
	}
	return
}
func UserGetAllData(username string, Type string, Session *mgo.Session) (CommandList []models.Data, err error) {
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
		CommandList = append(CommandList, device.Data...)
	}
	return
}
func UserGetAllDevice(username string, Session *mgo.Session) (Devices []models.DeviceInDB, err error) {
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
