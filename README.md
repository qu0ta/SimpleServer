# SimpleServer

**SimpleServer** — это минималистичный HTTP-сервер, написанный на Go. Этот проект создан для демонстрации простого сервера, который можно легко настроить и расширить.

## Возможности

- Легкий и понятный код.
- Базовая реализация HTTP-сервера.
- Простота настройки и расширения.

## Требования

- Go версии 1.16 или выше.

## Установка

1. Склонируйте репозиторий:

   ```bash
   git clone https://github.com/qu0ta/SimpleServer.git
   cd SimpleServer
   ```
2. Скомпилируйте и запустите сервер:

```bash
go run main.go
```

(Опционально) Создайте исполняемый файл:

```bash
go build -o simpleserver
./simpleserver
```
## Использование
После запуска сервер будет слушать входящие запросы на порту по умолчанию (например, 8080). Вы можете открыть ваш браузер и перейти по адресу:

```
http://localhost:8080
```
## Конфигурация
Вы можете изменить конфигурацию сервера, редактируя код в файле main.go. Например:

Изменить порт сервера.
Добавить обработчики для новых маршрутов.

## Документация
Для просмотра документации используйте:
```bash
godoc -http "localhost:8000"
```
И перейдите по этому адресу в поисковике