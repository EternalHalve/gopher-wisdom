package quotes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var QuotesList = []Quote{
	{
		ID:        1,
		Content:   "That wherever you are in the world, I'll search for you",
		Anime:     "Your Name",
		Character: "Taki",
	},
	{
		ID:        2,
		Content:   "A man must raise more grain than he has trampled in his life. His hands must build more homes than they have burned.",
		Anime:     "Vinland Saga",
		Character: "Thorfinn",
	},
}

func GetQuotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, QuotesList)
}

func GetQuotesByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}

	for _, a := range QuotesList {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Quotes not found"})
}

func PostQuotes(c *gin.Context) {
	var NewQuotes Quote

	if err := c.BindJSON(&NewQuotes); err != nil {
		return
	}

	QuotesList = append(QuotesList, NewQuotes)
	c.IndentedJSON(http.StatusCreated, NewQuotes)
}
