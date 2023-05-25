package middleware

import "github.com/gofiber/fiber/v2"

func Middleware(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")
	if token != "secret" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthrized",
		})
	}
	return ctx.Next()
}
