package controllers

import (
	"medpoint/internal/middleware"
	"medpoint/internal/models"

	"github.com/sev-2/raiden"
)

type DoctorsController struct {
	raiden.ControllerBase
	Http  string `path:"/doctors" type:"rest"`
	Model models.Doctors
}

// Register route dalam Raiden (controller terdaftar otomatis)
func (d *DoctorsController) RegisterRoutes() {
	// Raiden akan menautkan /doctors ke metode controller ini
	d.Http = "/doctors"
}

// BeforeGet akan dipanggil sebelum method Get, pastikan middleware dijalankan sebelum data diambil
func (c *DoctorsController) BeforeGet(ctx raiden.Context) error {
	// Menggunakan middleware Auth
	if err := middleware.AuthMiddleware(ctx); err != nil {
		return err
	}

	// Mengecek API Key setelah autentikasi berhasil
	apiKey := ctx.RequestContext().Request.Header.Peek("apikey")
	if len(apiKey) == 0 {
		return ctx.SendJson(map[string]string{
			"message": "No API key found in request",
			"hint":    "No apikey request header or url param was found.",
		})
	}

	if string(apiKey) != "your-expected-api-key" {
		return ctx.SendJson(map[string]string{
			"message": "Invalid API key",
		})
	}

	return nil
}
