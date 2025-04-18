// Code generated by raiden-cli; DO NOT EDIT.
package bootstrap

import (
	"github.com/sev-2/raiden/pkg/resource"
	"medpoint/internal/models"
)

func RegisterModels() {
	resource.RegisterModels(
		&models.Doctors{},
		&models.Facilities{},
		&models.Notifications{},
		&models.Payments{},
		&models.Reservations{},
		&models.Schedules{},
		&models.Services{},
		&models.Users{},
	)
}
