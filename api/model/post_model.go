package model

import (
	"time"
)

type Post struct {
	ID int64
	PostedAt time.Time
	Username string
}