package controllers

import (
	"gitlab.com/hooshyar/ChiChiNi-API/core/DB"
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"net/http"
	"strings"
)

func GetUserFromHeader(req *http.Request) (User models.UserInDB) {
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
		user := DB.GetUserOfSession(jwtToken, session)
		return user
	}
	return
}
