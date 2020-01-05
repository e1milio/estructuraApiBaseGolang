package auth

import "github.com/dgrijalva/jwt-go"

// JWTCustomClaims are custom claims extending default ones.
type JWTCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
