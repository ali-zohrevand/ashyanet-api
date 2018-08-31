package services

import (
	"encoding/json"
	"gitlab.com/hooshyar/ChiChiNi-API/core/DB"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/settings"
)

func CreateSettingsFile() (err error) {
	setting := models.SettingsInDB{}
	setting.Username = "testUser"
	session, err := DB.ConnectDB()
	if err != nil {
		return
	}
	defer session.Close()
	key := DB.GetValidKey(session)
	setting.Key = key.Key
	setting.Password = "123456"
	settingJsonByte, err := json.Marshal(setting)
	if err != nil {
		return
	}
	err = WriteFile(settings.SettingPath, string(settingJsonByte))
	return

}
