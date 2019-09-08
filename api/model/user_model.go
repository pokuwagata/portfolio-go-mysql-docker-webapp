package model

import (
	"time"
)

type User struct {
	ID int
	Username string `validate:"required"`
	Password string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Status string
}
