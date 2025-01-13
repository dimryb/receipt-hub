# receipt-hub

Система предназначена для хранения и обработки информации о кассовых чеках. 
Пользователи могут загружать данные о чеках, просматривать историю покупок, а также получать статистику по расходам. 
Система включает основные функции работы с базой данных, API, микросервисную архитектуру и деплой в контейнерах.

---

## Swagger-документация

Проект использует [Swaggo](https://github.com/swaggo/swag) для автоматической генерации Swagger-документации на основе комментариев в коде.

### Генерация документации

Для генерации или обновления Swagger-документации выполните следующую команду в корне проекта:

```bash
swag init -g .\cmd\main.go
```

### Результат генерации

После выполнения команды документация будет сохранена в директории `docs`:

- **`docs/swagger.json`** — описание API в формате JSON.
- **`docs/swagger.yaml`** — описание API в формате YAML.
- **`docs/docs.go`** — автоматически сгенерированный Go-код для подключения Swagger.

### Просмотр документации

1. Убедитесь, что ваш сервер настроен для обработки Swagger-документации (например, с использованием middleware `gin-swagger`).
2. Запустите сервер и перейдите в браузере по адресу:
   ```
   http://localhost:8080/swagger/index.html
   ```

### Установка Swaggo

Если `swag` еще не установлен, добавьте его в проект:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

