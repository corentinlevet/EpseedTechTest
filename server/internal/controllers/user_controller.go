package controllers

import (
	"epseed/internal/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

// GET /users
func (c *UserController) GetUsers(ctx *gin.Context) {
	if ctx.Request.Method != "GET" {
		ctx.JSON(405, gin.H{"message": "Method not allowed"})
		return
	}

	users, err := c.UserService.GetUsers()

	if err != nil {
		ctx.JSON(404, gin.H{"message": "Users not found"})
		return
	}

	ctx.JSON(200, users)
}
