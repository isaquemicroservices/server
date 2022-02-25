package middleware

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	config "github.com/isaqueveras/servers-microservices-backend/configuration"
)

// AuthorizationGin is middleware for gin
func AuthorizationGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			err   error
			sess  *config.Session
		)

		if token = c.GetHeader("Authorization"); token == "" || len(token) < 10 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You do not have permission"})
			return
		}

		if sess, err = decodeJWT(token[7:], config.Get().SecretKey); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			return
		}

		c.Set("session", *sess)
		c.Next()
	}
}

// GetSession get session on context of request
func GetSession(c *gin.Context) (sess *config.Session, err error) {
	var (
		v  interface{}
		ok bool
	)

	if v, ok = c.Get("session"); !ok {
		return nil, errors.New("Invalid session")
	}

	if val, ok2 := v.(config.Session); ok2 {
		sess = &val
		return
	}

	return nil, errors.New("Invalid session")
}

func decodeJWT(token string, secret string) (sess *config.Session, err error) {
	var decoded *jwt.Token
	sess = new(config.Session)

	if decoded, err = jwt.ParseWithClaims(token, sess, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method: " + token.Header["alg"].(string))
		}
		return []byte(secret), nil
	}); err != nil {
		return nil, err
	}

	if decoded == nil {
		return nil, errors.New("Was not possible to decode the token")
	}

	if claims, ok := decoded.Claims.(*config.Session); ok && decoded.Valid {
		if claims.Issuer != "isaqueveras.auth" {
			return nil, errors.New("Token issuer is not valid")
		}
		return claims, nil
	}

	return nil, errors.New("Token content is not valid")
}

// AdminOnly
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		if sess, err := GetSession(c); err != nil || !*sess.Permission.IsAdmin {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
