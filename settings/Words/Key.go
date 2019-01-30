package Words

var RuneCharInKey = "123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var LengthOfDeviceKey = 20
var KeyExist = "key Exist"

var StatusActivated = "Acticated"
var StatusValid = "Valid"
var KeyIsNotValid = "Key is not Valid"
var KeyAddedTodevice = "Key added to Device."

type WorldsKey struct {
	RuneCharInKey     string
	LengthOfDeviceKey int
	KeyExist          string
	StatusActivated   string
	StatusValid       string
	KeyIsNotValid     string
	KeyAddedTodevice  string
	TokenKey          string
}

var TokenKey = "42isThdjfhjkhfjksdhfkjsdfhkjkhjkhkjhjkhkhkhkjhjkhkhkjhkjjkheAnswer"
