# TODO API на Go + Fiber + PostgreSQL

Простое REST API для управления задачами (TODO-лист) с использованием Go, Fiber и PostgreSQL. Поддерживает создание, просмотр, редактирование и удаление задач.

## 📦 Стек:
- Golang + Fiber
- PostgreSQL (`pgx`)
- REST API

## 🚀 Запуск

1. Создай файл `.env`:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=yourpassword
DB_NAME=yourdbname
```

2. Установите зависимости
```bash
go mod tidy
```

3. Запустите сервер
```bash
make run
```

## 📸 Скриншоты

### Успешный запрос (POST /tasks)
![POST ЗАПРОС](docs/screenshot/postTasksInPostman.png)

### Запись в базе данных
![БАЗА ДАННЫХ](docs/screenshot/postTasksInPostgres.png)

### Успешное получение всех задач (GET /tasks)
![GET ЗАПРОС](docs/screenshot/getAllTasksInPostman.png)

### Успешное изменение задачи (PUT /tasks/:id)
![PUT ЗАПРОС](docs/screenshot/updateTasksInPostman.png)

### Запись в базе данных
![БАЗА ДАННЫХ](docs/screenshot/updateTasksInPostgres.png)

### Успешное удаление задачи (DELETE /tasks/:id)
![DELETE ЗАПРОС](docs/screenshot/deleteTasksInPostman.png)

### Запись в базе данных
![БАЗА ДАННЫХ](docs/screenshot/deleteTasksInPostgres.png)