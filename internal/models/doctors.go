package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"github.com/sev-2/raiden/pkg/postgres"
)

type Doctors struct {
	db.ModelBase
	Id             uuid.UUID          `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	UserId         uuid.UUID          `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable:false"`
	Specialization string             `json:"specialization,omitempty" column:"name:specialization;type:text;nullable:false"`
	CreatedAt      *postgres.DateTime `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable;default:CURRENT_TIMESTAMP"`
	CreatedBy      *uuid.UUID         `json:"created_by,omitempty" column:"name:created_by;type:uuid;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"doctors" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	FacilitiesThroughSchedulesDoctor []*Facilities `json:"facilities_through_schedules_doctor,omitempty" join:"joinType:manyToMany;through:schedules;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
	ScheduleDoctors                  []*Schedules  `json:"schedule_doctors,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:doctor_id"`
	UserCreated                      *Users        `json:"user_created,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:created_by"`
	User                             *Users        `json:"user,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
	UsersThroughSchedulesDoctor      []*Users      `json:"users_through_schedules_doctor,omitempty" join:"joinType:manyToMany;through:schedules;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
}
