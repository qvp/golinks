// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package sqlc

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Link struct {
	ID          int32
	Url         string
	Status      pgtype.Text
	CreatedAt   pgtype.Timestamp
	ProcessedAt pgtype.Timestamp
}

type LinkImage struct {
	ID     int32
	LinkID int32
	Url    string
}
