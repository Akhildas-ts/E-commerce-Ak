package helper

import (
	"ak/config"
	"ak/models"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthCustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func PasswordHashing(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return "", errors.New("hash Server issue")
	}

	hash := string(hashedPassword)

	return hash, nil

}

func GenerateTokenUsers(userID int, userEmail string, expirationTime time.Time) (string, error) {
	cfg, _ := config.LoadConfig()

	claims := &AuthCustomClaims{
		Id:    userID,
		Email: userEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.KEY))

	// tokenString, err:=jwt.NewWithClaims(jwt.SigningMethodHS256, &AuthCustomClaims{Id:    userID, Email: userEmail,StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix(),IssuedAt:  time.Now().Unix(),},}).SignedString([]byte(cfg.KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateAccessToken(user models.SignupDetailResponse) (string, error) {

	expirationTime := time.Now().Add(15 * time.Minute)
	tokenString, err := GenerateTokenUsers(user.ID, user.Email, expirationTime)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func GenerateRefreshToken(user models.SignupDetailResponse) (string, error) {

	expirationTime := time.Now().Add(24 * 90 * time.Hour)
	tokeString, err := GenerateTokenUsers(user.ID, user.Email, expirationTime)
	if err != nil {
		return "", err
	}
	return tokeString, nil

}