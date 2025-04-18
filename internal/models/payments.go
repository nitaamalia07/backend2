package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"github.com/sev-2/raiden/pkg/postgres"
)

type Payments struct {
	db.ModelBase
	Id            uuid.UUID          `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	ReservationId uuid.UUID          `json:"reservation_id,omitempty" column:"name:reservation_id;type:uuid;nullable:false"`
	Method        string             `json:"method,omitempty" column:"name:method;type:text;nullable:false"`
	Amount        float64            `json:"amount,omitempty" column:"name:amount;type:numeric;nullable:false"`
	Status        string             `json:"status,omitempty" column:"name:status;type:text;nullable:false"`
	PaidAt        *postgres.DateTime `json:"paid_at,omitempty" column:"name:paid_at;type:timestamp;nullable"`
	CreatedAt     *postgres.DateTime `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable;default:CURRENT_TIMESTAMP"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"payments" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Reservation *Reservations `json:"reservation,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:reservation_id"`
}
