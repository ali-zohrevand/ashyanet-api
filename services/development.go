package services

import (
	"gitlab.com/hooshyar/ChiChiNi-API/core/DB"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/services/log"
)

func AddTempData() (err error) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		panic(errConnectDB)
	}
	//...........................User............................
	Ali := models.User{}
	Ali.UserName = "ali"
	Ali.Email = "ali@a.ir"
	Ali.Password = "123456789"
	Ali.Role = "user"
	Ali.FirstName = "ali"
	Ali.LastName = "zohrevand"
	Hasan := models.User{}
	Hasan.UserName = "hasan"
	Hasan.Email = "Hasan@a.ir"
	Hasan.Password = "123456789"
	Hasan.Role = "user"
	Hasan.FirstName = "Hasan"
	Hasan.LastName = "zohrevand"
	userDb := DB.UserDataStore{}
	err = userDb.CreateUser(Ali, session)
	err = userDb.CreateUser(Hasan, session)
	//...........................Device............................
	Lamp := models.Device{}
	Lamp.Name = "lamp"
	Lamp.Description = "لامپ داخل اتاقل "
	Lamp.Key = getValidKey()
	Lamp.Type = "light"
	Lamp.Owners = append(Lamp.Owners, Ali.UserName)
	Lamp.Location = "room"
	MovementSensor := models.Device{}
	MovementSensor.Name = "MovementSensor"
	MovementSensor.Description = " سنسور حرکتی"
	MovementSensor.Key = getValidKey()
	MovementSensor.Type = "movement"
	MovementSensor.Owners = append(Lamp.Owners, Ali.UserName)
	Lamp.Location = "room"
	err = DB.CreateDeviceWithOutUser(Lamp, session)
	err = DB.CreateDeviceWithOutUser(MovementSensor, session)

	//............................Location...........................

	HomeLoattion := models.Location{}
	HomeLoattion.Name = "home"
	HomeLoattion.Description = "home base, tehran"
	RoomLocation := models.Location{}
	RoomLocation.Name = "room"
	RoomLocation.Description = "اتاق بچه ها"
	RoomLocation.Parent = "home"
	RoomLocation.Devices = append(RoomLocation.Devices, Lamp.Name, MovementSensor.Name)
	err = DB.CreateLocation(HomeLoattion, session)
	err = DB.CreateLocation(RoomLocation, session)

	return err
}
func getValidKey() string {
	session, errConnectDB := DB.ConnectDB()
	defer session.Close()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
	}
	key, _ := DB.GetValidKey(session)
	return key.Key
}
