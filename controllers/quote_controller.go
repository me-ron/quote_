package controllers

import (
	"net/http"
	"quote-generator-backend/models"
	"quote-generator-backend/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuoteController struct {
    Service *services.QuoteService
}

// AddQuote adds a new quote to the database.
func (h *QuoteController) AddQuote(c *gin.Context) {
	var quote models.Quote
	if err := c.ShouldBindJSON(&quote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Set a new ObjectID for the quote if it's not set
	if quote.ID.IsZero() {
		quote.ID = primitive.NewObjectID()
	}

	err := h.Service.AddQuote(quote)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to add quote"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Quote added successfully", "quote_id": quote.ID.Hex()})
}

func (qc *QuoteController) GetQuotesByCategory(c *gin.Context) {
    category := c.Param("category")

    quotes, err := qc.Service.GetQuotesByCategory(category)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, quotes)
}


func (qc *QuoteController) GetRandomQuotes(c *gin.Context) {
    // Get the user_id from query parameters
    userIDStr := c.DefaultQuery("user_id", "")
    var userID primitive.ObjectID
    if userIDStr == "" {
        userID = primitive.NilObjectID
    }else{
        var err error
        userID, err = primitive.ObjectIDFromHex(userIDStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
            return
        }
    }

    // Fetch random quotes with the limit and categories based on the user's preferences
    quotes, err := qc.Service.GetRandomQuotes(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, quotes)
}


// GetCategories fetches all distinct categories, optionally filtered by userID.
func (qc *QuoteController) GetCategories(c *gin.Context) {
    userIDParam := c.DefaultQuery("user_id", "")
    var userID primitive.ObjectID
    if userIDParam == "" {
        userID = primitive.NilObjectID
        
    }else{
        objID, err := primitive.ObjectIDFromHex(userIDParam)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
            return
        }
        userID = objID
    }

    categories, err := qc.Service.GetAllCategories(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, categories)
}