package quotes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	db.AutoMigrate(&Quote{})
	return db
}

func TestQuoteHandlers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := setupTestDB()

	t.Run("PostQuotes should create a new quote", func(t *testing.T) {
		router := gin.New()
		RegisterRoutes(router, db)

		input := Quote{
			Content:   "Limit Break!",
			Anime:     "Dragon Ball",
			Character: "Goku",
		}
		body, _ := json.Marshal(input)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/quotes", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusCreated, resp.Code)

		var result Quote
		json.Unmarshal(resp.Body.Bytes(), &result)
		assert.Equal(t, "Goku", result.Character)
	})

	t.Run("GetQuotes should return all quotes", func(t *testing.T) {
		router := gin.New()
		RegisterRoutes(router, db)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/quotes", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)

		var results []Quote
		json.Unmarshal(resp.Body.Bytes(), &results)
		assert.GreaterOrEqual(t, len(results), 1)
		assert.Equal(t, "Limit Break!", results[0].Content)
	})

	t.Run("GetQuotesByID should return 404 for non-existent quote", func(t *testing.T) {
		router := gin.New()
		RegisterRoutes(router, db)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/quotes/999", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code)
	})

	t.Run("PostQuotes should return 400 for missing content", func(t *testing.T) {
		router := gin.New()
		RegisterRoutes(router, db)

		body := bytes.NewBuffer([]byte(`{}`))

		req := httptest.NewRequest(http.MethodPost, "/api/v1/quotes", body)
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})
}
