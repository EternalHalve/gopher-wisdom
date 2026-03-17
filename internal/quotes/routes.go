package quotes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := NewQuoteHandler(db)

	v1 := r.Group("/api/v1")
	{
		quotesGroup := v1.Group("/quotes")
		{
			quotesGroup.GET("", h.GetQuotes)
			quotesGroup.GET("/:id", h.GetQuotesByID)
			quotesGroup.POST("", h.PostQuotes)
		}
	}
}
