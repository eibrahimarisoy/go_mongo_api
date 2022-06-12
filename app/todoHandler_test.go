package app

import (
	"net/http/httptest"
	"testing"

	services "github.com/eibrahimarisoy/go_mongo_api/mocks/service"
	"github.com/eibrahimarisoy/go_mongo_api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var td TodoHandler
var mockService *services.MockTodoService

func setUp(t *testing.T) func() {
	ctrl := gomock.NewController(t)

	mockService = services.NewMockTodoService(ctrl)

	td = TodoHandler{mockService}

	return func() {
		defer ctrl.Finish()
	}
}

func TestTodoHandler_TodoGetAll(t *testing.T) {
	trd := setUp(t)
	defer trd()

	router := fiber.New()

	router.Get("api/todos", td.GetAllTodo)

	var FakeDataForHandler = []models.Todo{
		{primitive.NewObjectID(), "Todo 1", "Todo 1 description"},
		{primitive.NewObjectID(), "Todo 2", "Todo 2 description"},
		{primitive.NewObjectID(), "Todo 3", "Todo 3 description"},
	}

	mockService.EXPECT().GetAllTodo().Return(FakeDataForHandler, nil)

	req := httptest.NewRequest("GET", "/api/todos", nil)
	res, _ := router.Test(req, 1)

	assert.Equal(t, 200, res.StatusCode)
}
