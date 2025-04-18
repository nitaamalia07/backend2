package types

import (
	"github.com/sev-2/raiden"
)

type UserRole struct {
	raiden.TypeBase
}

func (t *UserRole) Name() string {
	return "user_role"
}

func (r *UserRole) Format() string {
	return "user_role"
}

func (r *UserRole) Enums() []string {
	return []string{"superadmin","admin","doctor","patient"}
}

func (r *UserRole) Comment() *string {
	return nil
}

