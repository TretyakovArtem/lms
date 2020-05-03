package model

import "testing"

// TestBook ...
func TestBook(t *testing.T) *Book {
	return &Book{
		Title: "War and Piece",
		Year:  "1823",
	}
}
