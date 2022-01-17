package middleware

import (
	"errors"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/servers-microservices-backend/configuration"
)

type Session struct {
	Administrator *bool   `json:"administrator,omitempty"`
	Name          *string `json:"name,omitempty"`
	jwt.StandardClaims
}

// ValidateJWT access token validation
func ValidateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token-auth")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Erro on cookie"})
			return
		}

		token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(configuration.Get().SecretKey), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You do not have permission"})
			return
		}

		claims := token.Claims.(*jwt.StandardClaims)
		log.Println(claims.Issuer)

		c.Next()
	}
}

// AuthorizationGin is middleware for gin
func AuthorizationGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token   string
			err     error
			decoded *jwt.Token
		)

		if token = c.GetHeader("Authorization"); token == "" || len(token) < 10 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You do not have permission"})
			return
		}

		sess := new(Session)
		if decoded, err = jwt.ParseWithClaims(token, sess, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method: " + token.Header["alg"].(string))
			}
			return []byte(configuration.Get().SecretKey), nil
		}); err != nil {
			return
		}

		if claims, ok := decoded.Claims.(*Session); ok && decoded.Valid {
			if claims.Issuer != "isaqueveras.auth" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token issuer is not valid"})
				return
			}
			return
		}

		c.Next()
	}
}

// GetSession get session on context of request
func GetSession(c *gin.Context) (sess *Session, err error) {
	var (
		v  interface{}
		ok bool
	)

	if v, ok = c.Get("session"); !ok {
		return nil, errors.New("Invalid session")
	}

	if val, ok2 := v.(Session); ok2 {
		sess = &val
		return
	}

	return nil, errors.New("Invalid session")
}
