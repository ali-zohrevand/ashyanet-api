package Words

var FirstNameNeeded = "First name Could not be empty"
var UserExist = "User Exist"
var UserCreated = "User Created"
var UserVerifyMailProblem = "Verify mail did not sent, tray again."
var VerifyMailSent = "Verify Mail sent, please check your mail."
var UserNotExist = "User Not Exist"
var TimeExpired = "Time Expired."
var UserActivated = "User Activated."
var UserNotActive = "User Not Active."

var LocationExist = "Location Exist"
var DeviceCreated = "Device Created"
var LocationCreated = "Location Created"
var DeviceOrUserNotFound = "Device Or User Not Found"
var UserAddedToDevice = "User Added To Device"
var UserAddedToLocation = "User Added To Location"
var LocationNotFound = "Location Not Found"
var UserOrLocationNotFound = "user Or Location Not Found"

type WordsValidation struct {
	FirstNameNeeded        string
	UserExist              string
	UserCreated            string
	UserNotExist           string
	DeviceNotExist         string
	DeviceExist            string
	LocationExist          string
	DeviceCreated          string
	LocationCreated        string
	DeviceOrUserNotFound   string
	UserAddedToDevice      string
	LocationNotFound       string
	UserOrLocationNotFound string
	UserAddedToLocation    string
	UserActivated          string
	TimeExpired            string
	UserNotActive          string
	VerifyMailSent         string
	UserVerifyMailProblem  string
}
