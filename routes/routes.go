package routes

import (
	"quote-generator-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, qc *controllers.QuoteController, uc *controllers.UserController) {

    r.POST("/quotes", qc.AddQuote)
    r.GET("/quotes/:category", qc.GetQuotesByCategory)
    r.GET("/quotes/random", qc.GetRandomQuotes)
    r.GET("/quotes/categories", qc.GetCategories)
	r.POST("/login", uc.LoginOrCreateUser)
    r.POST("/users/:user_id/favorites", uc.AddFavorite)
    r.GET("/users/:user_id/favorites", uc.GetFavorites)
}
