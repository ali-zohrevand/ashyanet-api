package DB

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ali-zohrevand/ashyanet-api/Lab/Emqtt/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func CreateMqttUser(user models.MqttUser, Session *mgo.Session) (err error) {

	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	userInDB := models.MqttUserInDB{}
	userInDB.Id = bson.NewObjectId()
	userInDB.Username = user.Username
	userInDB.Is_superuser = user.Is_superuser
	sha := sha256.New()
	sha.Write([]byte(user.Password))
	passByte := sha.Sum(nil)
	passStr := hex.EncodeToString(passByte)
	fmt.Printf("sha256:\t\t%x\n", passByte)
	fmt.Println("pass: ", hex.EncodeToString(passByte))

	userInDB.Password = passStr
	err = sessionCopy.DB("mqtt").C("mqtt_user").Insert(userInDB)

	return
}
