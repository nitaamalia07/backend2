package middleware

import (
	"github.com/sev-2/raiden"
)

// AuthMiddleware: middleware untuk memeriksa validitas API Key
func AuthMiddleware(ctx raiden.Context) error {
	// Mengambil API Key dari header request
	apiKey := string(ctx.RequestContext().Request.Header.Peek("apikey"))

	// Jika API Key tidak ditemukan
	if apiKey == "" {
		return ctx.SendJson(map[string]string{
			"message": "API key tidak ditemukan",
		})
	}

	// Jika API Key tidak valid
	if apiKey != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImV5bHZid3pjYm5xcGFwdW10ZmRqIiwicm9sZSI6ImFub24iLCJpYXQiOjE3NDQxMTg5NDcsImV4cCI6MjA1OTY5NDk0N30.Jlz5TzTl_OH4s5jw8Kd6OOYdfq6OqtEXN0V_1U-Z32c" {
		return ctx.SendJson(map[string]string{
			"message": "API key salah",
		})
	}

	// Jika API Key valid, tidak ada error yang dikembalikan
	return nil
}
