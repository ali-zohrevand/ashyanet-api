package services

import (
	"errors"
	"github.com/ali-zohrevand/ashyanet-api/services/log"
	"github.com/ali-zohrevand/ashyanet-api/settings/ConstKey"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTData struct {
	// Standard claims are the standard jwt claims from the IETF standard
	// https://tools.ietf.org/html/rfc7519
	jwt.StandardClaims
}

func GenerateToken(userUUID string) (string, error) {
	//todo: add token and its ip to a
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 10000).Unix(),
		"iat": time.Now().Unix(),
		"sub": userUUID,
	}
	tokenString, err := token.SignedString([]byte(ConstKey.TokenKey))
	//		tokenString, err := token.SignedString([]byte(SECRET))
	if err != nil {
		log.ErrorHappened(err)
		return "", err
	}
	return tokenString, nil
}
func ValidateToken(token string) (IsValid bool) {
	_, err := jwt.ParseWithClaims(token, &JWTData{}, func(token *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != token.Method {
			return nil, errors.New("Invalid signing algorithm")
		}
		return []byte(ConstKey.TokenKey), nil
	})

	//todo check ip and it token
	//todo if ip of token is not equale log security alert and unathorized
	if err != nil {
		return false

	}
	return true
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
