package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type Services struct {
	db.ModelBase
	Id          uuid.UUID `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	FacilityId  uuid.UUID `json:"facility_id,omitempty" column:"name:facility_id;type:uuid;nullable:false"`
	Name        string    `json:"name,omitempty" column:"name:name;type:text;nullable:false"`
	Description *string   `json:"description,omitempty" column:"name:description;type:text;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"services" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Facility                            *Facilities     `json:"facility,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:facility_id"`
	ReservationServices                 []*Reservations `json:"reservation_services,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:service_id"`
	SchedulesThroughReservationsService []*Schedules    `json:"schedules_through_reservations_service,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:service_id;targetPrimaryKey:id;targetForeign:service_id"`
	UsersThroughReservationsService     []*Users        `json:"users_through_reservations_service,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:service_id;targetPrimaryKey:id;targetForeign:service_id"`
}
