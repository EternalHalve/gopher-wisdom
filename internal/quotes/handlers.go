package quotes

import (
	"net/http"

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

func PostQuotes(c *gin.Context) {
	var NewQuotes Quote

	if err := c.BindJSON(&NewQuotes); err != nil {
		return
	}

	QuotesList = append(QuotesList, NewQuotes)
	c.IndentedJSON(http.StatusCreated, NewQuotes)
}
