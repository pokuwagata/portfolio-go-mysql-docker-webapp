package model

type Session struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
	Token string
}