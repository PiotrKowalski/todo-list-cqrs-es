package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"todo-list-cqrs-es/internal/domain/todo_list"
)

var (
	database   = "todolist"
	collection = "todolists"
)

type todoListRepoAdapter struct {
	client *mongo.Client
}

func (t todoListRepoAdapter) Save(list *todo_list.TodoList) (*todo_list.TodoList, error) {
	coll := t.client.Database(database).Collection(collection)

	ret, err := coll.InsertOne(context.TODO(), list)
	if err != nil {
		return nil, err
	}

	list.Id = ret.InsertedID.(primitive.ObjectID)

	return list, nil
}

func (t todoListRepoAdapter) Load(id string) (*todo_list.TodoList, error) {
	coll := t.client.Database(database).Collection(collection)

	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", hex}}
	res := coll.FindOne(context.TODO(), filter)
	if err := res.Err(); err != nil {
		return nil, err
	}

	var todoList todo_list.TodoList
	err = res.Decode(&todoList)
	if err != nil {
		return nil, err
	}

	return &todoList, nil
}

func (t todoListRepoAdapter) Remove(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewTodoListRepoAdapter(URI string) *todoListRepoAdapter {
	log.Printf("Connecting to mongodb %v", URI)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	if err != nil {
		panic(err)
	}

	return &todoListRepoAdapter{client: client}
}
