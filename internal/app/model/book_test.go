package model_test

import (
	"testing"

	"github.com/p4thf1nderr/lms/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestBook_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		b       func() *model.Book
		isValid bool
	}{
		{
			name: "valid",
			b: func() *model.Book {
				return model.TestBook(t)
			},
			isValid: true,
		},
		{
			name: "empty title",
			b: func() *model.Book {
				book := model.TestBook(t)
				book.Title = ""

				return book
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.b().Validate())
			} else {
				assert.Error(t, tc.b().Validate())
			}
		})
	}
}
