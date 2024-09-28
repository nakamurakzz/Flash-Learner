package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

// Logger returns a middleware that logs HTTP requests using slog
func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			// Process request
			err := next(c)

			// Log request details
			req := c.Request()
			res := c.Response()

			slog.Info("HTTP request",
				"method", req.Method,
				"path", req.URL.Path,
				"status", res.Status,
				"duration", time.Since(start),
				"ip", c.RealIP(),
				"user-agent", req.UserAgent(),
			)

			return err
		}
	}
}

// ConfigureLogger sets up the global logger with the specified options
func ConfigureLogger(level slog.Level, addSource bool) {
	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: addSource,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)
}

// ErrorHandler returns a custom HTTP error handler
func ErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		message := "Internal Server Error"

		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			message = he.Message.(string)
		}

		// Log all 5xx errors
		if code >= 500 {
			slog.Error("Server error",
				"error", err.Error(),
				"path", c.Request().URL.Path,
				"method", c.Request().Method,
				"status", code,
			)
		}

		// Send the error response
		if !c.Response().Committed {
			if c.Request().Method == http.MethodHead {
				err = c.NoContent(code)
			} else {
				err = c.JSON(code, map[string]string{"error": message})
			}
			if err != nil {
				slog.Error("Failed to send error response", "error", err.Error())
			}
		}
	}
}
