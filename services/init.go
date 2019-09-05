package services


func InitServices() {
	//TODO delete createSetting File
	CreateSettingsFile()
	SaveSetingsInDB()
	CreateDefaultKey()
	AddTempData()
	AddInitUser()
	AddPermissionModelToDB()
	go MqttSubcribeRootTopic()
}
