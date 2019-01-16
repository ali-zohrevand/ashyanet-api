package services

import (
	"encoding/json"
	"fmt"
	"github.com/ali-zohrevand/ashyanet-api/OutputAPI"
	"github.com/ali-zohrevand/ashyanet-api/core/DB"
	"github.com/ali-zohrevand/ashyanet-api/models"
	"github.com/ali-zohrevand/ashyanet-api/services/log"
	"github.com/ali-zohrevand/ashyanet-api/services/validation"
	"github.com/casbin/casbin"
	scas "github.com/qiangmzsx/string-adapter"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strings"
	"time"
)

func Login(requestUser *models.User, request *http.Request) (int, []byte) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		return http.StatusInternalServerError, []byte("")
	}
	defer session.Close()
	out, errValidation, IsValid := validation.UserLoginValidation(*requestUser)
	if errValidation != nil || !IsValid {
		return http.StatusBadRequest, out
	}
	UserDatastore := DB.UserDataStore{}
	session, errConDB := DB.ConnectDB()
	if errConDB != nil {
		log.SystemErrorHappened(errConDB)
		return http.StatusUnauthorized, []byte("")
	}
	defer session.Close()
	if UserDatastore.CheckUserPassCorrect(*requestUser, session) {
		token, err := GenerateToken(requestUser.UserName)
		if err != nil {
			log.SystemErrorHappened(err)
			return http.StatusUnauthorized, []byte("")
		} else {
			//................................. Main Action is Hear .................................
			response, err := json.Marshal(OutputAPI.TokenAuthentication{token})
			if err != nil {
				log.SystemErrorHappened(err)
				return http.StatusUnauthorized, []byte("")
			}
			sessionObj := models.JwtSession{
				Id:            bson.NewObjectId(),
				JwtToken:      token,
				OwnerUsername: requestUser.UserName,
				TimeCreated:   time.Now(),
				Ip:            GetIpOfRequest(request),
			}
			var jwtSessionDataStore = DB.JwtSessionDataStore{}
			errAddToSessionDB := jwtSessionDataStore.CreateJwtSession(sessionObj, session)
			if errAddToSessionDB != nil {
				log.SystemErrorHappened(errConnectDB)
				return http.StatusInternalServerError, []byte("")
			}
			return http.StatusOK, response
			//..................................................................
		}
	} else {
		//todo: security log Happened beacuse user and password not match
	}
	return http.StatusUnauthorized, []byte("")
}
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
			role := DB.GetUserOfSession(jwtToken, session)
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
