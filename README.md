# Тестовое задание для Недорезов. [![Test](https://github.com/lirprocs/test_nedorezov/actions/workflows/test.yaml/badge.svg)](https://github.com/lirprocs/test_nedorezov/actions/workflows/test.yaml)
## Реализованы следующие эндпоинты:
1. POST /accounts - создание нового аккаунта.
2. POST /accounts/{id}/deposit - пополнение баланса.
3. POST /accounts/{id}/withdraw - снятие средств.
4. GET  /accounts/{id}/balance - проверка баланса.

## Установка
1. Клонируйте репозиторий
```bash
git clone https://github.com/lirprocs/test_nedorezov
```
2. Перейдите в директорию проекта:
```bash
cd test_nedorezov
```
3. Установите зависимости:
```bash
go mod tidy
```

## Запуск сервера
1. Перейдите в директорию проекта (Не нужно, еслу уже находитесь в ней):
```bash
cd test_nedorezov
```
2. Запустите сервер:
```bash
go run main.go
```

## Тестирование
1. Для запуска тестов выполните:
``` bash
go test ./...
```
