package routes

import (
	"kanban-app/controllers"
	"kanban-app/middleware"

	// "github.com/gofiber/fiber/middleware"
	"github.com/gofiber/fiber/v2"
)

func userRoute(app *fiber.App) {
	register := app.Group("/api/v1")
	user := register.Group("/user")

	user.Post("/register", controllers.CreateUser)
	user.Post("/login", controllers.LoginController)

}

func v1Route(app *fiber.App) {
	v1 := app.Group("/api/v1")

	todo := v1.Group("/todos")

	todo.Post("/", controllers.CreateTodo)

	todo.Get("/", middleware.Middleware, controllers.GetAllTodo)

	todo.Get("/:id", controllers.GetTodoById)

	todo.Patch("/:id", controllers.UpdateTodoById)

	todo.Delete("/:id", controllers.DeleteTodoById)

	category := v1.Group("/category")

	category.Post("/", controllers.CreateCategory)

	category.Get("/", controllers.GetAllCategory)

	category.Get("/:id", controllers.GetCategoryById)
}
