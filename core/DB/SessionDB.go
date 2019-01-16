package DB

import (
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/settings/ConstKey"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type JwtSessionDataStore struct {
}

func (ds *JwtSessionDataStore) CreateJwtSession(Session models.JwtSession, dbSession *mgo.Session) (err error) {
	sessionCopy := dbSession.Copy()
	defer sessionCopy.Close()

	err = sessionCopy.DB(ConstKey.DBname).C(ConstKey.JwtColletionName).Insert(Session)
	return
}

func IsUserSession(username string, token string, Ip string, dbSession *mgo.Session) (found bool) {
	found = false
	sessionCopy := dbSession.Copy()
	defer sessionCopy.Close()
	var JwtSession = models.JwtSession{}
	err := sessionCopy.DB(ConstKey.DBname).C(ConstKey.JwtColletionName).Find(bson.M{"token": token}).One(&JwtSession)
	if err != nil {
		return
	}
	if JwtSession.OwnerUsername == username {
		found = true
	}
	return
}
func GetUserOfSession(token string, dbSession *mgo.Session) string {
	sessionCopy := dbSession.Copy()
	defer sessionCopy.Close()
	var JwtSession = models.JwtSession{}
	sessionCopy.DB(ConstKey.DBname).C(ConstKey.JwtColletionName).Find(bson.M{"token": token}).One(&JwtSession)
	user, _ := FindUserByUsername(JwtSession.OwnerUsername, sessionCopy)
	return user.Role
}
