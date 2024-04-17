package main

import (
	"log"

	"github.com/Adikod/todo-list-microservice/api/handlers"
	"github.com/Adikod/todo-list-microservice/api/routes"
	"github.com/Adikod/todo-list-microservice/repositories"
	"github.com/Adikod/todo-list-microservice/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// Создание экземпляра конфигурации
	//cfg := config.NewConfig()

	// Создание экземпляра MongoDB клиента
	mongoClient := repositories.DBSet()

	// Получение коллекции из MongoDB
	collection := repositories.TaskData(mongoClient, "tasks")

	// Создание экземпляра TaskRepository
	taskRepo := repositories.NewMongoDBTaskRepository(collection)

	// Создание экземпляра TaskService
	taskService := services.NewTaskService(taskRepo)

	// Создание экземпляра TaskHandler
	taskHandler := handlers.NewTaskHandler(taskService)

	// Создание роутера Gin
	router := gin.Default()

	// Настройка маршрутов
	routes.SetupTaskRoutes(router, taskHandler)

	// Запуск сервера
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
