package controllers

import (
	"kanban-app/database"
	"kanban-app/models/entity"
	"kanban-app/models/request"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	user := request.UserRegisterRequest{}

	// PARSE REQUEST BODY
	if errParse := c.BodyParser(&user); errParse != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "fail to parsing data",
			"error":   errParse.Error(),
		})
	}

	// VALIDATION REQUEST DATA

	validate := validator.New()
	if errValidate := validate.Struct(&user); errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "some data is not valid",
			"error":   errValidate.Error(),
		})
	}

	newUser := entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	if errDb := database.DB.Create(&newUser).Error; errDb != nil {
		log.Println("todo.controller.go => CreateTodo :: ", errDb)
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "user registered successfully",
		"data":    newUser,
	})
}
