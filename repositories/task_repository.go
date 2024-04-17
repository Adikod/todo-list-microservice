package repositories

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"

	"github.com/Adikod/todo-list-microservice/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TaskRepository представляет интерфейс для работы с задачами в базе данных
type TaskRepository interface {
	Create(task *models.Task) error
	Update(id string, task *models.Task) error
	Delete(id string) error
	FindByID(id string) (*models.Task, error)
	FindByTitleAndActiveAt(title string, activeAt time.Time) (*models.Task, error)
	FindAll(status string) ([]models.Task, error)
}

// MongoDBTaskRepository представляет реализацию интерфейса TaskRepository для MongoDB
type MongoDBTaskRepository struct {
	collection *mongo.Collection
}

// NewMongoDBTaskRepository создает новый экземпляр MongoDBTaskRepository
func NewMongoDBTaskRepository(collection *mongo.Collection) *MongoDBTaskRepository {
	return &MongoDBTaskRepository{
		collection: collection,
	}
}

// Create создает новую задачу в MongoDB
func (r *MongoDBTaskRepository) Create(task *models.Task) error {
	_, err := r.collection.InsertOne(context.Background(), task)
	if err != nil {
		return err
	}
	return nil
}

// Update обновляет задачу в MongoDB
func (r *MongoDBTaskRepository) Update(id string, task *models.Task) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": task}

	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete удаляет задачу из MongoDB по ее ID
func (r *MongoDBTaskRepository) Delete(id string) error {
	filter := bson.M{"_id": id}

	_, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

// FindByID находит задачу в MongoDB по ее ID
func (r *MongoDBTaskRepository) FindByID(id string) (*models.Task, error) {
	var task models.Task
	filter := bson.M{"_id": id}

	err := r.collection.FindOne(context.Background(), filter).Decode(&task)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("задача с ID %s не найдена", id)
		}
		return nil, err
	}

	return &task, nil
}

// FindByTitleAndActiveAt находит задачу в MongoDB по ее заголовку и дате активности
func (r *MongoDBTaskRepository) FindByTitleAndActiveAt(title string, activeAt time.Time) (*models.Task, error) {
	var task models.Task
	filter := bson.M{"title": title, "activeAt": activeAt}

	err := r.collection.FindOne(context.Background(), filter).Decode(&task)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &task, nil
}

// FindAll возвращает список задач в MongoDB по заданному статусу
func (r *MongoDBTaskRepository) FindAll(status string) ([]models.Task, error) {
	var tasks []models.Task

	filter := bson.M{}
	if status == "done" {
		filter["done"] = true
	} else {
		filter["done"] = false
	}

	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func DBSet() *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/test"))
	if err != nil {
		log.Fatal()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal()
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("failed to connect to mongodb")
	}

	return client
}

var Client *mongo.Client = DBSet()

func TaskData(client *mongo.Client, CollectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("test").Collection(CollectionName)
	return collection
}
