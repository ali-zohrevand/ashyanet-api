package DB

import "gopkg.in/mgo.v2"

func ConnectDB() (session *mgo.Session, err error) {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")
	return s, err
}
