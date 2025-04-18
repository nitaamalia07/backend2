package types

import (
	"github.com/sev-2/raiden"
)

type Notifications struct {
	raiden.TypeBase
}

func (t *Notifications) Name() string {
	return "_notifications"
}

func (r *Notifications) Format() string {
	return "notifications[]"
}

func (r *Notifications) Enums() []string {
	return []string{}
}

func (r *Notifications) Comment() *string {
	return nil
}

