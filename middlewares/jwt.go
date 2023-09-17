package middlewares

import (
	"myproject/models/user/database"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jwtCustomClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
	jwt.RegisteredClaims
}

func GenerateToken(user database.User) string {
	claims := &jwtCustomClaims{
		user.ID,
		user.Username,
		user.Admin,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenized, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	return tokenized
}
