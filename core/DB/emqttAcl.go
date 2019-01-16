package DB

import (
	"errors"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/settings/ConstKey"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func CreateMqattAcl(acl models.MqttAcl, Session *mgo.Session) (err error) {

	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	//...............
	_, exist := CheckExist(ConstKey.MqttUserName, acl.Username, models.MqttAclIndb{}, ConstKey.EmqttDBName, ConstKey.EmqttAclColectionName, ConstKey.UserNotExist, sessionCopy)
	if exist {
		err = errors.New(ConstKey.UserExist)
		return
	}
	_, MqttuserExist := CheckExist(ConstKey.MqttUserName, acl.Username, models.MqttUserInDB{}, ConstKey.EmqttDBName, ConstKey.EmqttUserColletionName, ConstKey.UserNotExist, sessionCopy)
	if !MqttuserExist {
		err = errors.New(ConstKey.MqttUserNotFound)
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
	err = sessionCopy.DB(ConstKey.EmqttDBName).C(ConstKey.EmqttAclColectionName).Insert(aclInDB)

	return
}
func AddS() {

}
