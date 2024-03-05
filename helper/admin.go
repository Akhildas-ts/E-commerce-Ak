package helper

import (
	"ak/config"
	"ak/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type authCustomClaimsAdmin struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	jwt.StandardClaims
}

func GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error) {

	cfg, _ := config.LoadConfig()

	claims := &authCustomClaimsAdmin{
		Firstname: admin.Firstname,
		Lastname:  admin.Lastname,
		Email:     admin.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenstring, err := token.SignedString([]byte(cfg.KEY_FOR_ADMIN))

	if err != nil {
		fmt.Println("error from token generate check its ", err)
		return "", err
	}

	return tokenstring, nil
}

func ValidateToken(tokenstring string) (*authCustomClaimsAdmin, error) {

	cfg, _ := config.LoadConfig()

	token, err := jwt.ParseWithClaims(tokenstring, &authCustomClaimsAdmin{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])

		}

		return []byte(cfg.KEY_FOR_ADMIN), nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*authCustomClaimsAdmin); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// func GetTimeFromPeriod(timePeriod string) (time.Time,time.Time) {

// 	endDate := time.Now()

// 	if timePeriod == "week" {

// 		startDate := endDate.AddDate(0,0,-6)
// 		return startDate,endDate
// 	}

// 	if timePeriod == "month" {

// 		startDate := endDate.AddDate(0,-1,0)
// 		return startDate,endDate
// 	}

// 	if timePeriod == "year" {
// 		startDate := endDate.AddDate(-1,0,0)
// 		return startDate,endDate
// 	}

// 	return endDate.AddDate(0,0,-6),endDate
// }
