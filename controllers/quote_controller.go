package controllers

import (
	"net/http"
	"quote-generator-backend/models"
	"quote-generator-backend/services"
	"strconv"
	"strings"

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
    // Get limit from query parameters, default to 5 if not provided
    limit := 5
    if queryLimit := c.DefaultQuery("limit", "5"); queryLimit != "" {
        parsedLimit, err := strconv.Atoi(queryLimit)
        if err == nil {
            limit = parsedLimit
        }
    }

    // Get categories from query parameters
    categories := c.DefaultQuery("categories", "")
    var categoryList []string
    if categories != "" {
        categoryList = strings.Split(categories, ",")
    }

    // Fetch random quotes with the limit and categories
    quotes, err := qc.Service.GetRandomQuotes(limit, categoryList...)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, quotes)
}

