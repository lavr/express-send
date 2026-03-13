# Получение статуса и списка команд (v3)

Метод для получения статуса и списка команд чат-бота.

**Заголовки:**

- Authorization - "Bearer {token}"

**Параметры запроса:**

- **bot_id** [UUID] - идентификатор бота в системе Express

- **user_huid** [UUID] - идентификатор юзера, который запрашивает статус

- **ad_login** [String] (Default: null) - логин пользователя в Active Directory

- **ad_domain** [String] (Default: null) - домен пользователя в Active Directory

- **is_admin** [Boolean] (Default: null) - флаг, является ли пользователь админом в чате

- **chat_type** [String] - тип чата, в котором запрашивается статус

Пример параметров запроса:

```
?bot_id=dcfa5a7c-7cc4-4c89-b6c0-80325604f9f4
  &user_huid=ab103983-6001-44e9-889e-d55feb295494
  &ad_login=exlogin
  &ad_domain=exdomain
  &is_admin=true
  &chat_type=chat
```

**HTTP-статус ответа**

200

**Параметры ответа:**

- **status** [String] - "ok"|"error"

- **result** [Object]

  - **enabled** [Boolean] (Default: true) - бот включен/выключен

  - **status_message** [String] (Default: null) - текущий статус бота

  - **commands** [Array[Object]] (Default: []) - список команд бота

    - **description** [String] - описание команды

    - **body** [String] - тело команды

    - **name** [String] - имя команды

Пример ответа:

```
{
    "status": "ok",
    "result": {
      "enabled": true,
      "status_message": "it's work!",
      "commands": [
        {
          "description": "Add Comment",
          "name": "Comment",
          "body": "/comment #text"
        },
        {
          "description": "Move to Doing",
          "name": "Doing",
          "body": "/doing"
        },
        {
          "description": "Move to Doit",
          "name": "Doit",
          "body": "/doit"
        }
      ]
    }
  }
```
