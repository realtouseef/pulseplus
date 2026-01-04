package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/realtouseef/pulseplus/helpers"
)

type PingRequest struct {
	Url string `json:"url"`
}

func Ping(c *fiber.Ctx) error {
	payload := new(PingRequest)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	response, err := http.Get(payload.Url)
	helpers.HandleError(err)
	defer response.Body.Close()

	return c.Status(fiber.StatusOK).SendString(fmt.Sprintf("status code for %v is %v", payload.Url, response.StatusCode))
}
