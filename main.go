package main

import (
	"github.com/eibrahimarisoy/go_mongo_api/app"
	"github.com/eibrahimarisoy/go_mongo_api/configs"
	"github.com/eibrahimarisoy/go_mongo_api/repository"
	"github.com/eibrahimarisoy/go_mongo_api/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	appRoute := fiber.New()
	configs.ConnectDB()
	dbClient := configs.GetCollection(configs.DB, "todos")

	TodoRepositoryDB := repository.NewTodoRepositoryDB(dbClient)

	td := app.TodoHandler{
		Service: services.NewTodoService(TodoRepositoryDB),
	}
	api := appRoute.Group("/api")
	api.Post("/todos", td.CreateTodo)
	api.Get("/todos", td.GetAllTodo)
	api.Delete("/todos/:id", td.DeleteTodo)

	appRoute.Listen(":8000")
}
