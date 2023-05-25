package controllers

import (
	"kanban-app/database"
	"kanban-app/models/entity"
	"kanban-app/models/request"
	"kanban-app/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func LoginController(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)

	// PARSE REQUEST BODY
	if errParse := ctx.BodyParser(&loginRequest); errParse != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "fail to parsing data",
			"error":   errParse.Error(),
		})
	}
	// VALIDATION REQUEST DATA
	validate := validator.New()
	if errValidate := validate.Struct(loginRequest); errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "some data is not valid",
			"error":   errValidate.Error(),
		})
	}

	// CHECK VALIDATION USER DATA
	var user entity.User
	if err := database.DB.First(&user, "email = ?", loginRequest.Email).Error; err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credentials",
		})
	}

	// CHECK VALIDATION PASSWORD DATA
	isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password)
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credentials",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": "token",
	})
}
