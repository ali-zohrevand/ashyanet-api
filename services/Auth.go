package services

import (
	"fmt"
	"github.com/ali-zohrevand/ashyanet-api/core/DB"
	"github.com/ali-zohrevand/ashyanet-api/services/log"
	"github.com/casbin/casbin"
	scas "github.com/qiangmzsx/string-adapter"
	"net/http"
	"strings"
)

func RequireTokenAuthentication(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		rw.WriteHeader(http.StatusUnauthorized)
	}
	defer session.Close()
	authToken := req.Header.Get("Authorization")
	authArr := strings.Split(authToken, " ")
	if len(authArr) != 2 {
		rw.WriteHeader(http.StatusUnauthorized)
	} else {
		jwtToken := authArr[1]
		user, errGetUser := DB.JwtGetUser(jwtToken, session)
		if errGetUser != nil {
			rw.WriteHeader(http.StatusUnauthorized)
		}
		isValid := ValidateToken(jwtToken, user.UserName)

		if !isValid {
			rw.WriteHeader(http.StatusUnauthorized)
		} else {
			Action := req.Method
			Route := req.URL.Path

			role := user.Role
			if role == "" {
				role = "user"
			}
			fmt.Println(role, " ", Action, " ", Route)
			if CheckAccess(role, Route, Action) {
				next(rw, req)
			} else {
				rw.WriteHeader(http.StatusUnauthorized)

			}

		}

	}
}
func CheckAccess(sub, obj, act string) (Granted bool) {
	session, errConDB := DB.ConnectDB()
	if errConDB != nil {
		log.SystemErrorHappened(errConDB)
		return false
	}
	Granted = false
	permissionCasbin := DB.GetPermision(session)
	sa := scas.NewAdapter(permissionCasbin.Policy)
	e := casbin.NewEnforcer(casbin.NewModel(permissionCasbin.Model), sa)
	if e.Enforce(sub, obj, act) == true {
		return true
	} else {
		return false
	}

	return
}
