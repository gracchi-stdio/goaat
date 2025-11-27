package middleware

import (
	"context"

	"github.com/gracchi-stdio/goaat/internal/auth"
	"github.com/labstack/echo/v4"
)

// InjectUser middleware always injects the user session into context,
// whether authenticated or not. Use this for public pages that need
// to show different content based on auth state.
func InjectUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userSession := auth.GetSession(c)

		// Inject user into context (will be empty if not authenticated)
		ctx := context.WithValue(c.Request().Context(), auth.UserContextKey, userSession)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
