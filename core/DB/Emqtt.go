package DB

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/settings/Words"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func EmqttCreateUser(user models.MqttUser, Session *mgo.Session) (err error) {

	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	//...............
	_, exist := CheckExist(Words.MqttUserName, user.Username, models.MqttUserInDB{}, Words.EmqttDBName, Words.EmqttUserColletionName, Words.UserNotExist, sessionCopy)
	if exist {
		err = errors.New(Words.UserExist)
		return
	}
	//...............
	userInDB := models.MqttUserInDB{}
	userInDB.Id = bson.NewObjectId()
	userInDB.Username = user.Username
	userInDB.Is_superuser = user.Is_superuser
	userInDB.Created = user.Created
	sha := sha256.New()
	sha.Write([]byte(user.Password))
	passByte := sha.Sum(nil)
	passStr := hex.EncodeToString(passByte)
	userInDB.Password = passStr
	err = sessionCopy.DB(Words.EmqttDBName).C(Words.EmqttUserColletionName).Insert(userInDB)

	return
}
func EmqttGetUserByUserName(username string, Session *mgo.Session) (emqttUser models.MqttUserInDB, err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err = sessionCopy.DB(Words.EmqttDBName).C(Words.EmqttUserColletionName).Find(bson.M{Words.MqttUserName: username}).One(&emqttUser)
	return
}
func EmqttDeleteUser(username string, Session *mgo.Session) (err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err = sessionCopy.DB(Words.EmqttDBName).C(Words.EmqttUserColletionName).Remove(bson.M{Words.MqttUserName: username})

	return
}
func EmqttCreateAcl(acl models.MqttAcl, Session *mgo.Session) (err error) {

	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	//...............
	_, exist := CheckExist(Words.MqttUserName, acl.Username, models.MqttAclIndb{}, Words.EmqttDBName, Words.EmqttAclColectionName, Words.UserNotExist, sessionCopy)
	if exist {
		err = errors.New(Words.UserExist)
		return
	}
	_, MqttuserExist := CheckExist(Words.MqttUserName, acl.Username, models.MqttUserInDB{}, Words.EmqttDBName, Words.EmqttUserColletionName, Words.UserNotExist, sessionCopy)
	if !MqttuserExist {
		err = errors.New(Words.MqttUserNotFound)
		return
	}
	//...............
	aclInDB := models.MqttAclIndb{}
	aclInDB.Id = bson.NewObjectId()
	aclInDB.Username = acl.Username
	aclInDB.Clientid = acl.Clientid
	for i := 0; i < len(acl.Publish); i++ {
		aclInDB.Publish = append(aclInDB.Publish, acl.Publish[i])
	}
	for i := 0; i < len(acl.Subscribe); i++ {
		aclInDB.Subscribe = append(aclInDB.Subscribe, acl.Subscribe[i])
	}
	for i := 0; i < len(acl.Pubsub); i++ {
		aclInDB.Pubsub = append(aclInDB.Pubsub, acl.Pubsub[i])
	}
	err = sessionCopy.DB(Words.EmqttDBName).C(Words.EmqttAclColectionName).Insert(aclInDB)

	return
}
func EmqttDeleteByUsername(username string, Session *mgo.Session) (err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err = sessionCopy.DB(Words.EmqttDBName).C(Words.EmqttUserColletionName).Remove(bson.M{Words.MqttUserName: username})
	err = sessionCopy.DB(Words.EmqttDBName).C(Words.EmqttAclColectionName).Remove(bson.M{Words.MqttUserName: username})

	return
}
