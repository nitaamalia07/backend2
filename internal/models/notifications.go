package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"github.com/sev-2/raiden/pkg/postgres"
)

type Notifications struct {
	db.ModelBase
	Id        uuid.UUID          `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	UserId    uuid.UUID          `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable:false"`
	Title     string             `json:"title,omitempty" column:"name:title;type:text;nullable:false"`
	Message   string             `json:"message,omitempty" column:"name:message;type:text;nullable:false"`
	IsRead    *bool              `json:"is_read,omitempty" column:"name:is_read;type:boolean;nullable;default:false"`
	CreatedAt *postgres.DateTime `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable;default:CURRENT_TIMESTAMP"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"notifications" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	User *Users `json:"user,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
}
