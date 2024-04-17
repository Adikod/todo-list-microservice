package routes

import (
	"github.com/Adikod/todo-list-microservice/api/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Adikod/todo-list-microservice/docs"
)

// SetupTaskRoutes настраивает маршруты для задач
func SetupTaskRoutes(router *gin.Engine, taskHandler *handlers.TaskHandler) {
	api := router.Group("/api/todo-list/tasks") // Группа маршрутов для задач

	// Устанавливаем обработчики для различных методов HTTP
	{
		api.POST("", taskHandler.CreateTaskHandler)             // Обработчик для создания задачи (POST)
		api.PUT("/:id", taskHandler.UpdateTaskHandler)          // Обработчик для обновления задачи (PUT)
		api.DELETE("/:id", taskHandler.DeleteTaskHandler)       // Обработчик для удаления задачи (DELETE)
		api.PUT("/:id/done", taskHandler.MarkTaskAsDoneHandler) // Обработчик для пометки задачи как выполненной (PUT)
		api.GET("", taskHandler.GetTasksByStatusHandler)        // Обработчик для получения списка задач (GET)
	}

	// Маршрут для обслуживания документации Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
