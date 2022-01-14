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
	Administrator bool   `json:"administrator,omitempty"`
	Name          string `json:"name,omitempty"`
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
		if decoded, err = jwt.ParseWithClaims(token[7:], sess, func(t *jwt.Token) (interface{}, error) {
			return []byte(configuration.Get().SecretKey), nil
		}); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You do not have permission"})
			return
		}

		claims := decoded.Claims.(*Session)
		if claims.Issuer != "isaqueveras.auth" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You do not have permission"})
			return
		}

		c.Set("session", *sess)
		c.Next()
	}
}

// GetGinSession is a short hand to retrieve a session from a gin context
func GetGinSession(c *gin.Context) (sess *Session, err error) {
	var (
		value interface{}
		ok    bool
	)

	if value, ok = c.Get("session"); !ok {
		return nil, errors.New("invalid session")
	}

	if value, ok := value.(Session); ok {
		sess = &value
		return
	}

	return nil, errors.New("invalid session")
}
