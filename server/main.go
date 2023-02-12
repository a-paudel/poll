package main

import (
	"embed"
	"io/fs"
	"net/http"
	"poll/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed all:.embedded/client
var clientFiles embed.FS

func registerClientFiles(app *echo.Echo) {
	var files, _ = fs.Sub(clientFiles, ".embedded/client")

	app.Use(middleware.StaticWithConfig(
		middleware.StaticConfig{
			HTML5:      true,
			Filesystem: http.FS(files),
		},
	))
}

func main() {
	var app = echo.New()
	app.Pre(middleware.RemoveTrailingSlash())

	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "${time_rfc3339}\t${latency_human}\t${status}\t${method}\t${uri}\n"}))
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())

	registerClientFiles(app)
	routes.RegisterQuestionRoutes(app)

	for _, route := range app.Routes() {
		println(route.Method + "\t" + route.Path)
	}
	app.Start(":8000")

}
