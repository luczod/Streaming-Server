// Package handler is to foo
package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"authServer/internal/model"
	"authServer/internal/service"

	"github.com/labstack/echo/v4"
)

type IKeysHandler interface {
	AuthStreamingKey(ctx echo.Context) error
}

type keysHandler struct {
	KeysService service.IKeyService
}

func NewHandler(s service.IKeyService) IKeysHandler {
	return &keysHandler{
		KeysService: s,
	}
}

func (kh *keysHandler) AuthStreamingKey(ctx echo.Context) error {
	body := ctx.Request().Body
	defer body.Close()

	fields, _ := io.ReadAll(body)
	log.Default().Println("Auth...", string(fields))
	passedKeyValue := getStreamKey(fields)

	keys, err := kh.KeysService.AuthStreamingKey(passedKeyValue.Name, passedKeyValue.Key)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "problem with streaming key")
	}

	if keys.Key == "" {
		log.Default().Println("Forbidden User!")
		return ctx.String(http.StatusForbidden, "Forbidden")
	}

	log.Default().Println("User authenticated!")

	newStreamURL := fmt.Sprintf("rtmp://127.0.0.1:1935/hls-live/%s", keys.Name)
	log.Default().Println("Redirecting to:", newStreamURL)

	return ctx.Redirect(http.StatusFound, newStreamURL)
}

func getStreamKey(s []byte) model.Keys {
	var authValues model.Keys

	pairs := strings.Split(string(s), "&")

	fmt.Println(pairs)

	for _, pair := range pairs {
		splitPair := strings.Split(pair, "=")
		key := splitPair[0]
		value := splitPair[1]

		if key == "name" {
			allPassedValues := strings.Split(value, "_")
			authValues.Name = allPassedValues[0]
			authValues.Key = allPassedValues[1]
		}
	}

	return authValues
}
