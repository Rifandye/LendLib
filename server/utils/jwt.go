package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateToken(firstName, lastName, email, role string, userId int64) (string, error) {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	secretKey := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"firstName": firstName,
		"lastName": lastName,
		"email": email,
		"role": role,
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, string, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	secretKey := os.Getenv("JWT_SECRET")

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpexted signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, "", errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, "", errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, "", errors.New("invalid token claims")
	}

	// firstName := claims["firstName"].(string)
	// lastName := claims["lastName"].(string)
	// email := claims["email"].(string)
	

	role, ok := claims["role"].(string)
	if !ok {
		return 0, "", errors.New("role in token is not a string")
	}
	
	userIdFloat64, ok := claims["userId"].(float64)
	if !ok {
		return 0, "", errors.New("userId in token is not a float64")
	}

	userId := int64(userIdFloat64)

	return userId, role, nil
}