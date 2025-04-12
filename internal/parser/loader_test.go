package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadHtml_Success(t *testing.T) {
	ts := testHtmlPageServer()
	defer ts.Close()

	content, err := LoadHtml(ts.URL)

	assert.NoError(t, err)
	assert.NotZero(t, len(content))
}

func TestLoadHtml_Error(t *testing.T) {
	content, err := LoadHtml("http://fake.url")

	assert.Error(t, err)
	assert.Len(t, content, 0)
}
