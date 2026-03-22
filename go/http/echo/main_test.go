package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func BenchmarkDownload(b *testing.B) {
	e := echo.New()

	// Подготовка данных формы
	f := make(url.Values)
	f.Set("name", "test")
	f.Set("email", "test@example.com")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/download", strings.NewReader(f.Encode()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := bindFile(c)
		if err != nil {
			b.Fatalf("handler error: %v", err)
		}

		if rec.Code != http.StatusOK {
			b.Fatalf("expected status 200, got %d", rec.Code)
		}
	}
}

func BenchmarkDownloadWithMiddleware(b *testing.B) {
	e := echo.New()
	h := middleware(bindFile)

	f := make(url.Values)
	f.Set("name", "test")
	f.Set("email", "test@example.com")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/download", strings.NewReader(f.Encode()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := h(c)
		if err != nil {
			b.Fatalf("handler error: %v", err)
		}

		if rec.Code != http.StatusOK {
			b.Fatalf("expected status 200, got %d", rec.Code)
		}
	}
}
