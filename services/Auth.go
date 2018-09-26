package services

import (
	"fmt"
	"github.com/casbin/casbin"
	scas "github.com/qiangmzsx/string-adapter"
	"gitlab.com/hooshyar/ChiChiNi-API/core/DB"
	"gitlab.com/hooshyar/ChiChiNi-API/services/log"
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
		isValid := ValidateToken(jwtToken)
		if !isValid {
			rw.WriteHeader(http.StatusUnauthorized)
		} else {
			Action := req.Method
			Route := req.URL.Path
			user, errGetUser := DB.GetUserOfSession(jwtToken, session)
			if errGetUser != nil {
				rw.WriteHeader(http.StatusUnauthorized)
			}
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
