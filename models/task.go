package models

import "time"

// Task представляет задачу в списке дел
type Task struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`              // ID задачи
	Title     string    `json:"title" bson:"title"`                             // Заголовок задачи
	ActiveAt  time.Time `json:"activeAt" bson:"activeAt"`                       // Дата активации задачи
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"` // Дата создания задачи
	Done      bool      `json:"done" bson:"done"`                               // Статус выполнения задачи
}
