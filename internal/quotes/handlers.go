package quotes

import (
	"log"
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
	format := c.Query("format")
	var quote Quote

	if err := handler.DB.First(&quote, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "The requested fragment of wisdom remains undiscovered in these tunnels.",
			"hint":  "Perhaps the archives are incomplete, or your path lies elsewhere.",
		})
		return
	}

	response := gin.H{
		"content":   quote.Content,
		"character": quote.Character,
		"anime":     quote.Anime,
	}

	if format == "alien" {
		response["content"] = Alienify(quote.Content)
		response["dialect"] = "Zorgon-7"
	}

	c.IndentedJSON(http.StatusOK, response)
}

func (handler *QuoteHandler) PostQuotes(c *gin.Context) {
	var newQuote Quote

	if err := c.ShouldBindJSON(&newQuote); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handler.DB.Create(&newQuote)
	c.IndentedJSON(http.StatusCreated, newQuote)
}

func SeedData(db *gorm.DB) {
	var count int64

	if err := db.Model(&Quote{}).Count(&count).Error; err != nil {
		return
	}

	if count == 0 {
		quotes := []Quote{
			{Content: "I have no enemies... No one has any enemies", Anime: "Vinland Saga", Character: "Thorfinn"},
			{Content: "No matter how hard or impossible it is, never lose sight of your goal", Anime: "One Piece", Character: "Luffy"},
			{Content: "If I don't wield the sword, I can't protect you", Anime: "Bleach", Character: "Ichigo"},
		}

		if err := db.Create(&quotes).Error; err != nil {
			log.Printf("Could not seed: %v", err)
		}
	}
}
