package model

// Author ...
type Author struct {
	ID   uint   `db:"id"`
	Name string `db:"name"`
	Year string `db:"year"`
}
