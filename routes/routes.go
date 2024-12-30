package routes

import (
	"net/http"
	"quote-generator-backend/controllers"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") 
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func SetupRoutes(r *gin.Engine, qc *controllers.QuoteController, uc *controllers.UserController) {
    r.Use(CORSMiddleware())

    r.POST("/quotes", qc.AddQuote)
    r.GET("/quotes/:category", qc.GetQuotesByCategory)
    r.GET("/quotes/random", qc.GetRandomQuotes)
	r.POST("/login", uc.LoginOrCreateUser)
    r.POST("/users/:user_id/favorites", uc.AddFavorite)
    r.GET("/users/:user_id/favorites", uc.GetFavorites)
}
