package DB

import (
	"errors"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func CreateMqttAcl(acl models.MqttAcl, Session *mgo.Session) (err error) {

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
func AddS() {

}
