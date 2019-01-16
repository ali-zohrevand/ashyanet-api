package validation

import (
	"encoding/json"
	"github.com/ali-zohrevand/ashyanet-api/OutputAPI"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/asaskevich/govalidator"
	"strings"
)

func ObjectValidation(object interface{}) (Out []byte, err error, IsValid bool) {
	var errToShow OutputAPI.Message
	govalidator.TagMap["blacklist"] = govalidator.Validator(ValidBadChar)
	IsValid, err = govalidator.ValidateStruct(object)
	if err != nil {
		errToShow.Error = err.Error()
	}
	Out, _ = json.Marshal(errToShow)
	return
}
func UserLoginValidation(user models.User) (Out []byte, err error, IsValid bool) {
	IsValid = true
	var errToShow OutputAPI.Message
	if !ValidBadChar(user.Password) || !ValidateUserAndPass(user.Password) || !ValidateUserAndPass(user.UserName) {
		IsValid = false
	}
	Out, err = json.Marshal(errToShow)
	return
}
func ValidBadChar(input string) (Valid bool) {
	Valid = true
	slash := `\`
	BadCharList := []string{",", "/", "<", ">", "$", "'", "!", ")", "(", "&", "%", "~", "=", "+", "-", "?"}
	BadCharList = append(BadCharList, slash)
	for _, badChar := range BadCharList {
		if strings.Contains(input, badChar) {
			return false
		}
	}
	return
}
func ValidateUserAndPass(Input string) (IsValid bool) {
	IsValid = true
	if Input == "" {
		return false
	}
	return
}
