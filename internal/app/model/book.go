package model

// Book ...
type Book struct {
	ID   uint64 `db:"id"`
	Name string `db:"name"`
	Year string `db:"year"`
}
