package services

import (
	"testing"

	"github.com/eibrahimarisoy/go_mongo_api/mocks/repository"
	"github.com/eibrahimarisoy/go_mongo_api/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mockRepo *repository.MockTodoRepository
var service TodoService

var FakeData = []models.Todo{
	{primitive.NewObjectID(), "Todo 1", "Todo 1 description"},
	{primitive.NewObjectID(), "Todo 2", "Todo 2 description"},
	{primitive.NewObjectID(), "Todo 3", "Todo 3 description"},
}

func setUp(t *testing.T) func() {
	ct := gomock.NewController(t)
	defer ct.Finish()

	mockRepo = repository.NewMockTodoRepository(ct)
	service = NewTodoService(mockRepo)

	return func() {
		service = nil
		defer ct.Finish()
	}
}

func TestDefaultTodoService_TodoGetAll(t *testing.T) {
	td := setUp(t)
	defer td()

	mockRepo.EXPECT().GetAll().Return(FakeData, nil)
	result, err := service.GetAllTodo()

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.NotEmpty(t, result)
}
