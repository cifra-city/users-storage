// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlcore

import (
	"database/sql"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Username  string
	Title     sql.NullString
	Status    sql.NullString
	Avatar    sql.NullString
	Bio       sql.NullString
	City      uuid.NullUUID
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}
