# OpenID API

**Список методов**

| Метод | URL | Описание |
| --- | --- | --- |
| POST | /api/v3/botx/openid/refresh_access_token | Обновление access_token |

## Обновление access_token

Обрабатывается синхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **user_huid** [UUID] - `user_huid` пользователя, чей токен будет обновляться

- **ref** [UUID] (Default: null) - `sync_id` события с командой, которую не удалось выполнить, для последующего повтора отправки. Если обновление токена происходит не в результате выполнения команды, то параметр не указывается

Пример параметров запроса:

```
{
    "user_huid": "a465f0f3-1354-491c-8f11-f400164295cb",
    "ref": "a465f0f3-1354-491c-8f11-f400164295cb"
  }
```

Пример успешного ответа (202):

```
{
    "status": "ok",
    "result": true
  }
```

Ошибка от messaging сервиса (400):

```
{
    "status": "error",
    "reason": "error_from_messaging_service",
    "errors": [],
    "error_data": {
      "error": "some_error",
      "error_description": "Got error from Messaging service. Check BotX container logs (level :warn or upper) for more info."
    }
  }
```

Непредвиденная ошибка (503):

```
{
    "status": "error",
    "reason": "unexpected_error",
    "errors": [],
    "error_data": {
      "error_description": "Got unexpected error. Check BotX container logs (level :error or upper) for more info."
    }
  }
```
