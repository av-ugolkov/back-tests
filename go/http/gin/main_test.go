package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
)

func BenchmarkDownload(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)

	// Подготовка данных формы
	f := make(url.Values)
	f.Set("name", "test")
	f.Set("email", "test@example.com")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/download?"+f.Encode(), nil)
		rec := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		bindFile(c)

		if rec.Code != http.StatusOK {
			b.Fatalf("expected status 200, got %d", rec.Code)
		}
	}
}

func BenchmarkDownloadWithMiddleware(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)

	// Подготовка данных формы
	f := make(url.Values)
	f.Set("name", "test")
	f.Set("email", "test@example.com")

	handler := middleware(bindFile)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/download?"+f.Encode(), nil)
		rec := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		handler(c)

		if rec.Code != http.StatusOK {
			b.Fatalf("expected status 200, got %d", rec.Code)
		}
	}
}
