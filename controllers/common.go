package controllers

import (
	"github.com/ali-zohrevand/ashyanet-api/core/DB"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"net/http"
	"strings"
)

func GetUserFromHeader(req *http.Request) (User models.UserInDB, err error) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		return
	}
	authToken := req.Header.Get("Authorization")
	authArr := strings.Split(authToken, " ")
	if len(authArr) != 2 {
		return
	} else {
		jwtToken := authArr[1]
		user, err := DB.JwtGetUser(jwtToken, session)
		return user, err
	}
	return
}
