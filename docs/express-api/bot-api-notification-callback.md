# Обработка результата отправки бот-сообщения

Обрабатывается асинхронно.

**Заголовки:**

- Authorization - "Bearer {token}"

**Параметры запроса**

Успешный результат:

- **sync_id** [UUID5] - sync id отправляемого сообщения

- **status** [String] - "ok"

- **result** [Object] (Default: {}) - результат успешного запроса

Пример:

```
{
  "sync_id": "a7ffba12-8d0a-534e-8896-a0aa2d93a434",
  "status": "ok",
  "result": {}
}
```

Ошибка во время отправки:

- **sync_id** [UUID5] - sync id отправляемого сообщения

- **status** [String] - "error"

- **reason** [String] - краткая/общая причина ошибки

- **errors** [Array[String]] (Default: []) - детальное описание ошибки/ошибок

- **error_data** [Object] (Default: {}) - метаданные ошибки

Пример:

```
{
  "sync_id": "a7ffba12-8d0a-534e-8896-a0aa2d93a434",
  "status": "error",
  "reason": "chat_not_found",
  "errors": ["chat with specified id not found"],
  "error_data": {"group_chat_id": "918da23a-1c9a-506e-8a6f-1328f1499ee8"}
}
```
