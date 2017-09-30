/*
 * @author Philip Van Raalte
 * @date 2017-09-30.
 *
 * A web server, for single page applications, that uses Go and the Echo framework
 */
package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

/// Main function
/// Starts the serer
func main() {
	e := echo.New()

	//recover from panics and handles control to HTTPErrorHandler
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{}))

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 6,
	}))

	e.Static("/", "assets")

	e.HTTPErrorHandler = customHTTPErrorHandler

	e.Logger.Fatal(e.Start(":8000"))
}

/// Handles errors
/// Serves the index page which will handle routing
func customHTTPErrorHandler(_ error, c echo.Context) {
	c.File("assets/index.html")
}