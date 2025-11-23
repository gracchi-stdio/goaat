package handlers

import (
	"context"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/gracchi-stdio/goaat/internal/platform/db"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

// Auth initiates the OAuth flow
func (h *Handler) Auth(c echo.Context) error {
	provider := c.Param("provider")
	if provider == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Provider not specified")
	}

	// Add provider to context for gothic
	req := c.Request()
	// Use gothic.ProviderParamKey to avoid SA1029 and ensure gothic finds it
	ctx := context.WithValue(req.Context(), gothic.ProviderParamKey, provider)
	c.SetRequest(req.WithContext(ctx))

	h.AuthService.BeginAuth(c.Response(), c.Request())
	return nil
}

// AuthCallback handles the OAuth callback
func (h *Handler) AuthCallback(c echo.Context) error {
	provider := c.Param("provider")
	if provider == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Provider not specified")
	}

	// Add provider to context for gothic
	req := c.Request()
	ctx := context.WithValue(req.Context(), gothic.ProviderParamKey, provider)
	c.SetRequest(req.WithContext(ctx))

	user, err := h.AuthService.CompleteAuth(c.Response(), c.Request())
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	// Upsert user in database
	var avatarURL pgtype.Text
	if user.AvatarURL != "" {
		avatarURL = pgtype.Text{String: user.AvatarURL, Valid: true}
	} else {
		avatarURL = pgtype.Text{Valid: false}
	}

	dbUser, err := h.DB.UpsertUser(c.Request().Context(), db.UpsertUserParams{
		GithubID:  user.UserID,
		Email:     user.Email,
		Name:      user.Name,
		AvatarUrl: avatarURL,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to save user")
	}

	// Create a session for the user
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["user_id"] = dbUser.ID
	sess.Values["email"] = dbUser.Email
	sess.Values["name"] = dbUser.Name
	if dbUser.AvatarUrl.Valid {
		sess.Values["avatar_url"] = dbUser.AvatarUrl.String
	}
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to save session")
	}

	// Check for return_to URL
	returnTo := "/"
	if url, ok := sess.Values["return_to"].(string); ok && url != "" {
		returnTo = url
		delete(sess.Values, "return_to")
		sess.Save(c.Request(), c.Response())
	}

	return c.Redirect(http.StatusTemporaryRedirect, returnTo)
}

// Logout clears the session
func (h *Handler) Logout(c echo.Context) error {
	provider := c.Param("provider")
	if provider != "" {
		req := c.Request()
		ctx := context.WithValue(req.Context(), gothic.ProviderParamKey, provider)
		c.SetRequest(req.WithContext(ctx))
		h.AuthService.Logout(c.Response(), c.Request())
	}

	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	if c.Request().Header.Get("HX-Request") == "true" {
		c.Response().Header().Set("HX-Redirect", "/")
		return c.NoContent(http.StatusOK)
	}

	return c.Redirect(http.StatusSeeOther, "/")
}
