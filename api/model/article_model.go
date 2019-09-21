package model

import (
	"time"
)

type Article struct {
	ID int64
	Title string `validate:"required"`
	Content string `validate:"required"`
	PostId int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Username string
	ArticleStatus string
}