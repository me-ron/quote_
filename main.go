package handler

import (
	"log"
	"net/http"
	"quote-generator-backend/config"
	"quote-generator-backend/controllers"
	"quote-generator-backend/repositories"
	"quote-generator-backend/routes"
	"quote-generator-backend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Load environment variables from .env file if available
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using environment variables.")
	}

	// Initialize the database connection
	err = config.ConnectDB()
	if err != nil {
		log.Printf("Error connecting to DB: %s", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	db := config.DB
	quoteRepo := &repositories.QuoteRepository{Collection: db.Collection("quotes")}
	userRepo := &repositories.UserRepository{Collection: db.Collection("users")}

	quoteService := &services.QuoteService{Repo: quoteRepo}
	userService := &services.UserService{Repo: userRepo}

	quoteController := &controllers.QuoteController{Service: quoteService}
	userController := &controllers.UserController{Service: userService}

	// Setup the Gin router and handle the request
	router := gin.Default()
	router.Use(cors.Default())
	routes.SetupRoutes(router, quoteController, userController)

	// Use the Gin engine to handle the request and return the response
	router.ServeHTTP(w, r)
}
