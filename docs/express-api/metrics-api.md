# Metrics API

**Список методов**

| Метод | URL | Описание |
| --- | --- | --- |
| POST | /api/v3/botx/metrics/bot_function | Добавить новое использование функционала |

## Добавить новое использование функционала

Обрабатывается асинхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **user_huids** [Array[UUID]] - массив `user_huid`, участвующих в использовании фукнционала

- **group_chat_id** [UUID] - ID чата, в котором использовалcя функционал

- **bot_function** [String] - имя использованного функционала

Пример параметров запроса:

```
{
    "user_huids": ["a465f0f3-1354-491c-8f11-f400164295cb"],
    "group_chat_id": "a465f0f3-1354-491c-8f11-f400164295cb",
    "bot_function": "email_sent"
  }
```

Пример успешного ответа (200):

```
{
    "status": "ok",
    "result": true
  }
```
