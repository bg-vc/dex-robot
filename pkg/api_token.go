package pkg

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

var (
	ErrMissingHeader = errors.New("the length of the authorization header is zero")
)

type Context struct {
	ID       uint64
	Username string
}

func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

func Parse(tokenString string, secret string) (*Context, error) {
	ctx := &Context{}

	token, err := jwt.Parse(tokenString, secretFunc(secret))

	if err != nil {
		return ctx, err

	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.ID = uint64(claims["id"].(float64))
		ctx.Username = claims["username"].(string)
		return ctx, nil

	} else {
		return ctx, err
	}
}

func ParseRequest(c *gin.Context, accountType int) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	secret := ""
	if accountType == 1 {
		secret = viper.GetString("jwt_secret_admin")
	} else {
		secret = viper.GetString("jwt_secret_user")
	}

	if len(header) == 0 {
		return &Context{}, ErrMissingHeader
	}

	var t string

	fmt.Sscanf(header, "Bearer %s", &t)
	return Parse(t, secret)
}

func Sign(c Context, accountType int) (tokenString string, err error) {
	secret := ""
	if accountType == 1 {
		secret = viper.GetString("jwt_secret_admin")
	} else {
		secret = viper.GetString("jwt_secret_user")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.ID,
		"username": c.Username,
		"exp":      time.Now().Unix() + 60*30,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})
	tokenString, err = token.SignedString([]byte(secret))
	return
}
