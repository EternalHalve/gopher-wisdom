package quotes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("wisdom.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	DB.AutoMigrate(&Quote{})
}

func GetQuotes(c *gin.Context) {
	var quotes []Quote
	DB.Find(&quotes)
	c.IndentedJSON(http.StatusOK, quotes)
}

func GetQuotesByID(c *gin.Context) {
	id := c.Param("id")
	var quote Quote

	if err := DB.First(&quote, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Quote not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, quote)
}

func PostQuotes(c *gin.Context) {
	var newQuote Quote

	if err := c.BindJSON(&newQuote); err != nil {
		return
	}

	DB.Create(&newQuote) // INSERT INTO quotes ...
	c.IndentedJSON(http.StatusCreated, newQuote)
}
