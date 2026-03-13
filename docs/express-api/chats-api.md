# Chats API

**Список методов**

| Метод | URL | Описание |
| --- | --- | --- |
| GET | /api/v3/botx/chats/list | Получение списка чатов бота |
| GET | /api/v3/botx/chats/info | Получение информации о чате |
| POST | /api/v3/botx/chats/add_user | Добавление юзеров в чат |
| POST | /api/v3/botx/chats/remove_user | Удаление юзеров из чата |
| POST | /api/v3/botx/chats/add_admin | Добавление администратора в чат |
| POST | /api/v3/botx/chats/stealth_set | Включение стелс-режима в чате |
| POST | /api/v3/botx/chats/stealth_disable | Отключение стелс-режима в чате |
| POST | /api/v3/botx/chats/create | Создание чата |
| POST | /api/v3/botx/chats/pin_message | Закрепление сообщения в чате |
| POST | /api/v3/botx/chats/unpin_message | Открепление сообщения в чате |

## Получение списка чатов бота

Обрабатывается синхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

Отсутствуют

Пример успешного ответа:

```
{
    "status": "ok",
    "result": [
      {
        "group_chat_id": "740cf331-d833-5250-b5a5-5b5cbc697ff5",
        "chat_type": "group_chat",
        "name": "Chat Name",
        "description": "Desc",
        "members": [
          "6fafda2c-6505-57a5-a088-25ea5d1d0364",
          "705df263-6bfd-536a-9d51-13524afaab5c"
        ],
        "shared_history": false,
        "inserted_at": "2019-08-29T11:22:48.358586Z",
        "updated_at": "2019-08-30T21:02:10.453786Z"
      }
    ]
  }
```

## Получение информации о чате

Обрабатывается синхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **group_chat_id** [UUID] - идентификатор чата

Пример параметров запроса:

```
?group_chat_id=dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4
```

Пример успешного ответа:

```
{
    "status": "ok",
    "result": {
      "chat_type": "group_chat",
      "creator": "6fafda2c-6505-57a5-a088-25ea5d1d0364",
      "description": null,
      "group_chat_id": "740cf331-d833-5250-b5a5-5b5cbc697ff5",
      "inserted_at": "2019-08-29T11:22:48.358586Z",
      "members": [
        {
          "admin": true,
          "server_id": "32bb051e-cee9-5c5c-9c35-f213ec18d11e",
          "user_huid": "6fafda2c-6505-57a5-a088-25ea5d1d0364",
          "user_kind": "user"
        },
        {
          "admin": false,
          "server_id": "32bb051e-cee9-5c5c-9c35-f213ec18d11e",
          "user_huid": "705df263-6bfd-536a-9d51-13524afaab5c",
          "user_kind": "botx"
        }
      ],
      "name": "Group Chat Example",
      "shared_history": false
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

Messaging-сервис вернул ошибку (500):

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

## Добавление юзеров в чат

Обрабатывается синхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **group_chat_id** [UUID] - идентификатор чата

- **user_huids** [Array[UUID]] - список добавляемых юзеров

Пример запроса:

```
{
    "group_chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
    "user_huids": ["a465f0f3-1354-491c-8f11-f400164295cb"]
  }
```

Пример успешного ответа:

```
{
    "status": "ok",
    "result": true
  }
```

Чат не найден (404):

```
{
    "status": "error",
    "reason": "chat_not_found",
    "errors": ["Chat not found"],
    "error_data": {
      "group_chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4"
    }
  }
```

Бот не является админом чата (403):

```
{
    "status": "error",
    "reason": "no_permission_for_operation",
    "errors": ["Sender is not chat admin"],
    "error_data": {
      "group_chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
      "sender": "a465f0f3-1354-491c-8f11-f400164295cb"
    }
  }
```

Редактирование списка участников персонального чата невозможно (403):

```
{
    "status": "error",
    "reason": "chat_members_not_modifiable",
    "errors": ["Сan't add users to personal chats"],
    "error_data": {
      "group_chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4"
    }
  }
```

## Удаление юзеров из чата

Обрабатывается синхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **group_chat_id** [UUID] - идентификатор чата

- **user_huids** [Array[UUID]] - список удаляемых юзеров

Пример запроса:

```
{
    "group_chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
    "user_huids": ["a465f0f3-1354-491c-8f11-f400164295cb"]
  }
```

Пример успешного ответа:

```
{
    "status": "ok",
    "result": true
  }
```

Чат не найден (404):

```
{
    "status": "error",
    "reason": "chat_not_found",
    "errors": ["Chat not found"],
    "error_data": {
      "group_chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4"
    }
  }
```

Бот не является админом чата (403):

```
{
    "status": "error",
    "reason": "no_permission_for_operation",
    "errors": ["Sender is not chat admin"],
    "error_data": {
      "group_chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
      "sender": "a465f0f3-1354-491c-8f11-f400164295cb"
    }
  }
```

Редактирование списка участников персонального чата невозможно (403):

```
{
    "status": "error",
    "reason": "chat_members_not_modifiable",
    "errors": ["Сan't remove users from personal chats"],
    "error_data": {
      "group_chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4"
    }
  }
```

## Добавление администратора в чат

Обрабатывается синхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **group_chat_id** [UUID] - идентификатор чата

- **user_huids** [Array[UUID]] - список юзеров чата, которые назначаются администратором

Пример запроса:

```
{
    "group_chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
    "user_huids": ["a465f0f3-1354-491c-8f11-f400164295cb"]
  }
```

Пример успешного ответа:

```
{
    "status": "ok",
    "result": true
  }
```

Чат не найден (404):

```
{
    "status": "error",
    "reason": "chat_not_found",
    "errors": ["Chat not found"],
    "error_data": {
      "group_chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4"
    }
  }
```

Бот не является админом чата (403):

```
{
    "status": "error",
    "reason": "no_permission_for_operation",
    "errors": ["Sender is not chat admin"],
    "error_data": {
      "group_chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
      "sender": "a465f0f3-1354-491c-8f11-f400164295cb"
    }
  }
```

Редактирование списка администраторов персонального чата невозможно (400):

```
{
    "status": "error",
    "reason": "chat_members_not_modifiable",
    "errors": [],
    "error_data": {}
  }
```

## Включение стелс-режима в чате

Обрабатывается синхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **group_chat_id** [UUID] - идентификатор чата

- **disable_web** [Boolean] (Default: false) - если `true` - отключает доступ к чату с веб-клиента (по умолчанию `false`)

- **burn_in** [Integer] (Default: null) - время сгорания для прочитавшего, в секундах (по умолчанию `null` - выключено)

- **expire_in** [Integer] (Default: null) - время сгорания для всех участников чата, в секундах (по умолчанию `null` - выключено)

Пример запроса:

```
{
    "group_chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
    "disable_web": false,
    "burn_in": 60,
    "expire_in": 3200
  }
```

Пример успешного ответа:

```
{
    "status": "ok",
    "result": true
  }
```

Чат не найден (404):

```
{
    "status": "error",
    "reason": "chat_not_found",
    "errors": ["Chat not found"],
    "error_data": {
      "group_chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4"
    }
  }
```

Бот не является админом чата (403):

```
{
    "status": "error",
    "reason": "no_permission_for_operation",
    "errors": ["Sender is not chat admin"],
    "error_data": {
      "group_chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
      "sender": "a465f0f3-1354-491c-8f11-f400164295cb"
    }
  }
```

## Отключение стелс-режима в чате

Обрабатывается синхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **group_chat_id** [UUID] - идентификатор чата

Пример запроса:

```
{
    "group_chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4"
  }
```

Пример успешного ответа:

```
{
    "status": "ok",
    "result": true
  }
```

Чат не найден (404):

```
{
    "status": "error",
    "reason": "chat_not_found",
    "errors": ["Chat not found"],
    "error_data": {
      "group_chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4"
    }
  }
```

Бот не является админом чата (403):

```
{
    "status": "error",
    "reason": "no_permission_for_operation",
    "errors": ["Sender is not chat admin"],
    "error_data": {
      "group_chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
      "sender": "a465f0f3-1354-491c-8f11-f400164295cb"
    }
  }
```

## Создание чата

Обрабатывается синхронно.

В персональных чатах `name` игнорируется клиентами, можно указать любое, например "Personal chat".

Создателем чата будет являться бот, который инициировал запрос.

Бот будет добавлен в список участников чата, даже если он явно не указан в списке при отправке запроса.

Чтобы бот мог создавать чат, в панели администратора ему нужно задать свойство `allow_chat_creating` в значение `true`.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **name** [String] - имя чата

- **description** [String] (Default: null) - описание чата

- **chat_type** [String] - тип чата (`chat|group_chat|channel`)

- **members** [Array[UUID]] - список HUID участников чата

- **avatar** [String] (Default: null) - аватар чата в формате `data URL + base64 data (RFC 2397)`

- **shared_history** [Boolean] (Default: false) - если `true`, то созданный чат будет иметь признаки шаред-чата. Для персональных чатов всегда `false`.

Пример запроса:

```
{
    "name": "New Chat",
    "description": "Simple group chat",
    "chat_type": "group_chat",
    "members": [
      "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
      "a465f0f3-1354-491c-8f11-f400164295cb"
    ],
    "avatar": "data:image/png;base64,eDnXAc1FEUB0VFEFctII3lRlRBcetROeFfduPmXxE/8=",
    "shared_history": false
  }
```

Пример успешного ответа:

```
{
    "status": "ok",
    "result": {
      "chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4"
    }
  }
```

Боту запрещено создавать чаты (403):

```
{
    "status": "error",
    "reason": "chat_creation_is_prohibited",
    "errors": ["This bot is not allowed to create chats"],
    "error_data": {
      "bot_id": "a465f0f3-1354-491c-8f11-f400164295cb"
    }
  }
```

Ошибка создания чата (422):

```
{
    "status": "error",
    "reason": "|specified reason|",
    "errors": ["|specified errors|"],
    "error_data": {}
  }
```

## Закрепление сообщения в чате

Обрабатывается синхронно.

Бот закрепляет указанное сообщение в чате. Для закрепления в каналах у бота должны быть права администратора.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **chat_id** [UUID] - идентификатор чата

- **sync_id** [UUID] - идентификатор сообщения

Пример запроса:

```
{
    "chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
    "sync_id": "465f0f3-1354-491c-8f11-f400164295cb"
  }
```

Пример успешного ответа:

```
{
    "status": "ok",
    "result": "message_pinned"
  }
```

Бот не может закрепить сообщение (403):

```
{
    "status": "error",
    "reason": "no_permission_for_operation",
    "errors": [],
    "error_data": {
      "error_description": "Bot doesn't have permission for this operation in current chat",
      "group_chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
      "bot_id": "a465f0f3-1354-491c-8f11-f400164295cb"
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
      "error_description": "Chat with specified id not found",
      "group_chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4"
    }
  }
```

Messaging-сервис вернул ошибку (500):

```
{
    "status": "error",
    "reason": "error_from_messaging_service",
    "errors": [],
    "error_data": {
      "error_description": "Messaging service returns error. Check BotX container logs (level :warn or upper) for more info.",
      "group_chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
      "reason": "some_error"
    }
  }
```

## Открепление сообщения в чате

Обрабатывается синхронно.

Бот открепляет указанное сообщение в чате. Для открепления в каналах у бота должны быть права администратора.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **chat_id** [UUID] - идентификатор чата

Пример запроса:

```
{
    "chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4"
  }
```

Пример успешного ответа:

```
{
    "status": "ok",
    "result": "message_unpinned"
  }
```

Бот не может открепить сообщение (403):

```
{
    "status": "error",
    "reason": "no_permission_for_operation",
    "errors": [],
    "error_data": {
      "error_description": "Bot doesn't have permission for this operation in current chat",
      "group_chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
      "bot_id": "a465f0f3-1354-491c-8f11-f400164295cb"
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
      "error_description": "Chat with specified id not found",
      "group_chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4"
    }
  }
```

Messaging-сервис вернул ошибку (500):

```
{
    "status": "error",
    "reason": "error_from_messaging_service",
    "errors": [],
    "error_data": {
      "error_description": "Messaging service returns error. Check BotX container logs (level :warn or upper) for more info.",
      "group_chat_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
      "reason": "some_error"
    }
  }
```
