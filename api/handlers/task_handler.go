package handlers

import (
	"github.com/Adikod/todo-list-microservice/models"
	"github.com/Adikod/todo-list-microservice/services"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

// TaskHandler представляет обработчики HTTP-запросов для задач
type TaskHandler struct {
	service *services.TaskService // Сервис для работы с задачами
}

// NewTaskHandler создает новый экземпляр TaskHandler
func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service, // Инициализация сервиса
	}
}

// CreateTaskHandler обрабатывает запрос на создание новой задачи
// @Summary Создать новую задачу
// @Description Создает новую задачу на основе переданных данных
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body models.Task true "Данные для создания задачи"
// @Success 201 {string} string "task created successfully"
// @Failure 400 {object} ErrorResponse "Ошибка при разборе JSON"
// @Failure 500 {object} ErrorResponse "Ошибка при создании задачи"
// @Router /api/todo-list/tasks [post]
func (h *TaskHandler) CreateTaskHandler(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(403, gin.H{"error": err.Error()}) // Возвращаем ошибку 403, если не удалось разобрать JSON
		return
	}

	err := h.service.CreateTask(&task) // Создаем задачу через сервис
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()}) // Возвращаем ошибку 500, если не удалось создать задачу
		return
	}

	c.JSON(201, "task created successfully") // Возвращаем успешный статус и сообщение об успешном создании задачи
}

// UpdateTaskHandler обрабатывает запрос на обновление существующей задачи
// @Summary Обновить задачу
// @Description Обновляет существующую задачу на основе переданных данных
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "ID задачи для обновления"
// @Param task body models.Task true "Данные для обновления задачи"
// @Success 204 {string} string "task updated successfully"
// @Failure 400 {object} ErrorResponse "Ошибка при разборе JSON"
// @Failure 404 {object} ErrorResponse "Задача не найдена"
// @Router /api/todo-list/tasks/{id} [put]
func (h *TaskHandler) UpdateTaskHandler(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()}) // Возвращаем ошибку 400, если не удалось разобрать JSON
		return
	}

	err := h.service.UpdateTask(id, &task) // Обновляем задачу через сервис
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()}) // Возвращаем ошибку 404, если не удалось обновить задачу
		return
	}

	c.Status(204) // Возвращаем успешный статус без содержимого
}

// DeleteTaskHandler обрабатывает запрос на удаление задачи
// @Summary Удалить задачу
// @Description Удаляет задачу по ее ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "ID задачи для удаления"
// @Success 204 {string} string "task deleted successfully"
// @Failure 404 {object} ErrorResponse "Задача не найдена"
// @Router /api/todo-list/tasks/{id} [delete]
func (h *TaskHandler) DeleteTaskHandler(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteTask(id) // Удаляем задачу через сервис
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()}) // Возвращаем ошибку 404, если не удалось найти задачу для удаления
		return
	}

	c.Status(204) // Возвращаем успешный статус без содержимого
}

// MarkTaskAsDoneHandler обрабатывает запрос на пометку задачи как выполненной
// @Summary Пометить задачу как выполненную
// @Description Помечает задачу как выполненную по ее ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "ID задачи для пометки как выполненной"
// @Success 204 {string} string "task marked as done successfully"
// @Failure 404 {object} ErrorResponse "Задача не найдена"
// @Router /api/todo-list/tasks/{id}/done [post]
func (h *TaskHandler) MarkTaskAsDoneHandler(c *gin.Context) {
	id := c.Param("id")
	err := h.service.MarkTaskAsDone(id) // Помечаем задачу как выполненную через сервис
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()}) // Возвращаем ошибку 404, если не удалось найти задачу для пометки как выполненную
		return
	}

	c.Status(204) // Возвращаем успешный статус без содержимого
}

// GetTasksByStatusHandler обрабатывает запрос на получение списка задач по статусу
// @Summary Получить список задач по статусу
// @Description Возвращает список задач по указанному статусу
// @Tags tasks
// @Accept json
// @Produce json
// @Param status query string false "Статус задачи ('active' или 'done'), по умолчанию 'active'"
// @Success 200 {array} models.Task "Список задач"
// @Failure 500 {object} ErrorResponse "Ошибка при получении списка задач"
// @Router /api/todo-list/tasks [get]
func (h *TaskHandler) GetTasksByStatusHandler(c *gin.Context) {
	status := c.DefaultQuery("status", "active")
	tasks, err := h.service.GetTasksByStatus(status) // Получаем список задач по указанному статусу через сервис
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()}) // Возвращаем ошибку 500, если произошла ошибка при получении задач
		return
	}

	c.JSON(200, tasks) // Возвращаем список задач с успешным статусом
}
