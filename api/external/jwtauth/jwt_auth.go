package jwtauth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// 64 random hexadecimal characters (0-9 and A-F):
// https://www.grc.com/passwords.htm
const SECRET_KEY = "D01287268E2B030F0AE20D718F1EE60CA88FCDEC3D7F9E8368DF2209E90DC4E0"

type JwtCustomClaims struct {
	Name string
	Password string
	jwt.StandardClaims
}

func Init(g *echo.Group) {
	config := middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(SECRET_KEY),
	}
	g.Use(middleware.JWTWithConfig(config))
}