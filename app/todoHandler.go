package app

import (
	"net/http"

	"github.com/eibrahimarisoy/go_mongo_api/models"
	"github.com/eibrahimarisoy/go_mongo_api/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoHandler struct {
	Service services.TodoService
}

func (h TodoHandler) CreateTodo(c *fiber.Ctx) error {
	var todo models.Todo

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	result, err := h.Service.TodoInsert(todo)

	if err != nil || result.Status == false {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(result)
}

func (h TodoHandler) GetAllTodo(c *fiber.Ctx) error {
	result, err := h.Service.GetAllTodo()

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(result)
}

func (h TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	_id, _ := primitive.ObjectIDFromHex(id)

	result, err := h.Service.DeleteTodo(_id)

	if err != nil || result == false {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"State": result})
}
