# Notifications API

**Список методов**

| Метод | URL | Описание |
| --- | --- | --- |
| POST | /api/v4/botx/notifications/direct | Отправка директ нотификации в чат |
| POST | /api/v4/botx/notifications/internal | Отправка внутренней бот-нотификации |

## Отправка директ-нотификации v4

Обрабатывается асинхронно.

В ответ на запрос приходит `sync_id [UUID]` - идентификатор отправляемого сообщения.

В случае успеха или провала отправки на коллбек чат-боту отправляется запрос с результатом отправки.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "multipart/form-data" (при отсутствии файла можно использовать "application/json")

**Параметры запроса:**

- **group_chat_id** [UUID] - идентификатор чата в который придет сообщение

- **recipients** [Array[UUID]] (Default: null) - huid получателей события. По умолчанию все участники чата.

-
**notification** [Object]

  - **status** [String] - "ok" | "error". Служит идентификатором успешности или провала выполнения команды.

  - **body** [String] - текстовое сообщение. Отображается в чате, как текстовое сообщение.

  - **metadata** [Object] (Default: {}) - метаданные, которые будут отправлены в параметрах команды при нажатие на любую кнопку.

  - **opts** [Object] (Default: {}) - опции сообщения

    - **silent_response** [Boolean] (Default: false) - если значение `true`, то последующие сообщения пользователя не будут отображаться в чате, до тех пор пока не придет сообщения бота у которого это значение будет установлено в `false`. Разрешено только в личных чатах (1-1)

    - **buttons_auto_adjust** [Boolean] (Default: false) - если значение `true`, то кнопки при отображении будут автоматически переноситься на новый ряд, если не помещаются в заданном ряду

  - **keyboard** [Array[Array[Object]]] (Default: []) - кнопки команд расположенные на клавиатуре, представленные в виде двумерного массива объектов.

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

  - **bubble** [Array[Array[Object]]] (Default: []) - кнопки команд расположенные под сообщением, представленные в виде двумерного массива объектов.

    - **command** [String] - тело команды

    - **label** [String] - наименование команды

    - **data** [Object] (Default: {}) - объект с данными, которые будут отправлены в качестве параметров команды при нажатие на кнопку

    - **opts** [Object] - объект с клиентскими опциями кнопки

      - **silent** [Boolean] (Default: false) - если значение `true`, то при нажатии на кнопку в чат не будет отправлено сообщение с текстом команды и сама команда отправится боту в фоне

      - **h_size** [Integer] (Default: 1) - размер кнопки по горизонтали

      - **show_alert** [Boolean] (Default: false) - если значение `true`, то при нажатии на кнопку отобразится всплывающее уведомление с заданным в `alert_text` сообщением

      - **alert_text** [String] (Default: null) - текст уведомления. Если значение `null`, то выведется тело команды

      - **font_color** [String] (Default: null) - цвет текста в hex формате

      - **background_color** [String] (Default: null) - цвет фона/границ в hex-формате

      - **align** [String] (Default: "left") - выравнивание текста (`left|center|right`)

  -
**mentions** [Array[Object]] (Default: []) - список меншнов

****
****
****

> **Note**
> Примечание
> Отображение меншнов в тексте (body) задается по следующему шаблону:
> @{mention:
> mention_id
> } - для user/all mention
> @@{mention:
> mention_id
> } - для contact mention
> ##{mention:
> mention_id
> } - для chat/channel mention
> Например: @{mention:a465f0f3-1354-491c-8f11-f400164295cb}

    - **mention_type** [String] (Default: "user") - тип меншна (`user|chat|channel|contact|all`)

    - **mention_id** [UUID5] - ID меншна

    -
**mention_data** [Object]

для mention_type "all" - **null**

для mention_type "user" и "contact":

      - **user_huid** [UUID5] - huid пользователя, которого меншат

      - **name** [String] - имя пользователя

для mention_type "chat" и "channel":

      - **group_chat_id** [UUID5] - ID чата

      - **name** [String] - отображаемое имя чата

-
**file** [Object] (Default: null) - файл в base64-представлении.

  - **file_name** [String] - имя файла

  - **data** [String] - data URL + base64 data (RFC 2397)

- **opts** [Object] - опции запроса

  - **stealth_mode** [Boolean] (Default: false) - если `true`, то сообщение будет отправлено в чат только в том случае, если в чате включен стелс-режим

  - **notification_opts** [Object]

    - **send** [Boolean] (Default: true) - отправлять/не отправлять push-нотификацию

    - **force_dnd** [Boolean] (Default: false) - игнорировать/не игнорировать DND/Mute.

Пример запроса:

```
{
    "group_chat_id": "a465f0f3-1354-491c-8f11-f400164295cb",
    "recipients": null,
    "notification": {
      "status": "ok",
      "body": "@{mention:a465f0f3-1354-491c-8f11-f400164295cb} Выберите пункт меню",
      "metadata": {"account_id": 94},
      "opts": {"silent_response": true},
      "bubble": [
        [
          {"command": "/profile", "label": "Профиль", "opts": {"background_color": "#010205", "align": "center"}}
        ],
        [
          {"command": "/balance", "label": "Баланс", "data": {"operation_id": 1}},
          {"command": "/transactions", "label": "Транзакции", "data": {"operation_id": 2}}
        ]
      ],
      "keyboard": [
        [
          {"command": "/help", "label": "Помощь", "opts": {"silent": true, "h_size": 2}}
        ]
      ],
      "mentions": [
        {
          "mention_id": "a465f0f3-1354-491c-8f11-f400164295cb",
          "mention_type": "user",
          "mention_data": {
            "user_huid": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
            "name": "Gena the Croc"
          }
        }
      ]
    },
    "file": {
      "data": "data:image/png;base64,eDnXAc1FEUB0VFEFctII3lRlRBcetROeFfduPmXxE/8=",
      "file_name": "card.png"
    },
    "opts": {
      "stealth_mode": false,
      "notification_opts": {
        "send": true,
        "force_dnd": false
      }
    }
  }
```

**Успешный ответ (202):**

- **status** [String] (Value: "ok")

- **result** [Object]

  - **sync_id** [UUID] - идентификатор отправляемого сообщения

```
{
    "status": "ok",
    "result": {
      "sync_id": "f0f105d2-101f-59b0-9e10-e432efce2c36"
    }
  }
```

**Результат и возможные ошибки**

Следующий результат или ошибки будут отправлены чат-боту на callback-метод.

Успех:

```
{
  "sync_id": "f0f105d2-101f-59b0-9e10-e432efce2c36",
  "status": "ok",
  "result": {}
}
```

Чат не найден:

```
{
  "sync_id": "f0f105d2-101f-59b0-9e10-e432efce2c36",
  "status": "error",
  "reason": "chat_not_found",
  "errors": [],
  "error_data": {
    "group_chat_id": "705df263-6bfd-536a-9d51-13524afaab5c",
    "error_description": "Chat with specified id not found"
  }
}
```

Ошибка от Messaging-сервиса:

```
{
  "sync_id": "f0f105d2-101f-59b0-9e10-e432efce2c36",
  "status": "error",
  "reason": "error_from_messaging_service",
  "errors": [],
  "error_data": {
    "group_chat_id": "705df263-6bfd-536a-9d51-13524afaab5c",
    "reason": "some_error",
    "error_description": "Messaging service returns error. Check BotX container logs (level :warn or upper) for more info."
  }
}
```

Бот-отправитель не является участником чата:

```
{
  "sync_id": "f0f105d2-101f-59b0-9e10-e432efce2c36",
  "status": "error",
  "reason": "bot_is_not_a_chat_member",
  "errors": [],
  "error_data": {
    "group_chat_id": "705df263-6bfd-536a-9d51-13524afaab5c",
    "bot_id": "b165f00f-3154-412c-7f11-c120164257da",
    "error_description": "Bot is not a chat member"
  }
}
```

Стелс-режим отключен в чате, но требуется опцией:

```
{
  "sync_id": "f0f105d2-101f-59b0-9e10-e432efce2c36",
  "status": "error",
  "reason": "stealth_mode_disabled",
  "errors": [],
  "error_data": {
    "group_chat_id": "705df263-6bfd-536a-9d51-13524afaab5c",
    "bot_id": "b165f00f-3154-412c-7f11-c120164257da",
    "error_description": "Stealth mode disabled in specified chat"
  }
}
```

Итоговый список получателей события пуст:

```
{
  "sync_id": "f0f105d2-101f-59b0-9e10-e432efce2c36",
  "status": "error",
  "reason": "event_recipients_list_is_empty",
  "errors": [],
  "error_data": {
    "group_chat_id": "705df263-6bfd-536a-9d51-13524afaab5c",
    "bot_id": "b165f00f-3154-412c-7f11-c120164257da",
    "recipients_param": ["b165f00f-3154-412c-7f11-c120164257da"],
    "error_description": "Event recipients list is empty"
  }
}
```

Ошибка в процессе обработки:

```
{
  "sync_id": "f0f105d2-101f-59b0-9e10-e432efce2c36",
  "status": "error",
  "reason": "flow_processing_error",
  "errors": ["{error1}", "{error2}"],
  "error_data": {
    "error_description": "Got error on flow processing. Check BotX container logs (level :info or upper) for more info"
  }
}
```

Непредвиденная ошибка:

```
{
  "sync_id": "f0f105d2-101f-59b0-9e10-e432efce2c36",
  "status": "error",
  "reason": "unexpected_error",
  "errors": ["{error1}", "{error2}"],
  "error_data": {
    "error_description": "Got unexpected error. Check BotX container logs (level :error or upper) for more info"
  }
}
```

**Ошибки загрузки файла**

Список возможных ошибок можно найти в описание метода [Загрузка файла](../files-api/#загрузка-файла).

## Отправка внутренней бот-нотификации

Обрабатывается асинхронно.

В ответ на запрос приходит `sync_id [UUID]` - идентификатор отправляемого сообщения.

В случае успеха или провала отправки на коллбек чат-боту отправляется запрос с результатом отправки.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **group_chat_id** [UUID] - ID чата

- **data** [Object] - пользовательские данные

- **opts** [Object] (default: {}) - пользовательские опции

- **recipients** [Array[UUID]] (default: Null) - список ботов получателей. Если не передан, то отправляется всем ботам в чате

```
{
  "group_chat_id": "705df263-6bfd-536a-9d51-13524afaab5c",
  "data": {
    "message": "ping"
  },
  "opts": {
    "internal_token": "FEFctII3lRlRBcetROeFfduPmXxE"
  },
  "recipients": null
}
```

**Успешный ответ (202)**

- **status** [String] (Value: "ok")

- **result** [Object]

  - **sync_id** [UUID] - идентификатор отправляемого сообщения

```
{
    "status": "ok",
    "result": {
      "sync_id": "f0f105d2-101f-59b0-9e10-e432efce2c36"
    }
  }
```

Превышен лимит интенсивности запросов (429):

```
{
    "status": "error",
    "reason": "too_many_requests",
    "errors": [],
    "error_data": {
      "bot_id": "b165f00f-3154-412c-7f11-c120164257da"
    }
  }
```

**Результат и возможные ошибки**

Следующий результат или ошибки будут отправлены чат-боту на callback-метод.

Успех:

```
{
  "sync_id": "f0f105d2-101f-59b0-9e10-e432efce2c36",
  "status": "ok",
  "result": {}
}
```

Чат не найден:

```
{
  "sync_id": "f0f105d2-101f-59b0-9e10-e432efce2c36",
  "status": "error",
  "reason": "chat_not_found",
  "errors": [],
  "error_data": {
    "group_chat_id": "705df263-6bfd-536a-9d51-13524afaab5c",
    "error_description": "Chat with specified id not found"
  }
}
```

Ошибка от Messaging-сервиса:

```
{
  "sync_id": "f0f105d2-101f-59b0-9e10-e432efce2c36",
  "status": "error",
  "reason": "error_from_messaging_service",
  "errors": [],
  "error_data": {
    "group_chat_id": "705df263-6bfd-536a-9d51-13524afaab5c",
    "reason": "some_error",
    "error_description": "Messaging service returns error. Check BotX container logs (level :warn or upper) for more info."
  }
}
```

Бот-отправитель не является участником чата:

```
{
  "sync_id": "f0f105d2-101f-59b0-9e10-e432efce2c36",
  "status": "error",
  "reason": "bot_is_not_a_chat_member",
  "errors": [],
  "error_data": {
    "group_chat_id": "705df263-6bfd-536a-9d51-13524afaab5c",
    "bot_id": "b165f00f-3154-412c-7f11-c120164257da",
    "error_description": "Bot is not a chat member"
  }
}
```

Итоговый список получателей события пуст:

```
{
  "sync_id": "f0f105d2-101f-59b0-9e10-e432efce2c36",
  "status": "error",
  "reason": "event_recipients_list_is_empty",
  "errors": [],
  "error_data": {
    "group_chat_id": "705df263-6bfd-536a-9d51-13524afaab5c",
    "bot_id": "b165f00f-3154-412c-7f11-c120164257da",
    "recipients_param": ["b165f00f-3154-412c-7f11-c120164257da"],
    "error_description": "Event recipients list is empty"
  }
}
```
