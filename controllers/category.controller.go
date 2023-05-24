package controllers

import (
	"kanban-app/database"
	"kanban-app/models"
	"kanban-app/request"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CreateCategory(c *fiber.Ctx) error {
	categoryReq := request.CategoryCreateRequest{}

	if errParse := c.BodyParser(&categoryReq); errParse != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "fail to parsing data",
			"error":   errParse.Error(),
		})
	}

	category := models.Category{}
	category.Category = categoryReq.Category

	if errDb := database.DB.Create(&category).Error; errDb != nil {
		log.Println("category.controller.go => CreateCategory :: ", errDb)
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "catgory created",
		"data":    category,
	})
}

func GetCategoryById(c *fiber.Ctx) error {
	categoryId := c.Params("id")
	category := models.Category{}

	if err := database.DB.First(&category, "id = ?", categoryId).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "category not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "data transmitted",
		"data":    category,
	})
}

func GetAllCategory(c *fiber.Ctx) error {
	return nil
}
