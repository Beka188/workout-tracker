package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey []byte

func init() {
	key, err := loadConfig()
	if err != nil {
		fmt.Println("myerr", err)
		return
	}
	secretKey = key
}

// CustomClaims extends jwt.StandardClaims
type CustomClaims struct {
	Email  string `json:"email"`
	UserID int    `json:"user_id"`
	jwt.StandardClaims
}

// CreateToken generates a new JWT token
func CreateToken(email string, userID int) (string, error) {
	claims := CustomClaims{
		Email:  email,
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// VerifyToken validates the JWT token
func VerifyToken(tokenString string, expectedEmail string, expectedUserID int) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		if time.Now().Unix() > claims.ExpiresAt {
			return false, errors.New("token has expired")
		}

		if claims.Email != expectedEmail {
			return false, errors.New("email in token does not match expected email")
		}
		if claims.UserID != expectedUserID {
			return false, errors.New("user ID in token does not match expected user ID")
		}
		return true, nil
	}

	return false, errors.New("invalid token")
}
