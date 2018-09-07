package services

import (
	"encoding/json"
	"gitlab.com/hooshyar/ChiChiNi-API/core/DB"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"gitlab.com/hooshyar/ChiChiNi-API/settings/Words"
)

func CreateSettingsFile() (err error) {
	if IsFileExist(Words.SettingPath) {
		return
	}
	setting := models.SettingsInDB{}
	setting.Username = "testUser"
	session, err := DB.ConnectDB()
	if err != nil {
		return
	}
	defer session.Close()
	key, err := DB.GetValidKey(session)
	if err != nil {
		return
	}
	setting.Key = key.Key
	setting.Password = "123456"
	setting.Identifier = DB.GenerateKey()
	setting.Type = "gateway"
	settingJsonByte, err := json.Marshal(setting)
	if err != nil {
		return
	}
	err = WriteFile(Words.SettingPath, string(settingJsonByte))
	return

}
func SaveSetingsInDB() (err error) {
	//setting := models.SettingsInDB{}
	session, err := DB.ConnectDB()
	if err != nil {
		return
	}
	defer session.Close()
	Is := DB.IsCollectionEmptty(Words.DBname, Words.SettingsCollectiName, session)
	if !Is {
		return
	}
	content, er := ReadFile(Words.SettingPath)
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
func GetServerSettings() (settings models.SettingsInDB, err error) {
	session, err := DB.ConnectDB()
	if err != nil {
		return
	}
	defer session.Close()
	settings, err = DB.LoadSettings(session)
	if err != nil {
		return
	}
	return
}
