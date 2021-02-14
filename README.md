# Проектная работа - Bruteforce Protector Service

- Курс [Golang Developer. Professional](https://otus.ru/lessons/golang-professional/)
- [ТЗ](docs/specification.md)

## Сборка и тестирование
- `make run` - сборка и запуск докер образа с сервером
- `make stop` - остановка докер образа
- `make build` - сборка сервера и клиента
- `make build-server` - сборка сервера
- `make build-cli` - сборка клиента
- `make lint` - запуск линтера
- `make test` - запуск тестов
- `make release` - сборка клиента, сервера, запуск тестов и линтера
- `make generate` - генерация protobuf/grpc

## Команды bp-cli
- TBD 

## Roadmap

- [x] IP access list  
- [x] sliding window rate limiter  
- [x] bruteforce protector service methods  
- [ ] persistent storage (MongoDB/Redis)  
- [x] grpc api .proto  
- [x] grpc server  
- [x] bruteforce protector cli (grpc)
- [ ] ctx
- [ ] logger
- [ ] tests
- [x] Makefile
- [x] Docker
