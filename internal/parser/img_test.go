package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetImagesFromHtml_Success(t *testing.T) {
	images, err := GetImagesFromHtml(testHtmlPage)

	assert.NoError(t, err)
	assert.Len(t, images, 2)
}

func TestGetImagesFromHtml_Empty(t *testing.T) {
	images, err := GetImagesFromHtml(" no img tags here ")

	assert.NoError(t, err)
	assert.Len(t, images, 0)
}
