package model

// Instance ...
type Instance struct {
	ID         uint   `db:"id"`
	ProducerID uint   `db:"producer_id"`
	BookID     uint   `db:"book_id"`
	Year       string `db:"year"`
}
