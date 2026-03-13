# Events API

**Список методов**

| Метод | URL | Описание |
| --- | --- | --- |
| POST | /api/v3/botx/events/edit_event | Редактирование события |
| POST | /api/v3/botx/events/reply_event | Ответ на сообщение (reply) |
| GET | /api/v3/botx/events/:sync_id/status | Получение статуса сообщения |
| POST | /api/v3/botx/events/typing | Отправка typing |
| POST | /api/v3/botx/events/stop_typing | Отправка stop_typing |
| POST | /api/v3/botx/events/delete_event | Удаление сообщения |

## Редактирование события

Редактирование содержимого события (результата команды или нотификации).

Можно редактировать только события, отправленные текущим чат-ботом.

Обрабатывается асинхронно.

**Логика обновления полей события**

Если поле не указано в запросе, то оно не обновляется.

Для полей `keyboard/bubble/mentions`, если в запросе указан пустой массив `[]`, то набор становится пустым.

Если обновлен текст, но при этом есть меншны, то существующие меншны будут добавлены в новый текст.

Если обновляется список меншнов, но есть текст, то новый список меншнов будет добавлен в текст (а старый будет удален из текста).

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **sync_id** [UUID] - `uuid` редактируемого события

-
**payload** [Object]

  - **status** [String] - "ok" | "error" (Default: Skip). Служит идентификатором успешности или провала выполнения команды.

  - **body** [String] (Default: Skip) - текстовое сообщение. Отображается в чате как текстовое сообщение.

  - **metadata** [Object] (Default: Skip) - метаданные, которые будут отправлены в параметрах команды при нажатие на любую кнопку.

  - **opts** [Object] (Default: {}) - опции сообщения

    - **silent_response** [Boolean] (Default: false) - если значение `true`, то последующие сообщения пользователя не будут отображаться в чате, до тех пор пока не придет сообщения бота, у которого это значение будет установлено в `false`. Разрешено только в личных чатах (1-1)

    - **buttons_auto_adjust** [Boolean] (Default: false) - если значение `true`, то кнопки при отображении будут автоматически переноситься на новый ряд, если не помещаются в заданном ряду

  - **keyboard** [Array[Array[Object]]] (Default: Skip) - кнопки команд, расположенные на клавиатуре, представленные в виде двумерного массива объектов.

    - **command** [String] - тело команды

    - **label** [String] - наименование команды

    - **data** [Object] (Default: {}) - объект с данными, которые будут отправлены в качестве параметров команды при нажатии на кнопку

    - **opts** [Object] - объект с клиентскими опциями кнопки

      - **silent** [Boolean] (Default: false) - если значение `true`, то при нажатие на кнопку в чат не будет отправлено сообщение с текстом команды и сама команда отправится боту в фоне

      - **h_size** [Integer] (Default: 1) - размер кнопки по горизонтали

      - **show_alert** [Boolean] (Default: false) - если значение `true`, то при нажатии на кнопку отобразится всплывающее уведомление с заданным в `alert_text` сообщением

      - **alert_text** [String] (Default: null) - текст уведомления. Если значение `null`, то выведется тело команды

      - **font_color** [String] (Default: null) - цвет текста в hex-формате

      - **background_color** [String] (Default: null) - цвет фона/границ в hex-формате

      - **align** [String] (Default: "left") - выравнивание текста (`left|center|right`)

  - **bubble** [Array[Array[Object]]] (Default: Skip) - кнопки команд расположенные под сообщением, представленные в виде двумерного массива объектов.

    - **command** [String] - тело команды

    - **label** [String] - наименование команды

    - **data** [Object] (Default: {}) - объект с данными, которые будут отправлены в качестве параметров команды при нажатие на кнопку

    - **opts** [Object] - объект с клиентскими опциями кнопки

      - **silent** [Boolean] (Default: false) - если значение `true`, то при нажатии на кнопку в чат не будет отправлено сообщение с текстом команды и сама команда отправится боту в фоне

      - **h_size** [Integer] (Default: 1) - размер кнопки по горизонтали

      - **show_alert** [Boolean] (Default: false) - если значение `true`, то при нажатии на кнопку отобразится всплывающее уведомление с заданным в `alert_text` сообщением

      - **alert_text** [String] (Default: null) - текст уведомления. Если значение `null`, то выведется тело команды

      - **font_color** [String] (Default: null) - цвет текста в hex-формате

      - **background_color** [String] (Default: null) - цвет фона/границ в hex-формате

      - **align** [String] (Default: "left") - выравнивание текста (`left|center|right`)

  - **mentions** [Array[Object]] (Default: Skip) - список меншнов

    - **mention_type** [String] (Default: "user") - тип меншна (`user|chat|contact`)

    - **mention_id** [UUID5] (Default: autogenerate) - ID меншна

    -
**mention_data** [Object]

для mention_type "user" и "contact":

      - **user_huid** [UUID5] - huid пользователя которого меншат

      - **name** [String] (default: имя из Active Directory) - имя пользователя

для mention_type "chat":

      - **group_chat_id** [UUID5] - ID чата

      - **name** [String] - отображаемое имя чата

-
**file** [Object] (Default: Skip) - файл в base64-представление. Если передать `null`, то существующий файл удалится из события

  - **file_name** [String] - имя файла

  - **data** [String] - data URL + base64 data (RFC 2397)

- **opts** [Object] - опции запроса

  - **silent_response** [Boolean] (Default: false) - если значение `true`, то последующие сообщения пользователя не будут отображаться в чате, до тех пор пока не придет сообщения бота, у которого это значение будет установлено в `false`. Разрешено только в личных чатах (1-1).

  - **raw_mentions** [Boolean] (Default: false) - если `true`, то меншны не будут подставляться в начало текста сообщения, а будут подставлены в соответствие с заданным форматом

Пример запроса:

```
{
    "sync_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
    "payload": {
      "status": "ok",
      "body": "Операция завершена!",
      "bubble": [],
      "keyboard": [
        [
          {"command": "/balance", "label": "Баланс"},
          {"command": "/transactions", "label": "Транзакции"}
        ],
        [
          {"command": "/help", "label": "Помощь", "opts": {"silent": true, "h_size": 2}}
        ]
      ],
      "mentions": []
    },
    "opts": {}
  }
```

## Ответ на сообщение (reply)

Обрабатывается асинхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "multipart/form-data" (при отсутствие файла можно использовать "application/json")

**Параметры запроса:**

- **source_sync_id** [UUID] - идентификатор сообщения, на который будет отправлен ответ

-
**reply** [Object]

  - **status** [String] - "ok" | "error". Служит идентификатором успешности или провала выполнения команды.

  - **body** [String] - текстовое сообщение. Отображается в чате, как текстовое сообщение.

  - **metadata** [Object] (Default: {}) - метаданные, которые будут отправлены в параметрах команды при нажатие на любую кнопку.

  - **opts** [Object] (Default: {}) - опции сообщения

    - **silent_response** [Boolean] (Default: false) - если значение `true`, то последующие сообщения пользователя не будут отображаться в чате, до тех пор пока не придет сообщения бота, у которого это значение будет установлено в `false`. Разрешено только в личных чатах (1-1)

    - **buttons_auto_adjust** [Boolean] (Default: false) - если значение `true`, то кнопки, при отображение, будут автоматически переноситься на новый ряд, если не помещаются в заданном ряду

  - **keyboard** [Array[Array[Object]]] (Default: []) - кнопки команд расположенные на клавиатуре, представленные в виде двумерного массива объектов.

    - **command** [String] - тело команды

    - **label** [String] - наименование команды

    - **data** [Object] (Default: {}) - объект с данными, которые будут отправлены в качестве параметров команды при нажатие на кнопку

    - **opts** [Object] - объект с клиентскими опциями кнопки

      - **silent** [Boolean] (Default: false) - если значение `true`, то при нажатие на кнопку в чат не будет отправлено сообщение с текстом команды и сама команда отправится боту в фоне

      - **h_size** [Integer] (Default: 1) - размер кнопки по горизонтали

      - **show_alert** [Boolean] (Default: false) - если значение `true`, то при нажатии на кнопку отобразится всплывающее уведомление с заданным в `alert_text` сообщением

      - **alert_text** [String] (Default: null) - текст уведомления. Если значение `null`, то выведется тело команды

      - **font_color** [String] (Default: null) - цвет текста в hex-формате

      - **background_color** [String] (Default: null) - цвет фона/границ в hex-формате

      - **align** [String] (Default: "left") - выравнивание текста (`left|center|right`)

  - **bubble** [Array[Array[Object]]] (Default: []) - кнопки команд, расположенные под сообщением, представленные в виде двумерного массива объектов.

    - **command** [String] - тело команды

    - **label** [String] - наименование команды

    - **data** [Object] (Default: {}) - объект с данными, которые будут отправлены в качестве параметров команды при нажатие на кнопку

    - **opts** [Object] - объект с клиентскими опциями кнопки

      - **silent** [Boolean] (Default: false) - если значение `true`, то при нажатие на кнопку в чат не будет отправлено сообщение с текстом команды и сама команда отправится боту в фоне

      - **h_size** [Integer] (Default: 1) - размер кнопки по горизонтали

      - **show_alert** [Boolean] (Default: false) - если значение `true`, то при нажатии на кнопку отобразится всплывающее уведомление с заданным в `alert_text` сообщением

      - **alert_text** [String] (Default: null) - текст уведомления. Если значение `null`, то выведется тело команды

      - **font_color** [String] (Default: null) - цвет текста в hex-формате

      - **background_color** [String] (Default: null) - цвет фона/границ в hex-формате

      - **align** [String] (Default: "left") - выравнивание текста (`left|center|right`)

  -
**mentions** [Array[Object]] (Default: []) - список меншнов

    -
**mention_type** [String] (Default: "user") - тип меншна (`user|chat|contact`)

    -
**mention_id** [UUID5] (Default: autogenerate) - ID меншна

    -
**mention_data** [Object]

для mention_type "user" и "contact":

      - **user_huid** [UUID5] - huid пользователя которого меншат

      - **name** [String] (default: имя из Active Directory) - имя пользователя

для mention_type "chat":

      - **group_chat_id** [UUID5] - ID чата

      - **name** [String] - отображаемое имя чата

-
**file** [Binary]|[Object] (Default: null) - файл в бинарном представление (multipart-запрос) или `base64`.

  - Base64 file format

    - **file_name** [String] - имя файла

    - **data** [String] - data URL + base64 data (RFC 2397)

    - **caption** [String] (Default: null) - текст под файлом

- **opts** [Object] - опции запроса

  - **stealth_mode** [Boolean] (Default: false) - если `true`, то сообщение будет отправлено в чат только в том случае, если в чате включен стелс-режим

  - **notification_opts** [Object]

    - **send** [Boolean] (Default: true) - отправлять/не отправлять push-нотификацию

    - **force_dnd** [Boolean] (Default: false) - игнорировать/не игнорировать `DND/Mute`.

Пример запроса:

```
{
    "source_sync_id": "b165f00f-3154-412c-7f11-c120164257da",
    "reply": {
      "status": "ok",
      "body": "Хороший выбор! Я сохранил ваш вариант ответа.",
      "metadata": {},
      "opts": {"silent_response": false},
      "bubble": [],
      "keyboard": [],
      "mentions": []
    },
    "file": null,
    "opts": {
      "stealth_mode": false,
      "notification_opts": {
        "send": true,
        "force_dnd": false
      }
    }
  }
```

**Параметры успешного ответа:**

- **status** [String] (Value: "ok")

- **result** [String] (Value: "bot_reply_pushed")

Пример успешного ответа:

```
{
    "status": "ok",
    "result": "bot_reply_pushed"
  }
```

## Получение статуса сообщения

Обрабатывается синхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **sync_id** [UUID] - идентификатор события

Пример параметров запроса:

```
/api/v3/botx/events/dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4/status
```

**Пример успешного ответа:**

- **group_chat_id** [UUID] - идентификатор чата

- **sent_to** [Array[UUID]] - список huid пользователей которым было отправлено событие

- **read_by** [Array[Object]]

  - **user_huid** [UUID] - huid пользователя, прочитавшего событие

  - **read_at** [DateTime] - время прочтения

- **received_by** [Array[Object]]

  - **user_huid** [UUID] - huid пользователя, получившего событие

  - **received_at** [DateTime] - время получения

```
{
    "status": "ok",
    "result": {
      "group_chat_id": "740cf331-d833-5250-b5a5-5b5cbc697ff5",
      "sent_to": ["32bb051e-cee9-5c5c-9c35-f213ec18d11e"],
      "read_by": [
        {
          "user_huid": "6fafda2c-6505-57a5-a088-25ea5d1d0364",
          "read_at": "2019-08-29T11:22:48.358586Z"
        }
      ],
      "received_by": [
        {
          "user_huid": "6fafda2c-6505-57a5-a088-25ea5d1d0364",
          "received_at": "2019-08-29T11:22:48.358586Z"
        }
      ]
    }
  }
```

Событие не найдено (404):

```
{
  "status": "error",
  "reason": "event_not_found",
  "errors": [],
  "error_data": {
    "sync_id": "84a12e71-3efc-5c34-87d5-84e3d9ad64fd"
  }
}
```

Ошибка messaging-сервиса при запросе события (500):

```
{
  "status": "error",
  "reason": "messaging_service_error",
  "errors": [],
  "error_data": {
    "sync_id": "84a12e71-3efc-5c34-87d5-84e3d9ad64fd"
  }
}
```

## Отправка typing

Обрабатывается асинхронно.

Событие отправляется только пользователям с того же сервера, где зарегистрирован бот.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **group_chat_id** [UUID] - идентификатор чата

Пример запроса:

```
{
  "group_chat_id": "740cf331-d833-5250-b5a5-5b5cbc697ff5"
}
```

Пример успешного ответа (202):

```
{
    "status": "ok",
    "result": "typing_event_pushed"
  }
```

## Отправка stop_typing

Обрабатывается асинхронно.

Событие отправляется только пользователям с того же сервера, где зарегистрирован бот.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **group_chat_id** [UUID] - идентификатор чата

Пример запроса:

```
{
  "group_chat_id": "740cf331-d833-5250-b5a5-5b5cbc697ff5"
}
```

Пример успешного ответа (202):

```
{
    "status": "ok",
    "result": "stop_typing_event_pushed"
  }
```

## Удаление сообщения

Обрабатывается синхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **sync_id** [UUID] - идентификатор сообщения

Пример запроса:

```
{
  "sync_id": "740cf331-d833-5250-b5a5-5b5cbc697ff5"
}
```

Пример успешного ответа (202):

```
{
    "status": "ok",
    "result": "event_deleted"
  }
```

Сообщение не найдено (404):

```
{
    "status": "error",
    "reason": "sync_id_not_found",
    "errors": [],
    "error_data": {}
  }
```
