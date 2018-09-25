package DB

import (
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type JwtSessionDataStore struct {
}

func (ds *JwtSessionDataStore) CreateJwtSession(Session models.JwtSession, dbSession *mgo.Session) (err error) {
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
func GetUserOfSession(token string, dbSession *mgo.Session) (user models.UserInDB, err error) {
	sessionCopy := dbSession.Copy()
	defer sessionCopy.Close()
	var JwtSession = models.JwtSession{}
	sessionCopy.DB(Words.DBname).C(Words.JwtColletionName).Find(bson.M{"token": token}).One(&JwtSession)
	user, err = UserGetByUsername(JwtSession.OwnerUsername, sessionCopy)
	return user, err
}
