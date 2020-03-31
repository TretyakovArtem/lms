package store

import (
	"github.com/TretyakovArtem/lms/internal/app/model"
)

// CustomerRepository ...
type CustomerRepository struct {
	store *Store
}

// Create метод для занесения в базу читателя
func (r *CustomerRepository) Create(c *model.Customer) (*model.Customer, error) {

	r.store.db.QueryRow("INSERT INTO customers (name, address, phone) VALUES (?, ?, ?)", c.Name, c.Address, c.Phone)
	return nil, nil
}

// List список читателей
func (r *CustomerRepository) List() ([]model.Customer, error) {

	rows, err := r.store.db.Query(`SELECT id, name, address, phone FROM customers`)
	if err != nil {
		return nil, err
	}

	var customers []model.Customer

	for rows.Next() {
		var customer model.Customer
		_ = rows.Scan(&customer.ID, &customer.Name, &customer.Address, &customer.Phone)
		customers = append(customers, customer)
	}

	return customers, nil
}

// FindByName ..
func (r *CustomerRepository) FindByName(name string) (*model.Customer, error) {
	c := &model.Customer{}

	if err := r.store.db.QueryRow("SELECT id, name, address, phone FROM customers WHERE name LIKE CONCAT('%', ?, '%')", name).Scan(
		&c.ID,
		&c.Name,
		&c.Address,
		&c.Phone,
	); err != nil {
		return nil, err
	}

	return c, nil
}

// FindByID ..
func (r *CustomerRepository) FindByID(id string) (*model.Customer, error) {
	c := &model.Customer{}

	if err := r.store.db.QueryRow("SELECT id, name, address, phone FROM customers WHERE id = ?", id).Scan(
		&c.ID,
		&c.Name,
		&c.Address,
		&c.Phone,
	); err != nil {
		return nil, err
	}

	return c, nil
}

// Update метод для изменения данных в базе читателя
func (r *CustomerRepository) Update(c *model.Customer) (*model.Customer, error) {
	r.store.db.QueryRow("UPDATE customers SET name=?, address=?, phone=? WHERE id=?", c.Name, c.Address, c.Phone, c.ID)
	return nil, nil
}

// Delete метод для изменения данных в базе читателя
func (r *CustomerRepository) Delete(cID int) (*model.Customer, error) {
	r.store.db.QueryRow("DELETE FROM customers WHERE id = ?", cID)
	return nil, nil
}
