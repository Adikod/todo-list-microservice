package services

import (
	"errors"
	"github.com/Adikod/todo-list-microservice/models"
	"github.com/Adikod/todo-list-microservice/repositories"
	"time"
)

// TaskService представляет сервис для работы с задачами
type TaskService struct {
	taskRepository repositories.TaskRepository
}

// NewTaskService создает новый экземпляр TaskService
func NewTaskService(taskRepository repositories.TaskRepository) *TaskService {
	return &TaskService{
		taskRepository: taskRepository,
	}
}

func (s *TaskService) CreateTask(task *models.Task) error {
	// Проверяем, что все поля обязательны для заполнения
	if task.Title == "" || task.ActiveAt.IsZero() {
		return errors.New("все поля обязательны для заполнения")
	}

	// Валидация на длину заголовка
	if len(task.Title) > 200 {
		return errors.New("длина заголовка не должна превышать 200 символов")
	}

	// Валидация даты
	if task.ActiveAt.Before(time.Now()) {
		return errors.New("неверная дата")
	}

	// Проверка уникальности по полям title и activeAt
	existingTask, err := s.taskRepository.FindByTitleAndActiveAt(task.Title, task.ActiveAt)
	if err != nil {
		return err
	}

	if existingTask != nil {
		return errors.New("задача с таким заголовком и датой уже существует")
	}

	// Устанавливаем дату создания
	task.CreatedAt = time.Now()

	// Создаем задачу в базе данных
	err = s.taskRepository.Create(task)
	if err != nil {
		return err
	}

	return nil
}

// UpdateTask обновляет существующую задачу
func (s *TaskService) UpdateTask(id string, task *models.Task) error {
	// Проверяем, что все поля обязательны для заполнения
	if task.Title == "" || task.ActiveAt.IsZero() {
		return errors.New("все поля обязательны для заполнения")
	}

	// Валидация на длину заголовка
	if len(task.Title) > 200 {
		return errors.New("длина заголовка не должна превышать 200 символов")
	}

	// Валидация даты
	if task.ActiveAt.Before(time.Now()) {
		return errors.New("неверная дата")
	}

	// Проверяем, существует ли задача с данным ID
	_, err := s.taskRepository.FindByID(id)
	if err != nil {
		return err
	}

	// Обновляем задачу в базе данных
	err = s.taskRepository.Update(id, task)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTask удаляет задачу по ее ID
func (s *TaskService) DeleteTask(id string) error {
	// Проверяем, существует ли задача с данным ID
	_, err := s.taskRepository.FindByID(id)
	if err != nil {
		return err
	}

	// Удаляем задачу из базы данных
	err = s.taskRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

// MarkTaskAsDone помечает задачу как выполненную
func (s *TaskService) MarkTaskAsDone(id string) error {
	// Проверяем, существует ли задача с данным ID
	task, err := s.taskRepository.FindByID(id)
	if err != nil {
		return err
	}

	// Помечаем задачу как выполненную
	task.Done = true

	// Обновляем задачу в базе данных
	err = s.taskRepository.Update(id, task)
	if err != nil {
		return err
	}

	return nil
}

// GetTasksByStatus возвращает список задач по заданному статусу
func (s *TaskService) GetTasksByStatus(status string) ([]models.Task, error) {
	// Проверяем, что статус является допустимым
	if status != "active" && status != "done" {
		return nil, errors.New("недопустимый статус")
	}

	// Получаем список задач из репозитория
	tasks, err := s.taskRepository.FindAll(status)
	if err != nil {
		return nil, err
	}

	// Проверяем день недели и добавляем текст "ВЫХОДНОЙ" при необходимости
	for i := range tasks {
		if isWeekend(tasks[i].ActiveAt) {
			tasks[i].Title = "ВЫХОДНОЙ - " + tasks[i].Title
		}
	}

	return tasks, nil
}

// isWeekend возвращает true, если день является выходным (суббота или воскресенье)
func isWeekend(t time.Time) bool {
	return t.Weekday() == time.Saturday || t.Weekday() == time.Sunday
}
