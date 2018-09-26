package DB

import (
	"errors"
	"fmt"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func CheckExist(ObjectName string, name string, Object interface{}, DbNameString string, ColloctionName string, errorString string, Session *mgo.Session) (err error, Exist bool) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	err = sessionCopy.DB(DbNameString).C(ColloctionName).Find(bson.M{ObjectName: name}).One(&Object)
	if Object != nil && err == nil {
		err = errors.New(errorString)
		Exist = true
		return
	}
	Exist = false
	return
}

func addPath(path string, New string) string {
	path = path + "/" + New
	return path
}
func ConnectDB() (session *mgo.Session, err error) {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")
	return s, err
}
func IsCollectionEmptty(DbName string, CollName string, Session *mgo.Session) (Is bool) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	var OBJ []interface{}
	sessionCopy.DB(DbName).C(CollName).Find(bson.M{}).All(&OBJ)

	if len(OBJ) == 0 {
		return true
	}
	return false

}

func IsMongoUp() (Valid bool) {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		fmt.Println(err)
		return false
	}
	err = s.Ping()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func GenerateRandomString(length int) string {
	var charset string
	charset = Words.RuneCharInKey
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
