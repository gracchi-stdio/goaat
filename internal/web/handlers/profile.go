package handlers

import (
	"github.com/gracchi-stdio/goaat/internal/web/templates/components"
	"github.com/gracchi-stdio/goaat/internal/web/templates/pages"
	"github.com/labstack/echo/v4"
	"github.com/starfederation/datastar-go/datastar"
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

	sse := datastar.NewSSE(c.Response().Writer, c.Request())
	return sse.PatchElementTempl(components.Toast("Profile updated successfully", "success"))
}
