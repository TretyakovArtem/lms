package model

// Client ..
type Client struct {
	ID      uint   `db:"id"`
	Name    string `db:"name"`
	Address string `db:"address"`
	Phone   string `db:"phone"`
	Email   string `db:"email"`
}
