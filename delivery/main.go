package main

import (
	"log"
	"os"
	"quote-generator-backend/config"
	"quote-generator-backend/controllers"
	"quote-generator-backend/repositories"
	"quote-generator-backend/routes"
	"quote-generator-backend/services"

	"github.com/gin-gonic/gin"
)

func main() {
    err := config.ConnectDB()
	if (err != nil){
		log.Println(err.Error())
	}
	db := config.DB
    quoteRepo := &repositories.QuoteRepository{Collection: db.Collection("quotes")}
    userRepo := &repositories.UserRepository{Collection: db.Collection("users")}

    quoteService := &services.QuoteService{Repo: quoteRepo}
    userService := &services.UserService{Repo: userRepo}

    quoteController := &controllers.QuoteController{Service: quoteService}
    userController := &controllers.UserController{Service: userService}

    r := gin.Default()
    routes.SetupRoutes(r, quoteController, userController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" 
	}

    r.Run("0.0.0.0:"+port)
}
