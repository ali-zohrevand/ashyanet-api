package DB

import (
	"errors"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/settings/ConstKey"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserDataStore struct {
}

func (ds *UserDataStore) FindUser(userId string) (user models.User, err error) {
	return
}
func (ds *UserDataStore) CheckUserPassCorrect(user models.User, Session *mgo.Session) (IsUserPassValid bool) {
	IsUserPassValid = false
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	var Userdb = models.UserInDB{}
	err := sessionCopy.DB(ConstKey.DBname).C(ConstKey.UserCollectionName).Find(bson.M{"username": user.UserName}).One(&Userdb)
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
		err = errors.New(ConstKey.UserExist)
		return
	}
	userBack, err = FindUserByUsername(userToCreate.UserName, Session)
	if userBack.Email != "" {
		err = errors.New(ConstKey.UserExist)
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
	err = sessionCopy.DB(ConstKey.DBname).C(ConstKey.UserCollectionName).Insert(Userdb)
	return
}
func FindUserByEmail(email string, Session *mgo.Session) (userToCreate models.UserInDB, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err = sessionCopy.DB(ConstKey.DBname).C(ConstKey.UserCollectionName).Find(bson.M{"email": email}).One(&userToCreate)
	return
}
func FindUserByUsername(username string, Session *mgo.Session) (userToCreate models.UserInDB, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err = sessionCopy.DB(ConstKey.DBname).C(ConstKey.UserCollectionName).Find(bson.M{"username": username}).One(&userToCreate)
	return
}
