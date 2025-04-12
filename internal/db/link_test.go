package db

import (
	"context"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueries_LinkAdd(t *testing.T) {
	url := gofakeit.URL()
	link, err := queries.LinkAdd(context.Background(), url)
	assert.NoError(t, err)
	assert.Equal(t, link.Url, url)
}

func TestQueries_AddLinkImages(t *testing.T) {
	link, err := queries.LinkAdd(context.Background(), gofakeit.URL())
	assert.NoError(t, err)

	images := []string{
		"https://test.com/cat.png",
	}
	err = queries.SaveLinkImagesTx(context.Background(), int(link.ID), images)
	assert.NoError(t, err)
}
