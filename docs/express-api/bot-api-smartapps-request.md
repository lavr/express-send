# Синхронный запрос данных у SmartApp

Обрабатывается синхронно.

**Параметры запроса:**

- **bot_id** [UUID]

- **group_chat_id** [UUID]

- **sender_info** [Object]

  - **user_huid** [UUID]

  - **platform** [String]

  - **udid** [UUID]

- **method** [String]

- **payload** [Object] (default: null)

Пример запроса:

```
{
  "bot_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
  "group_chat_id": "8dada2c8-67a6-4434-9dec-570d244e78ee",
  "sender_info": {
    "user_huid": "ab103983-6001-44e9-889e-d55feb295494",
    "platform": "web",
    "udid": "49eac56a-c0d8-51d7-863e-925028f05110"
  },
  "method": "list.get",
  "payload": {
    "ref": "6fafda2c-6505-57a5-a088-25ea5d1d0364",
    "data": {"category_id": 1},
    "files": [
      {
        "file": "/uploads/files/3b1ee528-385e-0fb4-3bf5-62e9d72e4667/b0232da0bf3d406eb5653e37b2bb6517.bin?v=1713362406740",
        "file_name": "cts1-test.ast-innovation.ru.har",
        "file_size": 349372,
        "file_hash": "qVSzEUJITWP+TgCvcF3UCzQrBaY3RHqB92CHObz4E70=",
        "file_mime_type": "application/octet-stream",
        "chunk_size": 2097152,
        "file_encryption_algo": "stream",
        "file_id": "a0ec914f-8235-5021-9b8d-05c3cd303536",
        "type": "document"
      }
    ]
  },
}
```

**Параметры ответа:**

- **status** [String] - ok | error

- **result** [Any]

Пример успешного ответа:

```
{
  "status": "ok",
  "result": {
    "data": {
      "document_id": 1,
      "name": "Document_1"
    },
    "files": []
  }
}
```

Пример неуспешного ответа (404):

```
{
  "status": "error",
  "reason": "document_not_found",
  "errors": [],
  "error_data": {"document_id": 1}
}
```
