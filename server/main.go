package main

import (
	"fmt"
	"poll/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

func main() {
	app := echo.New()
	app.Pre(middleware.RemoveTrailingSlash())
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "${method}\t${status}\t${latency_human}\t${uri}\n"}))
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())
	app.Use(middleware.Gzip())
	app.Use(middleware.Secure())
	app.Use(middleware.BodyLimit("1M"))

	routes.RegisterQuestionRoutes(app)
	app.GET("/ws", hello)
	app.GET("/", homePage)

	app.Start(":8000")
}

func homePage(c echo.Context) error {
	// serve the index.html
	return c.File("index.html")
}

func hello(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// Write
			err := websocket.Message.Send(ws, "Hello, Client!")
			if err != nil {
				c.Logger().Error(err)
			}

			// Read
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
			}
			fmt.Printf("%s\n", msg)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
