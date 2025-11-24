package auth

import (
	"context"
	"encoding/gob"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const (
	SessionName = "session"
	UserKey     = "user_session"
)

type contextKey string

const UserContextKey contextKey = "user"

func init() {
	// Register the struct so gorilla/sessions can serialize it using gob
	gob.Register(UserSession{})
}

// UserSession represents the typed session data
type UserSession struct {
	UserID    int64
	Email     string
	Name      string
	AvatarURL string
	ReturnTo  string
}

// IsAuthenticated checks if the user is logged in
func (s UserSession) IsAuthenticated() bool {
	return s.UserID != 0
}

// GetUserFromContext retrieves the user session from the context
func GetUserFromContext(ctx context.Context) UserSession {
	if user, ok := ctx.Value(UserContextKey).(UserSession); ok {
		return user
	}
	return UserSession{}
}

// GetSession retrieves the current user session
func GetSession(c echo.Context) UserSession {
	sess, _ := session.Get(SessionName, c)

	// Try to get the struct directly from the session
	if val, ok := sess.Values[UserKey].(UserSession); ok {
		return val
	}

	// Return empty session if not found
	return UserSession{}
}

// SaveSession saves the user session
func SaveSession(c echo.Context, s UserSession) error {
	sess, _ := session.Get(SessionName, c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
	}

	// Store the entire struct
	sess.Values[UserKey] = s

	return sess.Save(c.Request(), c.Response())
}

// ClearSession invalidates the session (logout)
func ClearSession(c echo.Context) error {
	sess, _ := session.Get(SessionName, c)
	sess.Options.MaxAge = -1
	return sess.Save(c.Request(), c.Response())
}

// SetReturnTo saves the return URL in the session
func SetReturnTo(c echo.Context, url string) error {
	s := GetSession(c)
	s.ReturnTo = url
	return SaveSession(c, s)
}
