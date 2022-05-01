package auth

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/servers-microservices-backend/application/crm/auth"
	"github.com/isaqueveras/servers-microservices-backend/configuration"
	"github.com/isaqueveras/servers-microservices-backend/oops"
)

func create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Copy().Request.Context(), configuration.Get().ContextWithTimeout)
	defer cancel()

	var (
		req auth.User
		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		oops.Handling(err, c)
		return
	}

	if err = auth.Create(ctx, &req); err != nil {
		oops.Handling(err, c)
		return
	}

	c.JSON(201, nil)
}

func login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Copy().Request.Context(), configuration.Get().ContextWithTimeout)
	defer cancel()

	var (
		req  auth.CredentialsReq
		data *auth.User
		err  error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		oops.Handling(err, c)
		return
	}

	if data, err = auth.Login(ctx, &req); err != nil {
		oops.Handling(err, c)
		return
	}

	c.Set("session", configuration.Session{Name: data.Name})
	c.JSON(200, data)
}
