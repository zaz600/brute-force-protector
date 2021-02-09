# Требования к cli-клиенту

- Ключ для указания хоста/порта `-host localhost:53211`  
- Если ключ не указан должно быть подключение к серверу на `localhost`  
- Команды для управления списками доступа
  - `cmd blacklist ls`
  - `cmd blacklist add 1.1.1.1/24`
  - `cmd blacklist remove 1.1.1.1/24`  
  - `cmd whitelist ls`
  - `cmd whitelist add 1.1.1.1/24`
  - `cmd whitelist remove 1.1.1.1/24`
    
- Команды для сброса лимитов
  - `cmd reset login johndoe`
  - `cmd reset ip 8.8.8.8`
    
- Команда для проверки лимита
  - `cmd verify johndoe qwerty 8.8.8.8`