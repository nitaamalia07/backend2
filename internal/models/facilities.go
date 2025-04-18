package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
)

type Facilities struct {
	db.ModelBase
	Id      uuid.UUID `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	Name    string    `json:"name,omitempty" column:"name:name;type:text;nullable:false"`
	Address *string   `json:"address,omitempty" column:"name:address;type:text;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"facilities" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	DoctorsThroughSchedulesFacility []*Doctors   `json:"doctors_through_schedules_facility,omitempty" join:"joinType:manyToMany;through:schedules;sourcePrimaryKey:id;sourceForeignKey:facility_id;targetPrimaryKey:id;targetForeign:facility_id"`
	ScheduleFacilities              []*Schedules `json:"schedule_facilities,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:facility_id"`
	ServiceFacilities               []*Services  `json:"service_facilities,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:facility_id"`
	UsersThroughSchedulesFacility   []*Users     `json:"users_through_schedules_facility,omitempty" join:"joinType:manyToMany;through:schedules;sourcePrimaryKey:id;sourceForeignKey:facility_id;targetPrimaryKey:id;targetForeign:facility_id"`
}
