package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/labstack/gommon/color"
	"github.com/labstack/gommon/log"
)

// NewColorful creates a colorful Echo logger with custom formatting
func NewColorful() *log.Logger {
	l := log.New("app")
	l.SetLevel(log.DEBUG)
	l.SetOutput(os.Stdout)
	l.SetHeader("${time_rfc3339} ${level} ${short_file}:${line}")

	// Enable color globally
	color.Enable()

	return l
}

// ColorWriter wraps an io.Writer to add ANSI colors to log levels
type ColorWriter struct {
	w io.Writer
}

// NewColorWriter creates a new ColorWriter
func NewColorWriter(w io.Writer) *ColorWriter {
	return &ColorWriter{w: w}
}

// Write implements io.Writer and colorizes log levels
func (cw *ColorWriter) Write(p []byte) (n int, err error) {
	msg := string(p)

	// Color codes
	cyan := "\033[36m"
	green := "\033[32m"
	yellow := "\033[33m"
	red := "\033[31m"
	reset := "\033[0m"

	// Replace log levels with colored versions
	msg = replaceLevel(msg, "DEBUG", green+"DEBUG"+reset)
	msg = replaceLevel(msg, "INFO", cyan+"INFO"+reset)
	msg = replaceLevel(msg, "WARN", yellow+"WARN"+reset)
	msg = replaceLevel(msg, "ERROR", red+"ERROR"+reset)

	return fmt.Fprint(cw.w, msg)
}

// replaceLevel replaces the first occurrence of the level string
func replaceLevel(s, level, colored string) string {
	// Look for " LEVEL " pattern (with spaces)
	pattern := " " + level + " "
	coloredPattern := " " + colored + " "

	for i := 0; i <= len(s)-len(pattern); i++ {
		if s[i:i+len(pattern)] == pattern {
			return s[:i] + coloredPattern + s[i+len(pattern):]
		}
	}
	return s
}
