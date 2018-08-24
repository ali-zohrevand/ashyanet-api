package DB

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/pkg/errors"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/ConstKey"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func CreateMqttUser(user models.MqttUser, Session *mgo.Session) (err error) {

	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	//...............
	_, exist := CheckExist(ConstKey.MqttUserName, user.Username, models.MqttUserInDB{}, ConstKey.EmqttDBName, ConstKey.EmqttUserColletionName, ConstKey.UserNotExist, sessionCopy)
	if exist {
		err = errors.New(ConstKey.UserExist)
		return
	}
	//...............
	userInDB := models.MqttUserInDB{}
	userInDB.Id = bson.NewObjectId()
	userInDB.Username = user.Username
	userInDB.Is_superuser = user.Is_superuser
	sha := sha256.New()
	sha.Write([]byte(user.Password))
	passByte := sha.Sum(nil)
	passStr := hex.EncodeToString(passByte)
	userInDB.Password = passStr
	err = sessionCopy.DB(ConstKey.EmqttDBName).C(ConstKey.EmqttUserColletionName).Insert(userInDB)

	return
}
