package store_test

import (
	"testing"

	"github.com/TretyakovArtem/lms/internal/app/model"
	"github.com/TretyakovArtem/lms/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestCustomerRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("customers")

	_, err := s.Customer().Create(&model.Customer{
		Name:    "Артем Третьяков",
		Address: "ssss",
		Phone:   "878908908",
	})

	assert.NoError(t, err)
	// assert.NotNil(t, c)

}

func TestCustomerRepository_FindByName(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("customers")

	name := "Третьяков"

	_, err := s.Customer().FindByName(name)
	assert.Error(t, err)

	s.Customer().Create(&model.Customer{
		Name:    "Артем Третьяков",
		Address: "ssss",
		Phone:   "878908908",
	})

	c, err := s.Customer().FindByName(name)
	assert.NoError(t, err)
	assert.NotNil(t, c)

}
