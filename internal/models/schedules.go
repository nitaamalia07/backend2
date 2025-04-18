package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"github.com/sev-2/raiden/pkg/postgres"
)

type Schedules struct {
	db.ModelBase
	Id         uuid.UUID         `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	DoctorId   uuid.UUID         `json:"doctor_id,omitempty" column:"name:doctor_id;type:uuid;nullable:false"`
	FacilityId uuid.UUID         `json:"facility_id,omitempty" column:"name:facility_id;type:uuid;nullable:false"`
	Date       postgres.DateTime `json:"date,omitempty" column:"name:date;type:date;nullable:false"`
	Time       postgres.DateTime `json:"time,omitempty" column:"name:time;type:time;nullable:false"`
	CreatedBy  *uuid.UUID        `json:"created_by,omitempty" column:"name:created_by;type:uuid;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"schedules" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Doctor                              *Doctors        `json:"doctor,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:doctor_id"`
	Facility                            *Facilities     `json:"facility,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:facility_id"`
	ReservationSchedules                []*Reservations `json:"reservation_schedules,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:schedule_id"`
	ServicesThroughReservationsSchedule []*Services     `json:"services_through_reservations_schedule,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:schedule_id;targetPrimaryKey:id;targetForeign:schedule_id"`
	UserCreated                         *Users          `json:"user_created,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:created_by"`
	UsersThroughReservationsSchedule    []*Users        `json:"users_through_reservations_schedule,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:schedule_id;targetPrimaryKey:id;targetForeign:schedule_id"`
}
