package teststore_test

import (
	"testing"

	"github.com/p4thf1nderr/lms/internal/app/model"
	"github.com/p4thf1nderr/lms/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestBookRepository_Create(t *testing.T) {
	s := teststore.New()
	b := model.TestBook(t)

	assert.NoError(t, s.Book().Create(b))
	assert.NotNil(t, b)
}

func TestBookRepository_Index(t *testing.T) {
	s := teststore.New()

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
