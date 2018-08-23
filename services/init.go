package services

func InitServices() {
	CreateDefaultKey()
	AddTempData()
	AddInitUser()
	AddPermissionModelToDB()
}
