package config

import (
	"os"
)

// Config представляет конфигурацию приложения
type Config struct {
	Port        string // Порт сервера
	MongoURI    string // URI для подключения к MongoDB
	MongoDBName string // Имя базы данных MongoDB
}

// NewConfig создает новый экземпляр конфигурации
func NewConfig() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),                           // Получаем порт сервера из переменной среды или используем значение по умолчанию
		MongoURI:    getEnv("MONGO_URI", "mongodb://localhost:27017"), // Получаем URI для подключения к MongoDB из переменной среды или используем значение по умолчанию
		MongoDBName: getEnv("MONGO_DB_NAME", "todo"),                  // Получаем имя базы данных MongoDB из переменной среды или используем значение по умолчанию
	}
}

// getEnv возвращает значение переменной среды, если оно существует, иначе использует значение по умолчанию
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue // Если переменная среды не существует, возвращаем значение по умолчанию
	}
	return value // Возвращаем значение переменной среды
}
