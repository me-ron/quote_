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
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Initialize the database connection
	err := config.ConnectDB()
	if err != nil {
		log.Println(err.Error())
	}

	db := config.DB
	quoteRepo := &repositories.QuoteRepository{Collection: db.Collection("quotes")}
	userRepo := &repositories.UserRepository{Collection: db.Collection("users")}

	quoteService := &services.QuoteService{Repo: quoteRepo}
	userService := &services.UserService{Repo: userRepo}

	quoteController := &controllers.QuoteController{Service: quoteService}
	userController := &controllers.UserController{Service: userService}

	// Setup the Gin router
	app := gin.Default()
	routes.SetupRoutes(app, quoteController, userController)

	// Start serving the HTTP request with the Gin engine
	app.ServeHTTP(w, r)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

}
