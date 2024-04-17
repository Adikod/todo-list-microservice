# Todo List Microservice

Этот проект представляет собой микросервис для управления списком дел (todo list) с использованием RESTful API. Микросервис реализован на языке Go с использованием фреймворка Gin и базы данных MongoDB.

## Установка

Для установки необходимо склонировать репозиторий:

```bash 
git clone https://github.com/Adikod/todo-list-microservice.git
```

## Запуск

Для запуска есть несколько вариантов:
## Make
```bash
make build
make run
```
## Docker
```bash
docker build -t todo-list-microservice . 
docker run -d -p 8080:8080 todo-list-microservice
```
