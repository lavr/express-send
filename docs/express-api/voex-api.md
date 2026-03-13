# VoEx API

**Список методов**

| Метод | URL | Описание |
| --- | --- | --- |
| GET | /api/v3/botx/voex/conferences/:call_id | Получить информацию о конференции |
| GET | /api/v3/botx/voex/calls/:call_id | Получить информацию о звонке |

## Получить информацию о конференции

Обрабатывается синхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **call_id** [UUID] - идентификатор конференции

Пример успешного ответа (200):

```
{
    "status": "ok",
    "result": {
      "call_id": "a465f0f3-1354-491c-8f11-f400164295cb",
      "link": "https://example.com/join",
      "members": [
        "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
        "6fa5f1e9-1453-0ad7-2d6d-b791467e382a"
      ],
      "name": "Test Conference"
    }
  }
```

## Получить информацию о звонке

Обрабатывается синхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **call_id** [UUID] - идентификатор звонка

Пример успешного ответа (200):

```
{
    "status": "ok",
    "result": {
      "call_id": "a465f0f3-1354-491c-8f11-f400164295cb",
      "members": [
        "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
        "6fa5f1e9-1453-0ad7-2d6d-b791467e382a"
      ]
    }
  }
```
