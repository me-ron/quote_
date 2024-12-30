package main

import (
	"log"
	"net/http"
	"quote-generator-backend/config"
	"quote-generator-backend/controllers"
	"quote-generator-backend/repositories"
	"quote-generator-backend/routes"
	"quote-generator-backend/services"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
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

    app = gin.New()
	r := app.Group("/api")

    routes.SetupRoutes(r, quoteController, userController)
	
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}