package controllers

import (
	"epseed/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *services.AuthService
	UserService *services.UserService
}

// POST /auth/login
func (c *AuthController) Login(ctx *gin.Context) {
	if ctx.Request.Method != "POST" {
		ctx.JSON(405, gin.H{"message": "Method not allowed"})
		return
	}

	var json struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(400, gin.H{"message": "Bad request"})
		return
	}

	user, err := c.AuthService.Login(json.Username, json.Password)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Error while logging in"})
		return
	}

	if user == nil {
		ctx.JSON(404, gin.H{"message": "User not found"})
		return
	}

	ctx.JSON(200, gin.H{"message": "User logged in", "token": "token", "user_id": user.ID})
}

// POST /auth/signup
func (c *AuthController) Signup(ctx *gin.Context) {
	if ctx.Request.Method != "POST" {
		ctx.JSON(405, gin.H{"message": "Method not allowed"})
		return
	}

	var json struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(400, gin.H{"message": "Bad request"})
		return
	}

	err := c.AuthService.Signup(json.Username, json.Email, json.Password)
	if err != nil {
		if err.Error() == "User already exists" {
			ctx.JSON(409, gin.H{"message": "User already exists"})
			return
		}

		ctx.JSON(500, gin.H{"message": "Error while signing up"})
		return
	}

	user, _ := c.UserService.GetUserByEmail(json.Email)
	if user == nil {
		ctx.JSON(404, gin.H{"message": "User not found"})
		return
	}

	ctx.JSON(200, gin.H{"message": "User signed up", "token": "token", "user_id": user.ID})
}
