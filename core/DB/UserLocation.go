package DB

import (
	"errors"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
	"gopkg.in/mgo.v2"
)

func AddUserLocation(userLcoation models.UserLocation, Session *mgo.Session) (err error) {
	sessionCopy := Session.Copy()
	defer sessionCopy.Close()
	var LocationHasUser = false
	var UserHasLocation = false
	locationInDB, errFoundLocation := LocationGetByName(userLcoation.LocationName, sessionCopy)
	if errFoundLocation != nil {
		return errors.New(Words.LocationNotFound)
	}
	userInDB, errFoundUser := UserGetByUsername(userLcoation.UserName, sessionCopy)
	if errFoundUser != nil {
		return errors.New(Words.UserNotExist)
	}
	for _, location := range userInDB.Locations {
		if location == userLcoation.LocationName {
			UserHasLocation = true
		}
	}
	for _, user := range locationInDB.Users {
		if user == userLcoation.UserName {
			LocationHasUser = true
		}
	}
	if !UserHasLocation {
		userInDB.Locations = append(userInDB.Locations, userLcoation.LocationName)
		err = sessionCopy.DB(Words.DBname).C(Words.UserCollectionName).UpdateId(userInDB.Id, userInDB)

	}
	if !LocationHasUser {
		locationInDB.Users = append(locationInDB.Users, userLcoation.UserName)
		err = sessionCopy.DB(Words.DBname).C(Words.LocationCollectionName).UpdateId(locationInDB.Id, locationInDB)

	}
	return
}
