package services

import (
	"encoding/json"
	"gitlab.com/hooshyar/ChiChiNi-API/core/DB"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/ConstKey"
)

func CreateSettingsFile() (err error) {
	if IsFileExist(ConstKey.SettingPath) {
		return
	}
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
	setting.Identifier = DB.GenerateKey()
	settingJsonByte, err := json.Marshal(setting)
	if err != nil {
		return
	}
	err = WriteFile(ConstKey.SettingPath, string(settingJsonByte))
	return

}
func SaveSetingsInDB() (err error) {
	//setting := models.SettingsInDB{}
	session, err := DB.ConnectDB()
	if err != nil {
		return
	}
	defer session.Close()
	Is := DB.IsCollectionEmptty(ConstKey.DBname, ConstKey.SettingsCollectiName, session)
	if !Is {
		return
	}
	content, er := ReadFile(ConstKey.SettingPath)
	if er != nil {
		return
	}
	setting := models.SettingsInDB{}
	err = json.Unmarshal([]byte(content), &setting)
	if er != nil {
		return
	}
	err = DB.SaveSettings(setting, session)
	return err
}
