# Files API

**Список методов**

| Метод | URL | Описание |
| --- | --- | --- |
| GET | /api/v3/botx/files/download | Скачивание файла |
| POST | /api/v3/botx/files/upload | Загрузка файла |

## Скачивание файла

Обрабатывается синхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **group_chat_id** [UUID] - идентификатор чата

- **file_id** [UUID] - идентификатор файла

- **is_preview** [Boolean] (Default: false) - если `true` и файл имеет `preview`, то вернется содержимое `preview`, иначе оригинал

Пример запроса:

```
?group_chat_id=84a12e71-3efc-5c34-87d5-84e3d9ad64fd
&file_id=e48c5612-b94f-4264-adc2-1bc36445a226
&is_preview=false
```

Успешный ответ (200):

```
HTTP/1.1 200
Content-Type: image/jpeg

```

Чат не найден (404):

```
{
  "status": "error",
  "reason": "chat_not_found",
  "errors": [],
  "error_data": {
    "group_chat_id": "84a12e71-3efc-5c34-87d5-84e3d9ad64fd",
    "error_description": "Chat with specified id not found"
  }
}
```

Ошибка messaging-сервиса при запросе чата (500):

```
{
  "status": "error",
  "reason": "error_from_messaging_service",
  "errors": [],
  "error_data": {
    "group_chat_id": "84a12e71-3efc-5c34-87d5-84e3d9ad64fd",
    "reason": "some_error",
    "error_description": "Chat info fetching failed. Check BotX container logs (level :warn or upper) for more info."
  }
}
```

Метаданные файла не найдены (404):

```
{
  "status": "error",
  "reason": "file_metadata_not_found",
  "errors": [],
  "error_data": {
    "file_id": "e48c5612-b94f-4264-adc2-1bc36445a226",
    "group_chat_id": "84a12e71-3efc-5c34-87d5-84e3d9ad64fd",
    "error_description": "File with specified file_id and group_chat_id not found in file service"
  }
}
```

Файл-сервис вернул ошибку при запросе метаданных файла (500):

```
{
  "status": "error",
  "reason": "error_from_file_service",
  "errors": [],
  "error_data": {
    "file_id": "e48c5612-b94f-4264-adc2-1bc36445a226",
    "group_chat_id": "84a12e71-3efc-5c34-87d5-84e3d9ad64fd",
    "reason": "some_error",
    "error_description": "File metadata fetching failed. Check BotX container logs (level :warn or upper) for more info."
  }
}
```

Файл удален (204):

```
{
  "status": "error",
  "reason": "file_deleted",
  "errors": [],
  "error_data": {
    "link": "/uploads/files/b9197d3a-d855-5d34-ba8a-eff3a975ab20/c6e39e617aee4613bd0d6a47738b6368.jpeg?v=1627067439978",
    "error_description": "File at the specified link has been deleted"
  }
}
```

Ошибка при скачивании файла (500):

```
{
  "status": "error",
  "reason": "file_download_failed",
  "errors": [],
  "error_data": {
    "link": "/uploads/files/b9197d3a-d855-5d34-ba8a-eff3a975ab20/c6e39e617aee4613bd0d6a47738b6368.jpeg?v=1627067439978",
    "error_description": "Got error on file downloading. Check BotX container logs (level :warn or upper) for more info."
  }
}
```

Ошибка при скачивании файла из s3 файл-сервиса (500):

```
{
  "status": "error",
  "reason": "error_from_s3_file_service",
  "errors": [],
  "error_data": {
    "source_link": "/uploads/files/b9197d3a-d855-5d34-ba8a-eff3a975ab20/c6e39e617aee4613bd0d6a47738b6368.jpeg?v=1627067439978",
    "location": "https://bucketexample.s3.amazonaws.com/E54SAFSA",
    "error_description": "Got error on downloading file from s3 file service. Check BotX container logs (level :warn or upper) for more info."
  }
}
```

Файл не имеет `preview` (400):

```
{
  "status": "error",
  "reason": "file_without_preview",
  "errors": [],
  "error_data": {
    "file_id": "e48c5612-b94f-4264-adc2-1bc36445a226",
    "group_chat_id": "84a12e71-3efc-5c34-87d5-84e3d9ad64fd",
    "error_description": "Preview was requested but file doesn't have preview"
  }
}
```

## Загрузка файла

Обрабатывается синхронно.

Содержимое файла и параметры передается в формате `multipart/form-data` с boundary.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type: multipart/form-data; boundary=<boundary string>

**Параметры запроса:**

- **group_chat_id** [UUID5] - идентификатор чата

- **file_name** [String] - имя файла. Берется из content

- **mime_type** [String] - mime-тип файла. Берется из `content`

- **content** [Binary] - бинарное содержимое файла

- **meta** [Object] - метаданные файла

  - **duration** [Integer] (Default: null) - длительность видео/аудио в секундах

  - **caption** [String] (Default: null) - `caption` файла

Пример запроса:

```
POST /api/v3/files/upload HTTP/1.1
Content-Type: multipart/form-data; boundary=4d9aa92c

--4d9aa92c
Content-Disposition: form-data; name="content"; filename="image.jpeg"
Content-Type: image/jpeg

--4d9aa92c
Content-Disposition: form-data; name="group_chat_id"

d297a43f-e1c6-4e07-8081-e25c04364acc
--4d9aa92c
Content-Disposition: form-data; name="meta"
Content-Type: application/json

{
  "duration": null,
  "caption": "Мое видео"
}
--4d9aa92c--
```

**Успешный ответ (200):**

- **status** [String] - “ok”

- **result** [Object]

  - **type** [String] - “image”

  - **file** [String] - ссылка на файл

  - **file_mime_type** [String] - mime type файла

  - **file_name** [String] - имя файла

  - **file_preview** [String] (Default: null) - ссылка на превью

  - **file_preview_height** [Integer] (Default: null) - высота файла в px

  - **file_preview_width** [Integer] (Default: null) - ширина файла в px

  - **file_size** [Integer] - размер файла

  - **file_hash** [String] - хэш файла

  - **file_encryption_algo** [String] - “stream”

  - **chunk_size** [Integer] - размер чанков

  - **file_id** [UUID5] - идентификатор файла

  - **caption** [String] (Default: null) - caption файла

  - **duration** [Integer] (Default: null) - длительность аудио/видео

```
{
  "status": "ok",
  "result": {
    "type": "image",
    "file": "https://link.to/file",
    "file_mime_type": "image/png",
    "file_name": "pass.png",
    "file_preview": "https://link.to/preview",
    "file_preview_height": 300,
    "file_preview_width": 300,
    "file_size": 1502345,
    "file_hash": "Jd9r+OKpw5y+FSCg1xNTSUkwEo4nCW1Sn1AkotkOpH0=",
    "file_encryption_algo": "stream",
    "chunk_size": 2097152,
    "file_id": "8dada2c8-67a6-4434-9dec-570d244e78ee",
    "caption": "текст",
    "duration": null
  }
}
```

Чат не найден (404):

```
{
  "status": "error",
  "reason": "chat_not_found",
  "errors": [],
  "error_data": {
    "group_chat_id": "84a12e71-3efc-5c34-87d5-84e3d9ad64fd",
    "error_description": "Chat with specified id not found"
  }
}
```

Ошибка messaging-сервиса при запросе чата (500):

```
{
  "status": "error",
  "reason": "error_from_messaging_service",
  "errors": [],
  "error_data": {
    "group_chat_id": "84a12e71-3efc-5c34-87d5-84e3d9ad64fd",
    "reason": "some_error",
    "error_description": "Chat info fetching failed. Check BotX container logs (level :warn or upper) for more info."
  }
}
```

Неправильный `mimetype` (500):

```
{
  "status": "error",
  "reason": "flow_processing_error",
  "errors": ["invalid mimetype", "unsupported file type"],
  "error_data": {
    "error_description": "Got error on flow processing. Check BotX container logs (level :info or upper) for more info."
  }
}
```

Неподдерживаемый тип файла (500):

```
{
  "status": "error",
  "reason": "flow_processing_error",
  "errors": ["unsupported file type"],
  "error_data": {
    "error_description": "Got error on flow processing. Check BotX container logs (level :info or upper) for more info."
  }
}
```

Сервис ключей (KDC) вернул ошибку при запросе ключа бота (500):

```
{
  "status": "error",
  "reason": "flow_processing_error",
  "errors": ["sender public key fetching failed"],
  "error_data": {
    "error_description": "Got error on flow processing. Check BotX container logs (level :info or upper) for more info."
  }
}
```

Сервис ключей (KDC) вернул ошибку при запросе ключей получателей (500):

```
{
  "status": "error",
  "reason": "flow_processing_error",
  "errors": ["recipients public keys fetching failed"],
  "error_data": {
    "error_description": "Got error on flow processing. Check BotX container logs (level :info or upper) for more info."
  }
}
```

Файл-сервис вернул ошибку при загрузке файла (500):

```
{
  "status": "error",
  "reason": "flow_processing_error",
  "errors": ["uploading file to file service failed", "file service error: some_error"],
  "error_data": {
    "error_description": "Got error on flow processing. Check BotX container logs (level :info or upper) for more info."
  }
}
```
