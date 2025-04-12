package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"golinks/internal/db/sqlc"
)

// SaveLinkImagesTx save images to database and change status of link to "processed"
func (q *Queries) SaveLinkImagesTx(ctx context.Context, linkID int, images []string) error {
	tx, err := q.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	qtx := q.WithTx(tx)

	err = qtx.LinkSetStatus(ctx, sqlc.LinkSetStatusParams{
		ID:     int32(linkID),
		Status: pgtype.Text{String: "processed", Valid: true},
	})
	if err != nil {
		return err
	}

	linkImages := make([]sqlc.LinkImageAddMultipleParams, 0)
	for _, imageUrl := range images {
		linkImages = append(linkImages, sqlc.LinkImageAddMultipleParams{
			LinkID: int32(linkID),
			Url:    imageUrl,
		})
	}
	_, err = qtx.LinkImageAddMultiple(ctx, linkImages)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
