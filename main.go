package main

import (
	"github.com/zeddy-go/zeddy/app"
	"github.com/zeddy-go/zeddy/http/ginx"
	"log/slog"
	"template/module/user"
)

func main() {
	app.Use(
		ginx.NewModule(),
		user.NewModule(),
	)
	err := app.StartAndWait()
	if err != nil {
		slog.Warn("An error occurred", "error", err)
	}
}
