package main

import (
	"github.com/EternalHalve/gopher-wisdom/internal/quotes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/quotes", quotes.GetQuotes)
	router.GET("/quotes/:id", quotes.GetQuotesByID)
	router.POST("/quotes", quotes.PostQuotes)

	router.Run("localhost:8080")
}
