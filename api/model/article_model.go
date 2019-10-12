package model

import (
	"time"
)

type Article struct {
	ID            int64
	Title         string `validate:"required"`
	Content       string `validate:"required"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Username      string
	ArticleStatus string
}

type ViewArticle struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updatedAt"`
	Username  string    `json:"username"`
}

type FirstGetListResponse struct {
	Articles []ViewArticle `json:"articles"`
	Max      int           `json:"maxNumber"`
}

type GetListResponse struct {
	Articles []ViewArticle `json:"articles"`
}
