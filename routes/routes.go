package routes

import (
	"quote-generator-backend/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, qc *controllers.QuoteController, uc *controllers.UserController) {
    r.POST("/quotes", cors.Default(), qc.AddQuote)
    r.GET("/quotes/:category", cors.Default(), qc.GetQuotesByCategory)
    r.GET("/quotes/random", cors.Default(), qc.GetRandomQuotes)
	r.POST("/login", cors.Default(), uc.LoginOrCreateUser)
    r.POST("/users/:user_id/favorites", cors.Default(), uc.AddFavorite)
    r.GET("/users/:user_id/favorites", cors.Default(), uc.GetFavorites)
}