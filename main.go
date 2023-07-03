package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

// Embed a single file
//
//go:embed index.html
var f embed.FS

// Embed a directory
//
//go:embed assets/*
var embedDirStatic embed.FS

func main() {
	app := fiber.New()

	// Access file "image.png" under `assets/` directory via URL: `http://<server>/assets/image.png`.
	// Without `PathPrefix`, you have to access it via URL:
	// `http://<server>/assets/assets/image.png`.
	app.Use("/assets", filesystem.New(filesystem.Config{
		Root:       http.FS(embedDirStatic),
		PathPrefix: "assets",
		Browse:     true,
	}))

	app.Use("*", func(ctx *fiber.Ctx) error {
		return filesystem.SendFile(ctx, http.FS(f), "index.html")
	})

	log.Fatal(app.Listen(":3023"))
}
