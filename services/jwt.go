package services

import (
	"encoding/json"
	"errors"
	"github.com/ali-zohrevand/ashyanet-api/OutputAPI"
	"github.com/ali-zohrevand/ashyanet-api/core/DB"
	"github.com/ali-zohrevand/ashyanet-api/services/log"
	"github.com/ali-zohrevand/ashyanet-api/settings/Words"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

type JWTData struct {
	// Standard claims are the standard jwt claims from the IETF standard
	// https://tools.ietf.org/html/rfc7519
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	//todo: add token and its ip to a
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 1000).Unix(),
		"iat": time.Now().Unix(),
		"sub": username,
	}
	tokenString, err := token.SignedString([]byte(Words.TokenKey))
	//		tokenString, err := token.SignedString([]byte(SECRET))
	if err != nil {
		log.ErrorHappened(err)
		return "", err
	}
	return tokenString, nil
}
func ValidateToken(tokenString string, username string) (IsValid bool) {
	var jwtData JWTData
	token, err := jwt.ParseWithClaims(tokenString, &jwtData, func(token *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != token.Method {
			return nil, errors.New("Invalid signing algorithm")
		}

		return []byte(Words.TokenKey), nil
	})

	//todo check ip and it tokenString
	//todo if ip of tokenString is not equal log security alert and unauthorized
	subject := jwtData.Subject
	if subject != username || err != nil || !token.Valid {
		return false

	}
	return true
}
func IsJwtValid(token string) (int, []byte) {
	session, errConnectDB := DB.ConnectDB()
	if errConnectDB != nil {
		log.SystemErrorHappened(errConnectDB)
		var message OutputAPI.TokenValid
		message.Valid = false
		outJason, _ := json.Marshal(message)
		return http.StatusNotFound, outJason
	}
	User, err := DB.JwtGetUser(token, session)
	if err != nil {
		var message OutputAPI.TokenValid
		message.Valid = false
		outJason, _ := json.Marshal(message)
		return http.StatusNotFound, outJason

	}
	if ValidateToken(token, User.UserName) {
		var message OutputAPI.TokenValid
		message.Valid = true
		outJason, _ := json.Marshal(message)
		return http.StatusOK, outJason

	}
	var message OutputAPI.TokenValid
	message.Valid = false
	outJason, _ := json.Marshal(message)
	return http.StatusNotFound, outJason

}

/*
authToken := r.Header.Get("Authorization")
	authArr := strings.Split(authToken, " ")

	if len(authArr) != 2 {
		log.Println("Authentication header is invalid: " + authToken)
		http.Error(w, "Request failed!", http.StatusUnauthorized)
	}

	jwtToken := authArr[1]

	claims, err := jwt.ParseWithClaims(jwtToken, &JWTData{}, func(token *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != token.Method {
			return nil, errors.New("Invalid signing algorithm")
		}
		return []byte(SECRET), nil
	})

	if err != nil {
		log.Println(err)
		http.Error(w, "Request failed!", http.StatusUnauthorized)
	}

	data := claims.Claims.(*JWTData)

	userID := data.CustomClaims["userid"]

	// fetch some data based on the userID and then send that data back to the user in JSON format
	jsonData, err := getAccountData(userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Request failed!", http.StatusUnauthorized)
	}

	w.Write(jsonData)
*/
/*

	token, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		} else {
			return authBackend.PublicKey, nil
		}
	})

	if err == nil && token.Valid && !authBackend.IsInBlacklist(req.Header.Get("Authorization")) {
		next(rw, req)
	} else {
		rw.WriteHeader(http.StatusUnauthorized)
	}
*/
