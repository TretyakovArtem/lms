package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE URL")
	if databaseURL == "" {
		databaseURL = "root:123123@tcp(127.0.0.1)/library?charset=utf8"
	}

	os.Exit(m.Run())
}
