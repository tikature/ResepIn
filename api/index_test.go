package handler

import (
	"net/http/httptest"
	"testing"
)

func TestPlaceholderJpgServesSvgFallback(t *testing.T) {
	for _, path := range []string{
		"/images/placeholder.jpg",
		"/static/images/placeholder.jpg",
	} {
		t.Run(path, func(t *testing.T) {
			response := httptest.NewRecorder()
			request := httptest.NewRequest("GET", path, nil)

			serveStatic(response, request)

			if response.Code != 200 {
				t.Fatalf("status = %d, want 200", response.Code)
			}
			if got := response.Header().Get("Content-Type"); got != "image/svg+xml" {
				t.Fatalf("Content-Type = %q, want %q", got, "image/svg+xml")
			}
		})
	}
}
