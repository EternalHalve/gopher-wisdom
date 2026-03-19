package quotes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
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

func TestRateLimiter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := setupTestDB()
	router := gin.New()

	// 1. Setup a strict limiter for the test (2 requests per second)
	rate := limiter.Rate{
		Period: 1 * time.Second,
		Limit:  2,
	}
	store := memory.NewStore()
	instance := limiter.New(store, rate)
	middleware := mgin.NewMiddleware(instance)

	// 2. Attach middleware and routes
	router.Use(middleware)
	RegisterRoutes(router, db)

	t.Run("should allow first 2 requests and block the 3rd", func(t *testing.T) {
		// Request 1: OK
		resp1 := httptest.NewRecorder()
		req1, _ := http.NewRequest(http.MethodGet, "/api/v1/quotes", nil)
		router.ServeHTTP(resp1, req1)
		assert.Equal(t, http.StatusOK, resp1.Code)

		// Request 2: OK
		resp2 := httptest.NewRecorder()
		req2, _ := http.NewRequest(http.MethodGet, "/api/v1/quotes", nil)
		router.ServeHTTP(resp2, req2)
		assert.Equal(t, http.StatusOK, resp2.Code)

		// Request 3: TOO MANY REQUESTS
		resp3 := httptest.NewRecorder()
		req3, _ := http.NewRequest(http.MethodGet, "/api/v1/quotes", nil)
		router.ServeHTTP(resp3, req3)

		// This confirms the rate limiter is working
		assert.Equal(t, http.StatusTooManyRequests, resp3.Code)
		assert.Contains(t, resp3.Body.String(), "Limit exceeded")
	})
}
