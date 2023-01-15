package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type Claims struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

func ParseJWT(id uint, name string) (string, error) {
	now := time.Now()
	expireTime := now.Add(3 * time.Hour)

	claims := Claims{
		id,
		name,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Its me",
		},
	}
	jwtSecret := []byte("OWIJFOIEWJFOIEWOIFJEOWJFIEWFIefjwoiehrjbkwdfldn ,nblj;'pi21[urpilknvsdj")
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtSecret)
}

func DumpJWT(token string) (*Claims, error) {
	jwtSecret := []byte("OWIJFOIEWJFOIEWOIFJEOWJFIEWFIefjwoiehrjbkwdfldn ,nblj;'pi21[urpilknvsdj")
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims := tokenClaims.Claims.(*Claims); tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(200, Message{
				Code: -1,
				Msg:  "Token 无效",
			})
		}
		claims, err := DumpJWT(token)
		if err != nil {
			c.JSON(200, Message{
				Code: -1,
				Msg:  err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("id", claims.ID)
	}
}
