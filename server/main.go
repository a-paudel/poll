package main

import (
	"poll/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var app = echo.New()
	app.Pre(middleware.RemoveTrailingSlash())

	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "${time_rfc3339}\t${latency_human}\t${status}\t${method}\t${uri}\n"}))
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())

	routes.RegisterQuestionRoutes(app)

	for _, route := range app.Routes() {
		println(route.Method + "\t" + route.Path)
	}
	app.Start(":8000")

}
