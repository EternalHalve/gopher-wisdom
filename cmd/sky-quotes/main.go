package main

import (
    "github.com/gin-gonic/gin"
    "github.com/EternalHalve/sky-quotes/internal/quotes"
)

func main() {
    router := gin.Default()
    
    router.GET("/quotes", quotes.GetQuotes)

    router.Run("localhost:8080")
}