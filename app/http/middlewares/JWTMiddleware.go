package middlewares

import (
	"github.com/labstack/echo/v4"
	"estructuraapibasegolang/app/services/auth"
	"github.com/labstack/echo/v4/middleware"
)

// JWTMiddleware función intermediaria para retornar el middleware JWT del framework Echo
func JWTMiddleware() echo.MiddlewareFunc {

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims: &auth.JWTCustomClaims{},
		ContextKey: "auth_token",
		SigningMethod: "RS256",
		SigningKey: auth.GetPublicKey(),
	}

	return middleware.JWTWithConfig(config)

}