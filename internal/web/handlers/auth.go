package handlers

import (
	"context"
	"net/http"

	"github.com/gracchi-stdio/goaat/internal/auth"
	"github.com/gracchi-stdio/goaat/internal/platform/db"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
	"github.com/starfederation/datastar-go/datastar"
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

	// Check DB availability before proceeding
	if h.DB == nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, "database unavailable")
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
	s := auth.GetSession(c)
	s.UserID = dbUser.ID
	s.Email = dbUser.Email
	s.Name = dbUser.Name
	if dbUser.AvatarUrl.Valid {
		s.AvatarURL = dbUser.AvatarUrl.String
	}

	// Check for return_to URL
	returnTo := "/"
	if s.ReturnTo != "" {
		returnTo = s.ReturnTo
		s.ReturnTo = "" // Clear it
	}

	if err := auth.SaveSession(c, s); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to save session")
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

	auth.ClearSession(c)

	// Check if this is a Datastar request (SSE) - Datastar sends Accept: text/event-stream
	accept := c.Request().Header.Get("Accept")
	if accept == "text/event-stream" || c.Request().Header.Get("Datastar-Request") != "" {
		sse := datastar.NewSSE(c.Response().Writer, c.Request())
		return sse.Redirect("/")
	}

	return c.Redirect(http.StatusSeeOther, "/")
}
