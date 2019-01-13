package Words

import (
	"encoding/json"
	"fmt"
)

func InitDbWords() (dbW WorldsDB, err error) {
	dbW = WorldsDB{}
	dbW.DBname = DBname
	dbW.SettingsCollectiName = SettingsCollectiName
	dbW.DeviceCollectionName = DeviceCollectionName
	dbW.UserCollectionName = UserCollectionName
	dbW.JwtColletionName = JwtColletionName
	dbW.PermissionCollectionName = PermissionCollectionName
	dbW.LocationCollectionName = LocationCollectionName
	dbW.DeviceKeyLocationName = DeviceKeyLocationName
	dbW.DBNotConnectet = DBNotConnectet
	return

}
func InitKeyWords() (dbK WorldsKey, err error) {
	dbK = WorldsKey{}
	dbK.KeyIsNotValid = KeyIsNotValid
	dbK.StatusValid = StatusValid
	dbK.StatusActivated = StatusActivated
	dbK.LengthOfDeviceKey = LengthOfDeviceKey
	dbK.RuneCharInKey = RuneCharInKey
	dbK.KeyExist = KeyExist
	dbK.TokenKey = TokenKey
	dbK.KeyAddedTodevice = KeyAddedTodevice
	dbK.LengthOfDeviceKey = LengthOfDeviceKey
	return
}
func InitValidationWords() (dbV WordsValidation, err error) {
	dbV = WordsValidation{}
	dbV.DeviceNotExist = DeviceNotExist
	dbV.DeviceExist = DeviceExist
	dbV.DeviceCreated = DeviceCreated
	dbV.DeviceOrUserNotFound = DeviceOrUserNotFound
	dbV.FirstNameNeeded = FirstNameNeeded
	dbV.LocationCreated = LocationCreated
	dbV.LocationExist = LocationExist
	dbV.LocationNotFound = LocationNotFound
	dbV.UserAddedToDevice = UserAddedToDevice
	dbV.UserCreated = UserCreated
	dbV.UserExist = UserExist
	dbV.UserNotExist = UserNotExist
	dbV.UserOrLocationNotFound = UserOrLocationNotFound
	dbV.UserAddedToLocation = UserAddedToLocation
	dbV.UserActivated = UserActivated
	dbV.TimeExpired = TimeExpired
	dbV.UserNotActive = UserNotActive
	a, err := json.Marshal(dbV)
	fmt.Println(string(a))
	return
}
