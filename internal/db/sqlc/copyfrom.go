// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: copyfrom.go

package sqlc

import (
	"context"
)

// iteratorForLinkImageAddMultiple implements pgx.CopyFromSource.
type iteratorForLinkImageAddMultiple struct {
	rows                 []LinkImageAddMultipleParams
	skippedFirstNextCall bool
}

func (r *iteratorForLinkImageAddMultiple) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForLinkImageAddMultiple) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].LinkID,
		r.rows[0].Url,
	}, nil
}

func (r iteratorForLinkImageAddMultiple) Err() error {
	return nil
}

func (q *Queries) LinkImageAddMultiple(ctx context.Context, arg []LinkImageAddMultipleParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"link_image"}, []string{"link_id", "url"}, &iteratorForLinkImageAddMultiple{rows: arg})
}
