package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/nakamurakzz/flash-leaner/api/oapi"
)

func main() {
	server := NewServer()

	e := echo.New()
	e.Use(Logger())
	oapi.RegisterHandlers(e, server)
	e.HTTPErrorHandler = ErrorHandler()

	// Serve static files
	e.Static("/static", "web/static")

	// And we serve HTTP until the world ends.
	log.Fatal(e.Start("0.0.0.0:8000"))
}
