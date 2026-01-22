package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func serveStream() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		streamName := ctx.Param("live")
		filePath := ctx.Param("*")

		if filePath == "" {
			filePath = "index.m3u8"
		}

		fileStreamPath := filepath.Join("/hls/live/", streamName, filePath)
		log.Default().Println("Stream file requested:", fileStreamPath)

		return ctx.File(fileStreamPath)
	}
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/live/:live/*", serveStream())
	e.Logger.Fatal(e.Start(":8001"))
}
