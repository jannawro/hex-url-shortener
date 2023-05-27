package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jannawro/hex-url-shortener/internal/core/services/urlsrv"
	"github.com/jannawro/hex-url-shortener/internal/handlers/urlhdl"
	"github.com/jannawro/hex-url-shortener/internal/repositories/memorykv"
)

func main() {
    urlRepository := memorykv.New()
    urlService := urlsrv.New(urlRepository)
    urlHandler := urlhdl.NewHTTPHanlder(urlService)

    app := fiber.New()
    app.Use(logger.New())

    app.Get("/all", func(c *fiber.Ctx) error {
        return urlHandler.GetAll(c)
    })
    app.Get("/:id", func(c *fiber.Ctx) error {
        return urlHandler.Redirect(c)
    }) 
    app.Post("/", func(c *fiber.Ctx) error {
        return urlHandler.Create(c)
    })

    app.Listen(":3000")
}
