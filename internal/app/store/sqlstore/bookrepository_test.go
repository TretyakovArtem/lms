package sqlstore_test

import (
	"testing"

	"github.com/p4thf1nderr/lms/internal/app/model"
	"github.com/p4thf1nderr/lms/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestBookRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("books")

	s := sqlstore.New(db)
	b := model.TestBook(t)

	assert.NoError(t, s.Book().Create(b))
	assert.NotNil(t, b)
}

func TestBookRepository_Index(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("books")

	s := sqlstore.New(db)

	err := s.Book().Create(&model.Book{
		Title: "Война и мир",
		Year:  "1823",
	})

	err = s.Book().Create(&model.Book{
		Title: "Горе от ума",
		Year:  "1900",
	})

	books, err := s.Book().Index()

	assert.NoError(t, err)
	assert.NotNil(t, books)

}
