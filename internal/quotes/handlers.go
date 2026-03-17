package quotes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QuoteHandler struct {
	DB *gorm.DB
}

func NewQuoteHandler(db *gorm.DB) *QuoteHandler {
	return &QuoteHandler{DB: db}
}

func (handler *QuoteHandler) GetQuotes(c *gin.Context) {
	var quotes []Quote
	handler.DB.Find(&quotes)
	c.IndentedJSON(http.StatusOK, quotes)
}

func (handler *QuoteHandler) GetQuotesByID(c *gin.Context) {
	id := c.Param("id")
	var quote Quote

	if err := handler.DB.First(&quote, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Quote not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, quote)
}

func (handler *QuoteHandler) PostQuotes(c *gin.Context) {
	var newQuote Quote

	if err := c.BindJSON(&newQuote); err != nil {
		return
	}

	handler.DB.Create(&newQuote)
	c.IndentedJSON(http.StatusCreated, newQuote)
}
