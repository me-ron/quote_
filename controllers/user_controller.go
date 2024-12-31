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

func (uc *UserController) UpdateUserPreferences(c *gin.Context) {
    userIDStr := c.Param("user_id")
    userID, err := primitive.ObjectIDFromHex(userIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
        return
    }

    var preferences struct {
        Limit      int      `json:"limit"`
        Categories []string `json:"categories"`
    }

    if err := c.ShouldBindJSON(&preferences); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    err = uc.Service.UpdateUserPreferences(userID, preferences.Limit, preferences.Categories)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update preferences"})
        return
    }

    c.Redirect(http.StatusTemporaryRedirect, "/quotes/random?user_id="+userID.Hex())

}


