package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Book ...
type Book struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
	Year  string `json:"year"`
}

// Validate ...
func (b *Book) Validate() error {
	return validation.ValidateStruct(
		b,
		validation.Field(&b.Title, validation.Required, validation.Length(1, 255)),
		validation.Field(&b.Year, validation.Required, validation.Length(4, 4)),
	)
}
