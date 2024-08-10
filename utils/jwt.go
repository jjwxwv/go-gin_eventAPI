package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//secretkey
const secret = "supernova"

func GenerateToken(email string, userId int64) (string, error) {
	//generate new token
	//need to specify signing method and claims value
	//dont input password in claims value
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 2).Unix(), //assign expired to token
	})
	//key parameter in signedstring func can be any type but actually should convert to byte
	return token.SignedString([]byte(secret))
}

func VerifyToken(token string) (int64, error) {
	//check the token is valid by using parse method
	//parse receive token and the func that will return secretkey in the end
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		//checking that verifying token was signed with correct signing method by using type checking syntax
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("could not parse token")
	}
	isValid := parsedToken.Valid
	if !isValid {
		return 0, errors.New("invalid token")
	}
	//get the data from token and checking token type is correct
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}
	// email := claims["email"].(string)
	//for userId we store in type int64 but it turns out that the way this claim is stored and then retrieved later, it isn't stored as int64 but instead float64
	userId := int64(claims["userId"].(float64))
	return userId, nil
}