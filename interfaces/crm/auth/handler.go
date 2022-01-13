package auth

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/servers-microservices-backend/application/crm/auth"
	"github.com/isaqueveras/servers-microservices-backend/configuration"
)

func create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, configuration.Get().ContextWithTimeout)
	defer cancel()

	var (
		req auth.User
		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(500, gin.H{"message": "Error create user", "cause": err.Error()})
		return
	}

	if err = auth.Create(ctx, &req); err != nil {
		c.JSON(500, gin.H{"message": "Error create user", "cause": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "User created with success"})
}

func login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, configuration.Get().ContextWithTimeout)
	defer cancel()

	var (
		req  auth.CredentialsReq
		data *auth.User
		err  error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(500, gin.H{"message": "Error authenticating user", "cause": err.Error()})
		return
	}

	if data, err = auth.Login(ctx, &req); err != nil {
		c.JSON(500, gin.H{"message": "Error authenticating user", "cause": err.Error()})
		return
	}

	c.SetCookie("token", data.Token, 3600, "/", c.ClientIP(), false, true)

	c.JSON(200, data)
}

func logout(c *gin.Context) {
	c.SetCookie("token", "", -3600, "/", c.ClientIP(), false, true)
	c.JSON(200, gin.H{"message": "successfully logged out"})
}
