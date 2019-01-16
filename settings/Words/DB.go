package Words

var DBname = "db"
var UserCollectionName = "Users"
var JwtColletionName = "JwtTokenSession"
var PermissionCollectionName = "Permission"
var DeviceCollectionName = "Devices"
var LocationCollectionName = "Locations"
var DeviceKeyLocationName = "Keys"
var SettingsCollectiName = "settings"
var DBNotConnectet = "DB PROBLEM"

type WorldsDB struct {
	DBname                   string
	UserCollectionName       string
	JwtColletionName         string
	PermissionCollectionName string
	DeviceCollectionName     string
	LocationCollectionName   string
	DeviceKeyLocationName    string
	SettingsCollectiName     string
	DBNotConnectet           string
}
