package db

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func (rep *Repository) LinkAdd(ctx context.Context, url string) (int, error) {
	//language=SQL
	sql := `
	WITH inserted AS (
		INSERT INTO link (url, status)
		VALUES ($1, 'pending')
		ON CONFLICT (url) DO NOTHING
		RETURNING id
	)
	SELECT id FROM inserted
	UNION ALL
	SELECT id FROM link WHERE url = $1
	LIMIT 1;`

	var id int
	err := rep.pool.QueryRow(ctx, sql, url).Scan(&id)
	return id, err
}

func (rep *Repository) LinkGetByID(ctx context.Context, id int) (LinkRow, error) {
	// language=SQL
	sql := `
	SELECT id, url, status
	FROM link
	WHERE id = $1
	`

	var link LinkRow
	err := rep.pool.QueryRow(ctx, sql, id).Scan(&link.ID, &link.Url, &link.Status)
	return link, err
}

func (rep *Repository) LinkList(ctx context.Context) ([]LinkRow, error) {
	sql := `
	SELECT id, url, status
	FROM link
	`

	var links []LinkRow

	rows, err := rep.pool.Query(ctx, sql)
	if err != nil {
		return links, err
	}
	defer rows.Close()

	for rows.Next() {
		var link LinkRow
		err := rows.Scan(&link.ID, &link.Url, &link.Status)
		if err != nil {
			return links, err
		}
		links = append(links, link)
	}

	return links, err
}

func (rep *Repository) LinkImageAddMultiple(ctx context.Context, linkID int, images []string) error {
	//language=SQL
	sql := `
	INSERT INTO link_image (link_id, url)
	VALUES ($1, $2)
	`

	batch := &pgx.Batch{}
	for _, image := range images {
		batch.Queue(sql, linkID, image)
	}
	br := rep.pool.SendBatch(ctx, batch)
	defer br.Close()

	_, err := br.Exec()
	return err
}
