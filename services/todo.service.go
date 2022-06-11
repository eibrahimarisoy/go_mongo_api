package services

import (
	"fmt"

	"github.com/eibrahimarisoy/go_mongo_api/dto"
	"github.com/eibrahimarisoy/go_mongo_api/models"
	"github.com/eibrahimarisoy/go_mongo_api/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DefaultTodoService struct {
	Repo repository.TodoRepository
}

type TodoService interface {
	TodoInsert(todo models.Todo) (*dto.TodoDTO, error)
	GetAllTodo() ([]models.Todo, error)
	DeleteTodo(id primitive.ObjectID) (bool, error)
}

func (t DefaultTodoService) TodoInsert(todo models.Todo) (*dto.TodoDTO, error) {
	var res *dto.TodoDTO

	if len(todo.Title) <= 2 {
		res.Status = false
		return res, nil
	}
	result, err := t.Repo.Insert(todo)

	if err != nil || result == false {
		res.Status = false
		return res, err
	}

	res = &dto.TodoDTO{
		Status: result,
	}
	return res, nil
}

func (t DefaultTodoService) GetAllTodo() ([]models.Todo, error) {
	result, err := t.Repo.GetAll()
	fmt.Println(result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t DefaultTodoService) DeleteTodo(id primitive.ObjectID) (bool, error) {
	result, err := t.Repo.Delete(id)

	if err != nil || result == false {
		return false, err
	}

	return true, nil
}

func NewTodoService(repo repository.TodoRepository) DefaultTodoService {
	return DefaultTodoService{
		Repo: repo,
	}
}
