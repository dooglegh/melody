package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/olahol/melody.v1"
	"log"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.File("/", "index.html")
	e.File("/channel/:name", "chan.html")

	m := melody.New()

	e.GET("/channel/:name/ws", func(c echo.Context) error {
		m.HandleRequest(c.Response().Writer, c.Request())
		return nil
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		// m.Broadcast(msg)
		log.Println("[OnMessage]", string(msg))
		m.BroadcastFilter(msg, func(q *melody.Session) bool {
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})

	e.Start(":5000")
}
