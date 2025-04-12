package db

import (
	"os"
	"testing"
)

var queries *Queries

func TestMain(m *testing.M) {
	queries = NewQueries()
	defer ClosePool()

	exitCode := m.Run()

	os.Exit(exitCode)
}
