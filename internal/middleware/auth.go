package middleware

import (
	"context"
	"net/http"

	"github.com/gracchi-stdio/goaat/internal/auth"
	"github.com/labstack/echo/v4"
)

// RequireAuth checks if the user is logged in
func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userSession := auth.GetSession(c)
		if !userSession.IsAuthenticated() {
			// Log the unauthorized access attempt using the custom logger
			c.Logger().Warnf("Unauthorized access attempt to: %s", c.Request().URL.String())

			// Save the current URL to redirect back after login
			auth.SetReturnTo(c, c.Request().URL.String())
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
		c.Logger().Infof("Authenticated access by user: %s", userSession.Name)

		// Inject user into context for templates
		ctx := context.WithValue(c.Request().Context(), auth.UserContextKey, userSession)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
