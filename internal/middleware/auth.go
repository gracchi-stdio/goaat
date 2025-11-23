package middleware

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// RequireAuth checks if the user is logged in
func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		if sess.Values["user_id"] == nil {
			// Save the current URL to redirect back after login
			sess.Values["return_to"] = c.Request().URL.String()
			sess.Save(c.Request(), c.Response())
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
		return next(c)
	}
}
