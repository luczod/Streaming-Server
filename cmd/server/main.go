package main

import (
	"authServer/config/db"
	"authServer/config/env"
	"authServer/internal/handler"
	"authServer/internal/repository"
	"authServer/internal/service"
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sethvargo/go-envconfig"
)

func main() {
	ctx := context.Background()

	var envConfig env.EnvConfig

	if err := envconfig.Process(ctx, &envConfig); err != nil {
		log.Fatal(err)
	}

	db, err := db.OpenConn(envConfig)
	if err != nil {
		log.Fatal("error connect databbase")
	}

	// init app
	keysRepository := repository.NewKeysReposiroy(db)
	keysService := service.NewKeysService(keysRepository)
	keysHandler := handler.NewHandler(keysService)

	log.Default().Println("Routing...")
	e := echo.New()
	e.POST("/auth", keysHandler.AuthStreamingKey)
	e.GET("/healthcheck", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "WORKING")
	})

	e.Logger.Fatal(e.Start(":8000"))
}

// app=live&flashver=FMLE/3.0%20(compatible%3B%20FMSc/1.0)
// &swfurl=rtmp://localhost:1935/live
// &tcurl=rtmp://localhost:1935/live
// &pageurl=&addr=172.23.0.1
// &clientid=1&call=publish&name=yourstreamkey
// &type=live

// app=live&flashver=FMLE/3.0%20(compatible%3B%20FMSc/1.0)
// &swfurl=rtmp://localhost:1935/live
// &tcurl=rtmp://localhost:1935/live
// &pageurl=
// &addr=172.24.0.1
// &clientid=1
// &call=publish
// &name=yourstreamkey
// &type=live
