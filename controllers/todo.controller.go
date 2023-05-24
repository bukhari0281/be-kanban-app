package controllers

import (
	"kanban-app/database"
	"kanban-app/models"
	"kanban-app/request"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateTodo(c *fiber.Ctx) error {
	todoReq := request.TodoCreateRequest{}

	// PARSE REQUEST BODY
	if errParse := c.BodyParser(&todoReq); errParse != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "fail to parsing data",
			"error":   errParse.Error(),
		})
	}

	// VALIDATION REQUEST DATA

	validate := validator.New()
	if errValidate := validate.Struct(&todoReq); errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "some data is not valid",
			"error":   errValidate.Error(),
		})
	}

	todo := models.Todo{}
	todo.Name = todoReq.Name
	// todo.Category = todoReq.Category
	todo.IsComplete = todoReq.IsComplete
	if todoReq.Note != "" {
		todo.Note = &todoReq.Note
	}

	if errDb := database.DB.Create(&todo).Error; errDb != nil {
		log.Println("todo.controller.go => CreateTodo :: ", errDb)
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "todo created",
		"data":    todo,
	})
}

func GetTodoById(c *fiber.Ctx) error {
	todoId := c.Params("id")
	todo := models.Todo{}

	if err := database.DB.Preload("Category").First(&todo, "id = ?", todoId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "todo not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "data transmitted",
		"data":    todo,
	})
}

func GetAllTodo(c *fiber.Ctx) error {
	todos := []models.Todo{}

	if err := database.DB.Preload("Category").Find(&todos).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "data transmited",
		"data":    todos,
	})
}

func UpdateTodoById(c *fiber.Ctx) error {
	todoReq := request.TodoUpdateRequest{}

	// PARSE REQUEST BODY
	if errParse := c.BodyParser(&todoReq); errParse != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "fail to parsing data",
			"error":   errParse.Error(),
		})
	}

	// VALIDATION REQUEST DATA

	validate := validator.New()
	if errValidate := validate.Struct(&todoReq); errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "some data is not vali",
			"error":   errValidate.Error(),
		})
	}

	todoId := c.Params("id")
	todo := models.Todo{}

	if err := database.DB.First(&todo, "id = ?", todoId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "todo not found",
		})
	}

	todo.Name = todoReq.Name
	todo.Note = &todoReq.Note
	todo.IsComplete = todoReq.IsComplete

	if errSave := database.DB.Save(&todo).Error; errSave != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "todo updated",
		"data":    todo,
	})
}

func DeleteTodoById(c *fiber.Ctx) error {

	todoId := c.Params("id")
	todo := models.Todo{}

	if err := database.DB.First(&todo, "id = ?", todoId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "todo not found",
		})
	}

	if errDelete := database.DB.Delete(&todo).Error; errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "todo deleted",
	})
}
