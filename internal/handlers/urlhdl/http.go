package urlhdl

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/jannawro/hex-url-shortener/internal/core/ports"
)

type HTTPHandler struct {
    urlService ports.UrlService
}

func NewHTTPHanlder(urlService ports.UrlService) *HTTPHandler {
    return &HTTPHandler{
        urlService: urlService,
    }
}

func(hdl *HTTPHandler) Redirect(c *fiber.Ctx) error {
    url, err := hdl.urlService.Get(c.BaseURL() + "/" + c.Params("id"))
    if err != nil {
        return err
    }

    return c.Redirect(url.Original, 301)
}

func(hdl *HTTPHandler) Create(c *fiber.Ctx) error {
    payload := struct {
        Url string `json:"url"` 
    }{}

    if err := c.BodyParser(&payload); err != nil {
        return err
    }

    url, err := hdl.urlService.Create(c.BaseURL(), payload.Url)
    if err != nil {
        return err
    }

    return c.SendString("Your new shorter URL: " + url.Shorthand)
}

func(hdl *HTTPHandler) GetAll(c *fiber.Ctx) error {
    urls, err := hdl.urlService.GetAll()
    if err != nil {
        return err
    }

    return c.JSON(urls)
}
