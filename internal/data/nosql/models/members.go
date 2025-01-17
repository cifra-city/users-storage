package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/users-storage/internal/service/roles"
)

type Member struct {
	ID          uuid.UUID      `bson:"_id" json:"id"`
	UserId      uuid.UUID      `bson:"user_id" json:"user_id"`
	Role        roles.TeamRole `bson:"role" json:"role"`
	Description string         `bson:"description,omitempty" json:"description,omitempty"`
	CreatedAt   time.Time      `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `bson:"updated_at" json:"updated_at"`
}
