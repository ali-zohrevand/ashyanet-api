package DB

import (
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/settings/Words"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type JwtSessionDataStore struct {
}

func (ds *JwtSessionDataStore) JwtCreate(Session models.JwtSession, dbSession *mgo.Session) (err error) {
	sessionCopy := dbSession.Copy()
	defer sessionCopy.Close()

	err = sessionCopy.DB(Words.DBname).C(Words.JwtColletionName).Insert(Session)
	return
}

func IsUserSession(username string, token string, Ip string, dbSession *mgo.Session) (found bool) {
	found = false
	sessionCopy := dbSession.Copy()
	defer sessionCopy.Close()
	var JwtSession = models.JwtSession{}
	err := sessionCopy.DB(Words.DBname).C(Words.JwtColletionName).Find(bson.M{"token": token}).One(&JwtSession)
	if err != nil {
		return
	}
	if JwtSession.OwnerUsername == username {
		found = true
	}
	return
}
func JwtGetUser(token string, dbSession *mgo.Session) (user models.UserInDB, err error) {
	sessionCopy := dbSession.Copy()
	defer sessionCopy.Close()
	var JwtSession = models.JwtSession{}
	errHwtSessionL := sessionCopy.DB(Words.DBname).C(Words.JwtColletionName).Find(bson.M{"token": token}).One(&JwtSession)
	if errHwtSessionL != nil {
		return user, errHwtSessionL
	}
	user, err = UserGetByUsername(JwtSession.OwnerUsername, sessionCopy)
	return user, err
}
