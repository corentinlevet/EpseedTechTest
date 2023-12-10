package server

import (
	"epseed/internal/controllers"
	"epseed/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func initAuthGroup(router *gin.Engine) {
	authService := &services.AuthService{}
	userService := &services.UserService{}

	authController := &controllers.AuthController{
		AuthService: authService,
		UserService: userService,
	}

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/signup", authController.Signup)
	}
}

func initNotesGroup(router *gin.Engine) {
	noteService := &services.NoteService{}
	noteController := &controllers.NoteController{NoteService: noteService}

	notesGroup := router.Group("/notes")
	{
		notesGroup.GET("/", noteController.GetNotes)
		notesGroup.POST("/", noteController.CreateNote)
		notesGroup.PUT("/", noteController.UpdateNote)
		notesGroup.DELETE("/", noteController.DeleteNote)
	}
}

func initUsersGroup(router *gin.Engine) {
	userService := &services.UserService{}
	userController := &controllers.UserController{UserService: userService}

	usersGroup := router.Group("/users")
	{
		usersGroup.GET("/", userController.GetUsers)
	}
}

func InitRoutes() {
	router := gin.Default()

	cors := cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8081"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type"},
	})

	router.Use(cors)

	initAuthGroup(router)
	initNotesGroup(router)
	initUsersGroup(router)

	router.Run(":8080")
}
