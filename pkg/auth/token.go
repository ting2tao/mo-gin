package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"strconv"
)

const algorithm = "HS256"

type JWToken struct {
	timeout int64
	secret  string
}

func initJWT() *JWToken {
	timeOutStr := "2000"
	timeout, _ := strconv.Atoi(timeOutStr)
	token := new(JWToken)
	token.timeout = int64(timeout)
	token.secret = "2000"
	return token
}

func initStaffJWT() *JWToken {
	timeOutStr := "2000"
	timeout, _ := strconv.Atoi(timeOutStr)
	token := new(JWToken)
	token.timeout = int64(timeout)
	token.secret = "2000"
	return token
}

func (token *JWToken) JwtCreate(mc jwt.MapClaims) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod(algorithm), mc)
	// Sign and get the complete encoded token as a string
	tokenString, err := jwtToken.SignedString([]byte(token.secret))
	return tokenString, err
}

func (token *JWToken) JwtVerify(reqToken string) (claims jwt.MapClaims, err error) {
	jwtToken, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != algorithm {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		return []byte(token.secret), nil
	})
	if err != nil {
		return nil, err
	}

	var ok bool
	claims, ok = jwtToken.Claims.(jwt.MapClaims)
	if !(ok && jwtToken.Valid) {
		return nil, errors.New("validation failure")
	}

	//uid, ok := claims["uid"].(string)
	//if !ok {
	//	return "", errors.New("not available uid")
	//}
	return
}

func (token *JWToken) Param(reqToken, key string) (interface{}, error) {
	jwtToken, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != algorithm {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		return []byte(token.secret), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !(ok && jwtToken.Valid) {
		return "", errors.New("validation failure")
	}

	uid, ok := claims[key]
	if !ok {
		return "", errors.New("not available uid")
	}
	return uid, nil
}
