package controllers

import (
    "net/http"
    "quote-generator-backend/services"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
    Service *services.UserService
}

func (h *UserController) LoginOrCreateUser(c *gin.Context) {
	var request struct {
		Username string `json:"username" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userID, err := h.Service.LoginOrCreate(request.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to process request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userID.Hex()})
}

func (uc *UserController) AddFavorite(c *gin.Context) {
    var request struct {
        QuoteID string `json:"quote_id"`
    }

    if err := c.BindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID, _ := primitive.ObjectIDFromHex(c.Param("user_id"))
    quoteID, _ := primitive.ObjectIDFromHex(request.QuoteID)

    err := uc.Service.AddFavorite(userID, quoteID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Favorite added"})
}

func (uc *UserController) GetFavorites(c *gin.Context) {
    userID, _ := primitive.ObjectIDFromHex(c.Param("user_id"))
    quotes, err := uc.Service.GetFavorites(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, quotes)
}

