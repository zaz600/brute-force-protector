[![Go](https://github.com/zaz600/brute-force-protector/actions/workflows/go-pr-check.yml/badge.svg)](https://github.com/zaz600/brute-force-protector/actions/workflows/go-pr-check.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/zaz600/brute-force-protector)](https://goreportcard.com/report/github.com/zaz600/brute-force-protector)



# Проектная работа - Bruteforce Protector Service
- Курс [Golang Developer. Professional](https://otus.ru/lessons/golang-professional/)
- [ТЗ](docs/specification.md)

## Сборка и тестирование
- `make run` - сборка и запуск докер образа с сервером. Запуск осуществляется в фоне.
- `make run-log` - сборка и запуск докер образа с сервером. Не отсоединяется от консоли.
- `make stop` - остановка докер образа
- `make build` - сборка сервера и клиента
- `make build-server` - сборка сервера
- `make build-cli` - сборка клиента
- `make lint` - запуск линтера
- `make test` - запуск юнит-тестов
- `make release` - сборка клиента, сервера, запуск тестов и линтера
- `make generate` - генерация protobuf/grpc
- `make itest` - запуск интеграционных тестов в докере.
- `make itest-stop` - удаление контейнеров, используемых для запуска интеграционных тестов.

Проверить работу сервиса можно запуском интеграционных тестов:
- Убедиться, что у вас установлен `docker` и `docker-compose`
- Убедиться, что у вас MacOS/Linux
- Запустить `make itest`

## bp-cli
`bp-cli` - CLI для Bruteforce Protector. 

Адрес сервера (опционально) указывается при помощи ключа `-server 127.0.0.1:50051`, 
который можно добавить к командам.

### Список команд

- `bp-cli help` - справка по использованию.
- `bp-cli blacklist add <network>` - добавление подсети в черный список.
- `bp-cli blacklist remove <network>` - удаление подсети из черного списка.
- `bp-cli blacklist show` - вывод содержимого черного списка.
- `bp-cli whitelist add <network>` - добавление подсети в белый список.
- `bp-cli whitelist remove <network>` - удаление подсети из белого списка.
- `bp-cli whitelist show` - вывод содержимого белого списка.
- `bp-cli reset login <login>` - сброс лимита для логина.
- `bp-cli reset ip <ip>` - сброс лимита для IP.

## Roadmap

- [x] IP access list  
- [x] sliding window rate limiter  
- [x] bruteforce protector service methods  
- [x] persistent storage (MongoDB/Redis)  
- [x] grpc api .proto  
- [x] grpc server  
- [x] bruteforce protector cli (grpc)
- [x] ctx
- [ ] logger
- [x] tests
- [x] Makefile
- [x] Docker
