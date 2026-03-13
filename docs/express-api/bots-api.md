# Bots API

**Список методов**

| Метод | URL | Описание |
| --- | --- | --- |
| GET | /api/v2/botx/bots/:bot_id/token | Получение токена |
| GET | /api/v1/botx/bots/catalog?since=datatime | Получение списка ботов своего сервера |

## Получение токена

**Параметры запроса:**

- **bot_id** [UUID] - идентификатор чат-бота, необходим для шифрования данных

- **signature** [String] - bot_id, подписанный секретным ключом чат-бота (`secret_key`) с использованием `hmac-sha256`, представленный в виде base16 (hex) строки (используя традиционный [rfc4648](https://datatracker.ietf.org/doc/html/rfc4648#section-8) алфавит `0-9 A-F`).

Пример подписи:

```
:crypto.mac(:hmac, :sha256, "secret", "bot_id") |> Base.encode16

"904E39D3BC549C71F4A4BDA66AFCDA6FC90D471A64889B45CC8D2288E56526AD
```

Пример запроса:

```
/api/v2/botx/bots/8dada2c8-67a6-4434-9dec-570d244e78ee/token?signature=904E39D3BC549C71F4A4BDA66AFCDA6FC90D471A64889B45CC8D2288E56526AD
```

Пример ответа:

```
{
    "status": "ok",
    "result": ""
  }
```

## Получение списка ботов своего сервера

**Параметры запроса:**

- **since** [ISO-8601 timestamp] - минимально допустимое время последнего обновления бота в выборке

Пример ответа:

```
{
  "result": {
    "generated_at": ,
    "bots": [
      {
        "user_huid": "6fafda2c-6505-57a5-a088-25ea5d1d0364",
        "name": "First bot",
        "description": "My bot",
        "avatar": "https://cts.ccsteam.ru/uploads/profile_avatar/796640bd-7add-5274-9f02-7bd169b88dab/04db7577-238c-5aec-8d21-22e831839c21.jpg?v=21"
        "enabled": true,
      }
    ]
  }
}
```
