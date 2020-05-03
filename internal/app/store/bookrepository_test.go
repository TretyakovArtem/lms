package store_test

import (
	"testing"

	"github.com/TretyakovArtem/lms/internal/app/model"
	"github.com/TretyakovArtem/lms/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestBookRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("books")

	_, err := s.Book().Index(&model.Book{
		Name: "Война и мир",
		Year: "1823",
	})

	assert.NoError(t, err)
	// assert.NotNil(t, c)

}
