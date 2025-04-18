package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"github.com/sev-2/raiden/pkg/postgres"
)

type Users struct {
	db.ModelBase
	Id        uuid.UUID          `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	Name      string             `json:"name,omitempty" column:"name:name;type:text;nullable:false"`
	Email     string             `json:"email,omitempty" column:"name:email;type:text;nullable:false;unique"`
	Role      string             `json:"role,omitempty" column:"name:role;type:text;nullable:false"`
	CreatedAt *postgres.DateTime `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable;default:CURRENT_TIMESTAMP"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"users" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	DoctorCreateds                      []*Doctors       `json:"doctor_createds,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:created_by"`
	DoctorUsers                         []*Doctors       `json:"doctor_users,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:user_id"`
	DoctorsThroughSchedulesCreatedBy    []*Doctors       `json:"doctors_through_schedules_created_by,omitempty" join:"joinType:manyToMany;through:schedules;sourcePrimaryKey:id;sourceForeignKey:created_by;targetPrimaryKey:id;targetForeign:created_by"`
	FacilitiesThroughSchedulesCreatedBy []*Facilities    `json:"facilities_through_schedules_created_by,omitempty" join:"joinType:manyToMany;through:schedules;sourcePrimaryKey:id;sourceForeignKey:created_by;targetPrimaryKey:id;targetForeign:created_by"`
	NotificationUsers                   []*Notifications `json:"notification_users,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:user_id"`
	ReservationUsers                    []*Reservations  `json:"reservation_users,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:user_id"`
	ScheduleCreateds                    []*Schedules     `json:"schedule_createds,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:created_by"`
	SchedulesThroughReservationsUser    []*Schedules     `json:"schedules_through_reservations_user,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:user_id;targetPrimaryKey:id;targetForeign:user_id"`
	ServicesThroughReservationsUser     []*Services      `json:"services_through_reservations_user,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:user_id;targetPrimaryKey:id;targetForeign:user_id"`
	UsersThroughDoctorsCreatedBy        []*Users         `json:"users_through_doctors_created_by,omitempty" join:"joinType:manyToMany;through:doctors;sourcePrimaryKey:id;sourceForeignKey:created_by;targetPrimaryKey:id;targetForeign:created_by"`
	UsersThroughDoctorsUser             []*Users         `json:"users_through_doctors_user,omitempty" join:"joinType:manyToMany;through:doctors;sourcePrimaryKey:id;sourceForeignKey:user_id;targetPrimaryKey:id;targetForeign:user_id"`
}
