package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"github.com/sev-2/raiden/pkg/postgres"
)

type Reservations struct {
	db.ModelBase
	Id         uuid.UUID          `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	UserId     uuid.UUID          `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable:false"`
	ScheduleId uuid.UUID          `json:"schedule_id,omitempty" column:"name:schedule_id;type:uuid;nullable:false"`
	ServiceId  uuid.UUID          `json:"service_id,omitempty" column:"name:service_id;type:uuid;nullable:false"`
	Status     string             `json:"status,omitempty" column:"name:status;type:text;nullable:false"`
	CreatedAt  *postgres.DateTime `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable;default:CURRENT_TIMESTAMP"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"reservations" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	PaymentReservations []*Payments `json:"payment_reservations,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:reservation_id"`
	Schedule            *Schedules  `json:"schedule,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:schedule_id"`
	Service             *Services   `json:"service,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:service_id"`
	User                *Users      `json:"user,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
}
