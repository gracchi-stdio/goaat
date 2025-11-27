package handlers

import (
	"github.com/gracchi-stdio/goaat/internal/web/templates/pages"
	"github.com/labstack/echo/v4"
)

// ProfilePage renders the user profile page
func (h *Handler) ProfilePage(c echo.Context) error {
	if c.Request().Header.Get("datastar-request") != "" {
		return RenderWithDatastar(c, pages.ProfileContent())
	}
	return Render(c, pages.Profile())
}

// UpdateProfile handles profile updates with success alert
func (h *Handler) UpdateProfile(c echo.Context) error {
	// TODO: Implement actual profile update logic

	// For now, just show success alert and redirect
	c.Response().Header().Set("HX-Trigger", "profile-updated")

	// Redirect back to profile page
	return c.Redirect(302, "/admin/profile")
}
