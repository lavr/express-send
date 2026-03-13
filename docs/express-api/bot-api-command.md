# Обработка команд

Обрабатывается асинхронно.

**Заголовки:**

- authorization - "Bearer {token}"

**Параметры запроса:**

- **sync_id** [UUID] - идентификатор сообщения в системе Express

- **source_sync_id** [UUID] (Default: null) - идентификатор исходного сообщения (сообщения в котором находились элементы интерфейса) в системе Express

- **command** [Object]

  - **body** [String] - тело команды

  - **command_type** [String] - тип команды (user|system)

  - **data** [Object] - данные команды полученные на основе введеных пользователем данных через элементы UI или при нажатие на кнопку

  - **metadata** [Object] - метаданные заложенные в объекте сообщения от бота

- **attachments** [Array[Object]] - вложения, переданные в сообщении. Например: изображения, видео, файлы, ссылки, геолокации, контакты

  - **type** [String] - тип вложения

  - **data** [Object] - объект с данными вложения

- **from** [Object]

  - **user_huid** [UUID] (Default: null) - huid юзера который отправил команду

  - **group_chat_id** [UUID] (Default: null) - id чата в который отправили команда

  - **chat_type** [String] (Default: null) - тип чата (chat|group_chat|channel)

  - **ad_login** [String] (Default: null) - логин юзера который отправил команду

  - **ad_domain** [String] (Default: null) - домен юзера который отправил команду

  - **username** [String] (Default: null) - имя юзера который отправил команду

  - **is_admin** [Boolean] (Default: null) - является ли юзер админом чата

  - **is_creator** [Boolean] (Default: null) - является ли юзер создателем чата

  - **manufacturer** [String] (Default: null) - имя бренда производителя

  - **device** [String] (Default: null) - название девайса

  - **device_software** [String] (Default: null) - ОС девайса

  - **device_meta** [Object] (Default: null)

    - **pushes** [Boolean] - разрешение приложению на отправку пушей

    - **timezone** [String] - таймзона пользователя

    - **permissions** [Object] - различные разрешения приложения (использование микрофона, камеры и т.д.)

  - **platform** [String] (Default: null) - название клиентской платформы (web|android|ios|desktop)

  - **platform_package_id** [String] (Default: null) - идентификатор пакета с данными приложения и устройства

  - **app_version** [String] (Default: null) - версия приложения Express

  - **locale** [String] (Default: null) - локаль текущей сессии

  - **host** [String] - имя хоста с которого пришла команда

- **async_files** [Array[Object]] - метаданные файлов для отложенной обработки

  - **type** [String] - тип файла, один из: "image", "video", "document", "voice"

  - **file** [String] - ссылка на файл

  - **file_mime_type** [String] - mimetype файла

  - **file_name** [String] - имя файла

  - **file_preview** [String] (Default: Null) - ссылка на превью

  - **file_preview_height** [Integer] (Default: Null) - высота превью в px

  - **file_preview_width** [Integer] (Default: Null) - ширина превью в px

  - **file_size** [Integer] - размер файла в байтах

  - **file_hash** [String] - хэш файла

  - **file_encryption_algo** [String] - "stream"

  - **chunk_size** [Integer] - размер чанков

  - **file_id** [UUID] - ID файла

  - **duration** [Integer] (Default: Null) - длительность видео/аудио

  - **caption** [String] (Default: Null) - подпись под файлом

- **bot_id** [UUID] - идентификатор бота в системе Express

- **proto_version** [Integer] - версия протокола (BotX -> Bot) используемая при отправке команды

- **entities** [Array[Object]] - особые сущности переданные в сообщении. Например: меншны, хэштеги, ссылки, форварды

  - **type** [String] - тип сущности

  - **data** [Object] - объект с данными сущности

Пример запроса:

```
{
  "sync_id": "a465f0f3-1354-491c-8f11-f400164295cb",
  "command": {
    "body": "/doit #6",
    "command_type": "user",
    "data": {},
    "metadata": {"account_id": 94}
  },
  "attachments": [
    {
      "type": "image",
      "data": {
        "content": "data:image/jpg;base64,eDnXAc1FEUB0VFEFctII3lRlRBcetROeFfduPmXxE/8=",
        "file_name": "image.jpg",
      }
    }
  ],
  "async_files": [],
  "from": {
    "user_huid": "ab103983-6001-44e9-889e-d55feb295494",
    "group_chat_id": "8dada2c8-67a6-4434-9dec-570d244e78ee",
    "ad_login": "example_login",
    "ad_domain": "example.com",
    "username": "Bob",
    "is_admin": true,
    "is_creator": true,
    "chat_type": "group_chat",
    "host": "cts.ccteam.ru",
    "app_version": "1.21.11",
    "device": "Chrome 92.0",
    "device_software": "macOS 10.15.7",
    "device_meta": {
      "permissions": {"microphone": true, "notifications": true},
      "pushes": false,
      "timezone": "Europe/Samara"
    },
    "platform": "web",
    "locale": "en",
    "manufacturer": "Google",
    "platform_package_id": "ru.unlimitedtech.express"
  },
  "bot_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
  "proto_version": 4,
  "entities": [
    {
      "type": "mention",
      "data": {
        "mention_type": "contact",
        "mention_id": "c06a96fa-7881-0bb6-0e0b-0af72fe3683f",
        "mention_data": {
          "user_huid": "ab103983-6001-44e9-889e-d55feb295494",
          "name": "Вася Иванов",
          "conn_type": "cts"
        }
      }
    }
  ]
}
```

**HTTP-статус ответа**

202

**Параметры ответа**

Успешный ответ:

- Не имеет значения

Ошибка (Бот отключен):

- reason - причина ошибки. Значение должно быть `bot_disabled`;

- error_data - объект с данными об ошибке. Должен включать в себя ключ `status_message`.

**Пример ответа**

Успешный ответ:

- http status: 202

```
{
  "result": "accepted"
}
```

Ошибка (Бот отключен):

- http status: 503

```
{
  "reason": "bot_disabled",
  "error_data": {"status_message": "please stand by"},
  "errors": []
}
```

## Системные команды

BotX отправляет системные команды после некоторых операций.

Системные команды имеют зарезервированное пространство имен `system:`.

Системные команды отправляются от лица сервера и в них отсутствуют данные, относящиеся к конкретному пользователю и его устройству, например:

- open_id_access_token

- pds_token

- ad_login, ad_domain, device, locale

- и пр.

### chat_created

Событие отправляется при создание чата.

Название: **system:chat_created**.

Параметры: см. пункт "Обработка команд".

**Вложенные данные** (поле `command["data"]`):

- **group_chat_id** [UUID5] - uuid чата

- **chat_type** [String] - тип чата

- **name** [String] - имя чата

- **creator** [UUID5] - huid создателя чата

- **members** [Array[Object]]

  - **huid** [UUID5] - huid участника

  - **name** [String] - имя участника

  - **user_kind** [String] - тип участника

  - **admin** [Boolean] - является/не является админом

Пример:

```
{
  "sync_id": "a465f0f3-1354-491c-8f11-f400164295cb",
  "command": {
    "body": "system:chat_created",
    "data": {
      "group_chat_id": "8dada2c8-67a6-4434-9dec-570d244e78ee",
      "chat_type": "group_chat",
      "name": "Meeting Room",
      "creator": "ab103983-6001-44e9-889e-d55feb295494",
      "members": [
        {
          "huid": "ab103983-6001-44e9-889e-d55feb295494",
          "name": "Bob",
          "user_kind": "user",
          "admin": true
        },
       {
          "huid": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
          "name": "Funny Bot",
          "user_kind": "botx",
          "admin": false
        }
      ]
    },
    "command_type": "system",
    "metadata": {}
  },
  "async_files": [],
  "attachments": [],
  "entities": [],
  "from": {
    "user_huid": null,
    "group_chat_id": "8dada2c8-67a6-4434-9dec-570d244e78ee",
    "ad_login": null,
    "ad_domain": null,
    "username": null,
    "chat_type": "group_chat",
    "manufacturer": null,
    "device": null,
    "device_software": null,
    "device_meta": {},
    "platform": null,
    "platform_package_id": null,
    "is_admin": null,
    "is_creator": null,
    "app_version": null,
    "locale": "en",
    "host": "cts.ccteam.ru"
  },
  "bot_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
  "proto_version": 4,
  "source_sync_id": null
}
```

### added_to_chat

Событие отправляется при добавление участников в чат.

Название: **system:added_to_chat**.

Параметры: см. пункт "Обработка команд".

**Вложенные данные** (поле `command["data"]`):

- **added_members** [Array[UUID5]] - список добавленных мемберов

Пример:

```
{
  "sync_id": "a465f0f3-1354-491c-8f11-f400164295cb",
  "command": {
    "body": "system:added_to_chat",
    "data": {
      "added_members": [
        "ab103983-6001-44e9-889e-d55feb295494",
        "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4"
      ]
    },
    "command_type": "system",
    "metadata": {}
  },
  "async_files": [],
  "attachments": [],
  "entities": [],
  "from": {
    "user_huid": null,
    "group_chat_id": "8dada2c8-67a6-4434-9dec-570d244e78ee",
    "ad_login": null,
    "ad_domain": null,
    "username": null,
    "chat_type": "group_chat",
    "manufacturer": null,
    "device": null,
    "device_software": null,
    "device_meta": {},
    "platform": null,
    "platform_package_id": null,
    "is_admin": null,
    "is_creator": null,
    "app_version": null,
    "locale": "en",
    "host": "cts.ccteam.ru"
  },
  "bot_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
  "proto_version": 4,
  "source_sync_id": null
}
```

### deleted_from_chat

Событие отправляется при удалении администратором участников чата.

Название: **system:deleted_from_chat**.

Параметры: см. пункт "Обработка команд".

**Вложенные данные** (поле `command["data"]`):

- **deleted_members** [Array[UUID5]] - список идентификаторов удаленных пользователей

Пример:

```
{
  "sync_id": "a465f0f3-1354-491c-8f11-f400164295cb",
  "command": {
    "body": "system:deleted_from_chat",
    "data": {
      "deleted_members": [
        "ab103983-6001-44e9-889e-d55feb295494",
        "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4"
      ]
    },
    "command_type": "system"
  },
  "async_files": [],
  "attachments": [],
  "entities": [],
  "from": {
    "user_huid": null,
    "group_chat_id": "8dada2c8-67a6-4434-9dec-570d244e78ee",
    "ad_login": null,
    "ad_domain": null,
    "username": null,
    "chat_type": "group_chat",
    "manufacturer": null,
    "device": null,
    "device_software": null,
    "device_meta": {},
    "platform": null,
    "platform_package_id": null,
    "is_admin": null,
    "is_creator": null,
    "app_version": null,
    "locale": "en",
    "host": "cts.ccteam.ru"
  },
  "bot_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
  "proto_version": 4,
  "source_sync_id": null
}
```

### left_from_chat

Событие отправляется при выходе участников из чата.

Название: **system:left_from_chat**.

Параметры: см. пункт "Обработка команд".

**Вложенные данные** (поле `command["data"]`):

- **"left_members"**: [[Array[UUID5]] - список идентификаторов покинувших чат пользователей

Пример:

```
{
  "sync_id": "a465f0f3-1354-491c-8f11-f400164295cb",
  "command": {
    "body": "system:left_from_chat",
    "data": {
      "left_members": [
        "ab103983-6001-44e9-889e-d55feb295494",
        "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4"
      ]
    },
    "command_type": "system",
    "metadata": {}
  },
  "async_files": [],
  "attachments": [],
  "entities": [],
  "from": {
    "user_huid": null,
    "group_chat_id": "8dada2c8-67a6-4434-9dec-570d244e78ee",
    "ad_login": null,
    "ad_domain": null,
    "username": null,
    "chat_type": "group_chat",
    "manufacturer": null,
    "device": null,
    "device_software": null,
    "device_meta": {},
    "platform": null,
    "platform_package_id": null,
    "is_admin": null,
    "is_creator": null,
    "app_version": null,
    "locale": "en",
    "host": "cts.ccteam.ru"
  },
  "bot_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
  "proto_version": 4,
  "source_sync_id": null
}
```

### smartapp_event

Событие отправляется клиентом при взаимодействии со SmartApp-приложением.

Название: **system:smartapp_event**.

Параметры: см. пункт "Обработка команд".

**Вложенные данные** (поле `command["data"]`):

- **ref** [UUID] - уникальный идентификатор запроса

- **smartapp_id** [UUID] - ID SmartApp

- **data** [Object] - пользовательские данные

- **opts** [Object] - опции запроса

- **smartapp_api_version** [Integer] - версия протокола smartapp <-> bot

Пример:

```
{
  "sync_id": "a465f0f3-1354-491c-8f11-f400164295cb",
  "command": {
    "body": "system:smartapp_event",
    "data": {
        "ref": "6fafda2c-6505-57a5-a088-25ea5d1d0364",
        "smartapp_id": "8dada2c8-67a6-4434-9dec-570d244e78ee",
        "data": {
            "url": "example.com",
            "method": "POST",
            "headers": {
                "key": "value"
            },
            "body": "XawLhuWPa8ThvXzGo1R3SlnMxo0R8+H7JRC6Y5UkpHA="
        },
        "opts": {},
        "smartapp_api_version": 1
    },
    "command_type": "system",
    "metadata": {}
  },
  "async_files": [],
  "attachments": [],
  "entities": [],
  "from": {
    "user_huid": "b9197d3a-d855-5d34-ba8a-eff3a975ab20",
    "group_chat_id": "8dada2c8-67a6-4434-9dec-570d244e78ee",
    "ad_login": null,
    "ad_domain": null,
    "username": null,
    "chat_type": "group_chat",
    "manufacturer": null,
    "device": null,
    "device_software": null,
    "device_meta": {},
    "platform": null,
    "platform_package_id": null,
    "is_admin": false,
    "is_creator": false,
    "app_version": null,
    "locale": "en",
    "host": "cts.ccteam.ru"
  },
  "bot_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
  "proto_version": 4,
  "source_sync_id": null
}
```

Пример (вместе с вложением/файлом):

```
{
  "sync_id": "a465f0f3-1354-491c-8f11-f400164295cb",
  "command": {
    "body": "system:smartapp_event",
    "data": {
        "ref": "6fafda2c-6505-57a5-a088-25ea5d1d0364",
        "smartapp_id": "8dada2c8-67a6-4434-9dec-570d244e78ee",
        "data": {
            "url": "example.com",
            "method": "POST",
            "headers": {
                "key": "value"
            },
            "body": "XawLhuWPa8ThvXzGo1R3SlnMxo0R8+H7JRC6Y5UkpHA="
        },
        "opts": {},
        "smartapp_api_version": 1
    },
    "command_type": "system",
    "metadata": {}
  },
  "attachments": [
    {
      "type": "image",
      "data": {
        "content": "data:image/jpg;base64,eDnXAc1FEUB0VFEFctII3lRlRBcetROeFfduPmXxE/8=",
        "file_name": "image.jpg",
      }
    },
    {
      "type": "link",
      "data": {
        "url": "http://ya.ru/xxx",
        "url_title": "Header in link",
        "url_preview": "http://ya.ru/xxx.jpg",
        "url_text": "Some text in link"
      }
    }
  ],
  "async_files": [],
  "entities": [],
  "from": {
    "user_huid": "b9197d3a-d855-5d34-ba8a-eff3a975ab20",
    "group_chat_id": "8dada2c8-67a6-4434-9dec-570d244e78ee",
    "ad_login": null,
    "ad_domain": null,
    "username": null,
    "chat_type": "group_chat",
    "manufacturer": null,
    "device": null,
    "device_software": null,
    "device_meta": {},
    "platform": null,
    "platform_package_id": null,
    "is_admin": false,
    "is_creator": false,
    "app_version": null,
    "locale": "en",
    "host": "cts.ccteam.ru"
  },
  "bot_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
  "proto_version": 4,
  "source_sync_id": null
}
```

Пример (вместе с метаданными файла `async_files`):

```
{
  "sync_id": "a465f0f3-1354-491c-8f11-f400164295cb",
  "command": {
    "body": "system:smartapp_event",
    "data": {
        "ref": "6fafda2c-6505-57a5-a088-25ea5d1d0364",
        "smartapp_id": "8dada2c8-67a6-4434-9dec-570d244e78ee",
        "data": {
            "url": "example.com",
            "method": "POST",
            "headers": {
                "key": "value"
            },
            "body": "XawLhuWPa8ThvXzGo1R3SlnMxo0R8+H7JRC6Y5UkpHA="
        },
        "opts": {},
        "smartapp_api_version": 1
    },
    "command_type": "system",
    "metadata": {}
  },
  "attachments": [],
  "entities": [],
  "async_files": [
    {
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
      "file_id": "8dada2c8-67a6-4434-9dec-570d244e78ee"
    }
  ],
  "from": {
    "user_huid": "b9197d3a-d855-5d34-ba8a-eff3a975ab20",
    "group_chat_id": "8dada2c8-67a6-4434-9dec-570d244e78ee",
    "ad_login": null,
    "ad_domain": null,
    "username": null,
    "chat_type": "group_chat",
    "manufacturer": null,
    "device": null,
    "device_software": null,
    "device_meta": {},
    "platform": null,
    "platform_package_id": null,
    "is_admin": false,
    "is_creator": false,
    "app_version": null,
    "locale": "en",
    "host": "cts.ccteam.ru"
  },
  "bot_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
  "proto_version": 4,
  "source_sync_id": null
}
```

### internal_bot_notification

Событие отправляется чат-ботом при взаимодействии с другими чат-ботами.

Название: **system:internal_bot_notification**.

Параметры: см. пункт "Обработка команд".

**Вложенные данные** (поле `command["data"]`):

- **data** [Object] - пользовательские данные

- **opts** [Object] - опции запроса

Пример:

```
{
  "sync_id": "a465f0f3-1354-491c-8f11-f400164295cb",
  "command": {
    "body": "system:internal_bot_notification",
    "data": {
      "data": {
        "message": "ping",
      },
      "opts": {
        "internal_token": "KyKfLJD1zMjNSJ1cQ4+8Lz"
      }
    },
    "command_type": "system",
    "metadata": {}
  },
  "async_files": [],
  "attachments": [],
  "entities": [],
  "from": {
    "user_huid": "b9197d3a-d855-5d34-ba8a-eff3a975ab20",
    "group_chat_id": "8dada2c8-67a6-4434-9dec-570d244e78ee",
    "ad_login": null,
    "ad_domain": null,
    "username": null,
    "chat_type": "group_chat",
    "manufacturer": null,
    "device": null,
    "device_software": null,
    "device_meta": {},
    "platform": null,
    "platform_package_id": null,
    "is_admin": false,
    "is_creator": false,
    "app_version": null,
    "locale": "en",
    "host": "cts.ccteam.ru"
  },
  "bot_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
  "proto_version": 4,
  "source_sync_id": null
}
```

### cts_login

Событие отправляется при успешном логине пользователя на CTS.

Название: **system:cts_login**.

Параметры: см. пункт "Обработка команд".

**Вложенные данные** (поле `command["data"]`):

- **user_huid** [UUID5] - huid пользователя

- **cts_id** [UUID5] - id сервера, на который был произведен логин

Пример:

```
{
  "sync_id": "a465f0f3-1354-491c-8f11-f400164295cb",
  "command": {
    "body": "system:cts_login",
    "data": {
      "user_huid": "b9197d3a-d855-5d34-ba8a-eff3a975ab20",
      "cts_id": "8dada2c8-67a6-4434-9dec-570d244e78ee"
    },
    "command_type": "system",
    "metadata": {}
  },
  "async_files": [],
  "attachments": [],
  "entities": [],
  "from": {
    "user_huid": null,
    "group_chat_id": null,
    "ad_login": null,
    "ad_domain": null,
    "username": null,
    "chat_type": null,
    "manufacturer": null,
    "device": null,
    "device_software": null,
    "device_meta": {},
    "platform": null,
    "platform_package_id": null,
    "is_admin": null,
    "is_creator": null,
    "app_version": null,
    "locale": "en",
    "host": "cts.ccteam.ru"
  },
  "bot_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
  "proto_version": 4,
  "source_sync_id": null
}
```

### cts_logout

Событие отправляется при успешном выходе пользователя с CTS.

Название: **system:cts_logout**.

Параметры: см. пункт "Обработка команд".

**Вложенные данные** (поле `command["data"]`):

- **user_huid** [UUID5] - huid пользователя

- **cts_id** [UUID5] - id сервера, с которого был произведен выход

Пример:

```
{
  "sync_id": "a465f0f3-1354-491c-8f11-f400164295cb",
  "command": {
    "body": "system:cts_logout",
    "data": {
      "user_huid": "b9197d3a-d855-5d34-ba8a-eff3a975ab20",
      "cts_id": "8dada2c8-67a6-4434-9dec-570d244e78ee"
    },
    "command_type": "system",
    "metadata": {}
  },
  "async_files": [],
  "attachments": [],
  "entities": [],
  "from": {
    "user_huid": null,
    "group_chat_id": null,
    "ad_login": null,
    "ad_domain": null,
    "username": null,
    "chat_type": null,
    "manufacturer": null,
    "device": null,
    "device_software": null,
    "device_meta": {},
    "platform": null,
    "platform_package_id": null,
    "is_admin": null,
    "is_creator": null,
    "app_version": null,
    "locale": "en",
    "host": "cts.ccteam.ru"
  },
  "bot_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
  "proto_version": 4,
  "source_sync_id": null
}
```

### event_edit

Событие отправляется при редактировании сообщения пользователем.

Название: **system:event_edit**.

Параметры: см. пункт "Обработка команд".

**Вложенные данные** (поле `command["data"]`):

- **body** [String] (Optional) - обновленное тело сообщения

Пример:

```
{
  "sync_id": "a465f0f3-1354-491c-8f11-f400164295cb",
  "command": {
    "body": "system:event_edit",
    "data": {
      "body": "Edited"
    },
    "command_type": "system",
    "metadata": {}
  },
  "async_files": [],
  "attachments": [
    {
      "content": "data:image/jpg;base64,eDnXAc1FEUB0VFEFctII3lRlRBcetROeFfduPmXxE/8=",
      "file_name": "image.jpg",
    }
  ],
  "entities": [],
  "from": {
    "user_huid": "b9197d3a-d855-5d34-ba8a-eff3a975ab20",
    "group_chat_id": "8dada2c8-67a6-4434-9dec-570d244e78ee",
    "ad_login": null,
    "ad_domain": null,
    "username": null,
    "chat_type": null,
    "manufacturer": null,
    "device": null,
    "device_software": null,
    "device_meta": {},
    "platform": null,
    "platform_package_id": null,
    "is_admin": null,
    "is_creator": null,
    "app_version": null,
    "locale": "en",
    "host": "cts.ccteam.ru"
  },
  "bot_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
  "proto_version": 4,
  "source_sync_id": null
}
```

### chat_deleted_by_user

Событие отправляется при удалении чата пользователем (чат не удаляется на сервере, только скрывается на устройстве).

Название: **system:chat_deleted_by_user**.

Параметры: см. пункт "Обработка команд".

**Вложенные данные** (поле `command["data"]`):

- **group_chat_id** [UUID] - id удалённого чата

- **user_huid** [UUID] - пользователь, удаливший чат

Пример:

```
{
  "sync_id": "a465f0f3-1354-491c-8f11-f400164295cb",
  "command": {
    "body": "system:chat_deleted_by_user",
    "data": {
      "group_chat_id": "6fa5f1e9-1453-0ad7-2d6d-b791467e382a",
      "user_huid": "37b821f5-a2be-5dc1-b107-87807ce97e56"
    },
    "command_type": "system",
    "metadata": {}
  },
  "async_files": [],
  "attachments": [],
  "entities": [],
  "from": {
    "user_huid": null,
    "group_chat_id": "6fa5f1e9-1453-0ad7-2d6d-b791467e382a",
    "ad_login": null,
    "ad_domain": null,
    "username": null,
    "chat_type": null,
    "manufacturer": null,
    "device": null,
    "device_software": null,
    "device_meta": {},
    "platform": null,
    "platform_package_id": null,
    "is_admin": null,
    "is_creator": null,
    "app_version": null,
    "locale": "en",
    "host": "cts.ccteam.ru"
  },
  "bot_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
  "proto_version": 4,
  "source_sync_id": null
}
```

### conference_created

Событие отправляется при создании конференции с ботом или добавлении бота в существующую.

Название: **system:conference_created**.

Параметры: см. пункт "Обработка команд".

**Вложенные данные** (поле `command["data"]`):

- **call_id** [UUID] - id конференции

Пример:

```
{
  "sync_id": "a465f0f3-1354-491c-8f11-f400164295cb",
  "command": {
    "body": "system:conference_created",
    "data": {
      "call_id": "6fa5f1e9-1453-0ad7-2d6d-b791467e382a"
    },
    "command_type": "system",
    "metadata": {}
  },
  "async_files": [],
  "attachments": [],
  "entities": [],
  "from": {
    "user_huid": null,
    "group_chat_id": "6fa5f1e9-1453-0ad7-2d6d-b791467e382a",
    "ad_login": null,
    "ad_domain": null,
    "username": null,
    "chat_type": "voex_call",
    "manufacturer": null,
    "device": null,
    "device_software": null,
    "device_meta": {},
    "platform": null,
    "platform_package_id": null,
    "is_admin": null,
    "is_creator": null,
    "app_version": null,
    "locale": "en",
    "host": "cts.ccteam.ru"
  },
  "bot_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
  "proto_version": 4,
  "source_sync_id": null
}
```

### conference_deleted

Событие отправляется при удалении конференции с ботом или удалении бота из существующей.

Название: **system:conference_deleted**.

Параметры: см. пункт "Обработка команд".

**Вложенные данные** (поле `command["data"]`):

- **call_id** [UUID] - id конференции

Пример:

```
{
  "sync_id": "a465f0f3-1354-491c-8f11-f400164295cb",
  "command": {
    "body": "system:conference_deleted",
    "data": {
      "call_id": "6fa5f1e9-1453-0ad7-2d6d-b791467e382a"
    },
    "command_type": "system",
    "metadata": {}
  },
  "async_files": [],
  "attachments": [],
  "entities": [],
  "from": {
    "user_huid": null,
    "group_chat_id": "6fa5f1e9-1453-0ad7-2d6d-b791467e382a",
    "ad_login": null,
    "ad_domain": null,
    "username": null,
    "chat_type": "voex_call",
    "manufacturer": null,
    "device": null,
    "device_software": null,
    "device_meta": {},
    "platform": null,
    "platform_package_id": null,
    "is_admin": null,
    "is_creator": null,
    "app_version": null,
    "locale": "en",
    "host": "cts.ccteam.ru"
  },
  "bot_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
  "proto_version": 4,
  "source_sync_id": null
}
```

### call_started

Событие отправляется при начале звонка или конференции с ботом.

Название: **system:call_started**.

Параметры: см. пункт "Обработка команд".

**Вложенные данные** (поле `command["data"]`):

- **call_id** [UUID] - id конференции

- **call_type** [String] - тип звонка (voex_room|voex_call)

Пример:

```
{
  "sync_id": "a465f0f3-1354-491c-8f11-f400164295cb",
  "command": {
    "body": "system:call_started",
    "data": {
      "call_id": "6fa5f1e9-1453-0ad7-2d6d-b791467e382a",
      "call_type": "voex_room"
    },
    "command_type": "system",
    "metadata": {}
  },
  "async_files": [],
  "attachments": [],
  "entities": [],
  "from": {
    "user_huid": null,
    "group_chat_id": "6fa5f1e9-1453-0ad7-2d6d-b791467e382a",
    "ad_login": null,
    "ad_domain": null,
    "username": null,
    "chat_type": "voex_call",
    "manufacturer": null,
    "device": null,
    "device_software": null,
    "device_meta": {},
    "platform": null,
    "platform_package_id": null,
    "is_admin": null,
    "is_creator": null,
    "app_version": null,
    "locale": "en",
    "host": "cts.ccteam.ru"
  },
  "bot_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
  "proto_version": 4,
  "source_sync_id": null
}
```

### call_ended

Событие отправляется при окончании звонка или конференции с ботом.

Название: **system:call_ended**.

Параметры: см. пункт "Обработка команд".

**Вложенные данные** (поле `command["data"]`):

- **call_id** [UUID] - id конференции

- **call_type** [String] - тип звонка (voex_room|voex_call)

Пример:

```
{
  "sync_id": "a465f0f3-1354-491c-8f11-f400164295cb",
  "command": {
    "body": "system:call_ended",
    "data": {
      "call_id": "6fa5f1e9-1453-0ad7-2d6d-b791467e382a",
      "call_type": "voex_room"
    },
    "command_type": "system",
    "metadata": {}
  },
  "async_files": [],
  "attachments": [],
  "entities": [],
  "from": {
    "user_huid": null,
    "group_chat_id": "6fa5f1e9-1453-0ad7-2d6d-b791467e382a",
    "ad_login": null,
    "ad_domain": null,
    "username": null,
    "chat_type": "voex_call",
    "manufacturer": null,
    "device": null,
    "device_software": null,
    "device_meta": {},
    "platform": null,
    "platform_package_id": null,
    "is_admin": null,
    "is_creator": null,
    "app_version": null,
    "locale": "en",
    "host": "cts.ccteam.ru"
  },
  "bot_id": "dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4",
  "proto_version": 4,
  "source_sync_id": null
}
```

## Типы Attachments

### image

Пример:

```
{
  "type": "image",
  "data": {
    "content": "data:image/jpg;base64,eDnXAc1FEUB0VFEFctII3lRlRBcetROeFfduPmXxE/8=",
    "file_name": "image.jpg",
    }
}
```

**Описание data-полей:**

- **content** [String] - содержимое файла в формате data URL + base64 data (RFC 2397)

- **file_name** [String] (Optional) - имя файла

### video

Пример:

```
{
  "type": "video",
  "data": {
    "content": "data:video/mp4;base64,eDnXAc1FEUB0VFEFctII3lRlRBcetROeFfduPmXxE/8=",
    "file_name": "video.mp4",
    "duration": 100
    }
}
```

**Описание data-полей:**

- **content** [String] - содержимое файла в формате data URL + base64 data (RFC 2397)

- **file_name** [String] (Optional) - имя файла

- **duration** [Integer] - длительность воспроизведения

### document

Пример:

```
{
  "type": "document",
  "data": {
    "content": "data:application/vnd.openxmlformats-officedocument.wordprocessingml.document;base64,eDnXAc1FEUB0VFEFctII3lRlRBcetROeFfduPmXxE/8=",
    "file_name": "document.docx",
    }
}
```

**Описание data-полей:**

- **content** [String] - содержимое файла в формате data URL + base64 data (RFC 2397)

- **file_name** [String] (Optional) - имя файла

### voice

Пример:

```
{
  "type": "voice",
  "data": {
    "content": "data:audio/mpeg;base64,eDnXAc1FEUB0VFEFctII3lRlRBcetROeFfduPmXxE/8=",
    "duration": 100
    }
}
```

**Описание data-полей:**

- **content** [String] - содержимое файла в формате data URL + base64 data (RFC 2397)

- **duration** [Integer] - длительность воспроизведения

### location

Пример:

```
{
  "type": "location",
  "data": {
    "location_name": "Центр вселенной",
    "location_address": "Россия, Тверская область",
    "location_lat": 58.04861,
    "location_lng": 34.28833,
    }
}
```

**Описание data-полей:**

- **location_name** [String] - название местоположения

- **location_address** [String] - адрес

- **location_lat** [String] - широта

- **location_lng** [String] - долгота

### contact

Пример:

```
{
  "type": "contact",
  "data": {
    "file_name": "Контакт",
    "contact_name": "Иванов Иван",
    "content": "data:text/vcard;base64,eDnXAc1FEUB0VFEFctII3lRlRBcetROeFfduPmXxE/8="
  }
}
```

**Описание data-полей:**

- **content** [String] - содержимое файла в формате data URL + base64 data (RFC 2397)

- **file_name** [String] - имя файла

- **contact_name** [String] - имя контакта

### link

Пример:

```
{
  "type": "link",
  "data": {
    "url": "http://ya.ru/xxx",
    "url_title": "Header in link",
    "url_preview": "http://ya.ru/xxx.jpg",
    "url_text": "Some text in link"
    }
}
```

**Описание data-полей:**

- **url** [String] - URL ссылки

- **url_title** [String] - заголовок

- **url_preview** [String] - URL изображения для предварительного просмотра ссылки

- **url_text** [String] - текст ссылки

### sticker

Пример:

```
{
  "type": "sticker",
  "data": {
    "id": "ab103983-6001-44e9-889e-d55feb295494",
    "link": "https://cts1dev.ccsteam.ru/uploads/sticker_pack/c06a96fa-7881-0bb6-0e0b-0af72fe3683f/d78fd455d32544a88167d4fe592615b0.png?v=1641977923967",
    "pack": "c06a96fa-7881-0bb6-0e0b-0af72fe3683f",
    "version": 1640076750
    }
}
```

**Описание data-полей:**

- **id** [UUID] - id стикера

- **link** [String] - ссылка на стикер

- **pack** [UUID] - id стикерпака

- **version** [Integer] - версия файла

## Типы Entities

### mention

Пример:

```
{
  "type": "mention",
  "data": {
    "mention_type": "contact",
    "mention_id": "c06a96fa-7881-0bb6-0e0b-0af72fe3683f",
    "mention_data": {
      "user_huid": "ab103983-6001-44e9-889e-d55feb295494",
      "name": "Вася Иванов",
      "conn_type": "cts"
    }
  }
}
```

**Описание data-полей:**

- **mention_type** [String] - тип mention (contact|chat|channel|user|all)

- **mention_id** [UUID] - id mention используемого при подстановке в текст

-
**mention_data** [Object]

для mention_type "all" - {}

для mention_type "contact|user":

  - **user_huid** [UUID] - huid контакта

  - **name** [String] - имя контакта

  - **conn_type** [String] - тип соединения контакта

для mention_type "chat|channel":

  - **group_chat_id** [UUID] - id чата

  - **name** [String] - имя чата

### forward

Пример:

```
{
  "type": "forward",
  "data": {
    "group_chat_id": "918da23a-1c9a-506e-8a6f-1328f1499ee8",
    "sender_huid": "c06a96fa-7881-0bb6-0e0b-0af72fe3683f",
    "forward_type": "chat",
    "source_chat_name": "Simple Chat",
    "source_sync_id": "a7ffba12-8d0a-534e-8896-a0aa2d93a434",
    "source_inserted_at": "2020-04-21T22:09:32.178Z"
  }
}
```

**Описание data-полей:**

- **group_chat_id** [UUID] - id чата откуда переслали сообщение

- **sender_huid** [UUID] - huid автора сообщения

- **forward_type** [String] - chat|channel

- **source_chat_name** [String] (Optional) - имя чата откуда переслали сообщение

- **source_sync_id** [UUID] - sync_id пересылаемого сообщения

- **source_inserted_at** [DateTime] - ts пересылаемого сообщения

### reply

Пример:

```
{
  "type": "reply",
  "data": {
    "source_sync_id": "a7ffba12-8d0a-534e-8896-a0aa2d93a434",
    "sender": "c06a96fa-7881-0bb6-0e0b-0af72fe3683f",
    "body": "все равно документацию никто не читает...",
    "mentions": [],
    "attachment": null,
    "reply_type": "chat",
    "source_group_chat_id": "918da23a-1c9a-506e-8a6f-1328f1499ee8",
    "source_chat_name": "Serious Dev Chat",
  }
}
```

**Описание data-полей:**

- **source_sync_id** [UUID] - sync_id исходного сообщения

- **sender** [UUID] - huid автора сообщения

- **body** [String] - текст исходного сообщения

- **mentions** [Array[Mention]] - меншены исходного сообщения

- **attachment** [Attachment] - вложение исходного сообщения

- **reply_type** [String] - chat|botx|group_chat|channel

- **source_group_chat_id** [UUID] (Optional) - ID чата откуда переслали сообщение

- **source_chat_name** [String] (Optional) - имя чата откуда переслали сообщение
