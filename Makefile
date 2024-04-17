# Компиляция и запуск приложения
run:
	go run main.go

# Компиляция приложения
build:
	go build -o todo-list-microservice main.go

# Очистка скомпилированных файлов
clean:
	del todo-list-microservice
