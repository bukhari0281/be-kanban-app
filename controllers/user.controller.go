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
	registerReq := request.UserRegisterRequest{}

	// PARSE REQUEST BODY
	if errParse := c.BodyParser(&registerReq); errParse != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "fail to parsing data",
			"error":   errParse.Error(),
		})
	}

	// VALIDATION REQUEST DATA

	validate := validator.New()
	if errValidate := validate.Struct(&registerReq); errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "some data is not valid",
			"error":   errValidate.Error(),
		})
	}

	register := entity.User{}
	register.Name = registerReq.Name
	register.Email = registerReq.Email
	register.Password = registerReq.Password

	if errDb := database.DB.Create(&register).Error; errDb != nil {
		log.Println("todo.controller.go => CreateTodo :: ", errDb)
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "user registered successfully",
		"data":    register,
	})
}
