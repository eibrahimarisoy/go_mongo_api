package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/eibrahimarisoy/go_mongo_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockgen -destination=../mocks/repository/mockTodoService.go -package=repository github.com/eibrahimarisoy/go_mongo_api/repository TodoRepository

type TodoRepositoryDB struct {
	TodoCollection *mongo.Collection
}

type TodoRepository interface {
	Insert(todo models.Todo) (bool, error)
	GetAll() ([]models.Todo, error)
	Delete(id primitive.ObjectID) (bool, error)
}

func (t *TodoRepositoryDB) Insert(todo models.Todo) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	todo.Id = primitive.NewObjectID()
	result, err := t.TodoCollection.InsertOne(ctx, todo)

	if err != nil || result.InsertedID == nil {
		return false, err
	}

	return true, nil

}

func (t TodoRepositoryDB) GetAll() ([]models.Todo, error) {
	var todo models.Todo
	var todos []models.Todo

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := t.TodoCollection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	for result.Next(ctx) {
		if err := result.Decode(&todo); err != nil {
			log.Fatal(err)
			fmt.Println(todo)
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (t TodoRepositoryDB) Delete(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := t.TodoCollection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		return false, err
	}

	return result.DeletedCount > 0, nil

}

func NewTodoRepositoryDB(dbClient *mongo.Collection) *TodoRepositoryDB {
	return &TodoRepositoryDB{TodoCollection: dbClient}
}
