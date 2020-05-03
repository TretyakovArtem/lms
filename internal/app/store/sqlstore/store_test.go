package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL_TEST")
	if databaseURL == "" {
		databaseURL = "host=localhost user=pathfinder password=qaz123 dbname=library_test sslmode=disable"
	}

	os.Exit(m.Run())
}
