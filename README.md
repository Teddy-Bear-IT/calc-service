# Calculator Web Service

Это веб-сервис, который вычисляет арифметические выражения. Пользователи могут отправлять POST-запросы с арифметическими выражениями и получать вычисленные результаты.

## Как использовать

1. Склонируйте репозиторий:
   git clone [https://github.com/Teddy-Bear-IT/calc-service.git](https://github.com/Teddy-Bear-IT/calc-service.git)
2. Запустите сервис:

   ```go
   go run ./cmd/calc_service/...
   ```

3. Сервис запустился по адресу `http://localhost:8080/`

## API Endpoint

- URL: `/api/v1/calculate`
- Method: POST
- Request Body

### Запрос

```json
{
  "expression": "Выражение, которое ввёл пользователь"
}
```

### Статусы(Возможные ответы)

1. Успешный - 200 КОД

```json
{
  "result": "Выражение, которое ввёл пользователь"
}
```

2. Ошибка входных данных - 422 КОД

```json
{
  "error": "Expression is not valid"
}
```

3. Ошибка сервера - 500 код

```json
{
  "error": "Internal server error"
}
```

# Примеры запросов к серверу и его реакции

## Успешный запрос(200)

### Входные данные

```
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```

### Выходные данные

```json
{
  "result": 6
}
```

## Ошибка входных данных(422)

### Входные данные

```
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*a"
}'
```

### Выходные данные

```go
{
  "error": "Expression is not valid"
}
```

## Ошибка сервера(500)

### Входные данные

```
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "5/0"
}'
```

### Выходные данные

```go
{
  "error": "Internal server error"
}
```

# Запуск проекта

```
go run .cmd/

```
