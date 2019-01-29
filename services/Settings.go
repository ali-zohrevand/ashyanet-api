package services

import (
	"encoding/json"
	"github.com/ali-zohrevand/ashyanet-api/core/DB"
	"github.com/ali-zohrevand/ashyanet-api/models"

	"github.com/ali-zohrevand/ashyanet-api/services/Tools"
	"github.com/ali-zohrevand/ashyanet-api/settings/Words"
)

func CreateSettingsFile() (err error) {
	if Tools.IsFileExist(Words.SettingPath) {
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
	setting.Type = "server"
	setting.Url = "https://127.0.0.1:5000/active"
	setting.MailHost = "smtp.gmail.com"
	setting.MailPort = "465"
	setting.MailVerifyUsername = "ashyanet@gmail.com"
	setting.MailVerifyPassword = "mahdi1369QWE"
	settingJsonByte, err := json.Marshal(setting)

	if err != nil {
		return
	}
	err = Tools.WriteFile(Words.SettingPath, string(settingJsonByte))
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
	content, er := Tools.ReadFile(Words.SettingPath)
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
