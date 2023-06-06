package status

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestHealthCheck(t *testing.T) {
	e := echo.New()
	e.GET("/api/v1/products/status", HealthCheck())

	req := httptest.NewRequest(http.MethodGet, "/api/v1/products/status", http.NoBody)
	resp := httptest.NewRecorder()

	e.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("status code was expected %d, but it was obtained %d", http.StatusOK, resp.Code)
	}
}
