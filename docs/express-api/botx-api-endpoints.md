# eXpress BotX API — Обзор эндпоинтов

Базовый URL: `https://<cts_host>/api/...`

Авторизация: заголовок `Authorization: Bearer <token>` (если не указано иное).

---

## 1. Аутентификация

### Получение токена бота

```
GET /api/v2/botx/bots/{bot_id}/token?signature={signature}
```

Авторизация: **не требуется** (Bearer-токен не нужен).

**Параметры:**

| Параметр    | Тип    | Описание |
|-------------|--------|----------|
| `bot_id`    | path   | UUID бота |
| `signature` | query  | HMAC-SHA256 подпись: `UPPER(HEX(HMAC-SHA256(key=secret_key, msg=bot_id)))` |

**Ответ (200):**

```json
{
  "status": "ok",
  "result": "<token_string>"
}
```

Токен имеет неограниченный срок жизни.

**Реализация:** [`internal/auth/auth.go`](../../internal/auth/auth.go)

---

## 2. Отправка сообщений

### Отправка текстового уведомления

```
POST /api/v4/botx/notifications/direct
```

Content-Type: `application/json`

**Тело запроса:**

```json
{
  "group_chat_id": "<UUID чата>",
  "notification": {
    "status": "ok",
    "body": "Текст сообщения"
  }
}
```

**Ответ (200 / 202):**

```json
{
  "status": "ok",
  "result": {
    "sync_id": "<UUID>"
  }
}
```

**Реализация:** [`internal/botapi/client.go` — `SendNotification`](../../internal/botapi/client.go)

---

### Отправка файла

```
POST /api/v4/botx/notifications/direct
```

Content-Type: `application/json`

Используется тот же эндпоинт, что и для текста, но с полем `file`.

**Тело запроса:**

```json
{
  "group_chat_id": "<UUID чата>",
  "notification": {
    "status": "ok",
    "body": "Подпись к файлу (опционально)"
  },
  "file": {
    "file_name": "document.pdf",
    "data": "data:<mime>;base64,<encoded_content>"
  }
}
```

Поле `notification` опционально — можно отправить файл без текстовой подписи.

Поле `data` содержит файл в формате Data URI: `data:<MIME-тип>;base64,<содержимое>`.

**Ответ:** аналогичен текстовому уведомлению (200 / 201 / 202).

**Реализация:** [`internal/botapi/client.go` — `UploadFile`](../../internal/botapi/client.go)

---

## 3. Управление чатами

### Список чатов бота

```
GET /api/v3/botx/chats/list
```

Возвращает все чаты, в которых состоит бот.

**Ответ (200):**

```json
{
  "result": [
    {
      "group_chat_id": "<UUID>",
      "name": "Название чата",
      "description": "Описание",
      "chat_type": "group_chat",
      "members": ["<huid1>", "<huid2>"],
      "shared_history": true
    }
  ]
}
```

**Поля `ChatInfo`:**

| Поле             | Тип      | Описание |
|------------------|----------|----------|
| `group_chat_id`  | string   | UUID чата |
| `name`           | string   | Название чата |
| `description`    | string?  | Описание (может быть `null`) |
| `chat_type`      | string   | Тип чата (`group_chat`, `personal_chat` и т.д.) |
| `members`        | []string | Список HUID участников |
| `shared_history` | bool     | Доступна ли общая история |

**Реализация:** [`internal/botapi/client.go` — `ListChats`](../../internal/botapi/client.go)

---

## 4. Пользователи

### По HUID

```
GET /api/v3/botx/users/by_huid?user_huid={huid}
```

### По email

```
GET /api/v3/botx/users/by_email?email={email}
```

### По AD-логину

```
GET /api/v3/botx/users/by_login?ad_login={login}&ad_domain={domain}
```

**Ответ (200) — одинаковый для всех трёх:**

```json
{
  "status": "ok",
  "result": {
    "user_huid": "<UUID>",
    "name": "Иван Иванов",
    "emails": ["ivan@example.ru"],
    "ad_login": "ivan.ivanov",
    "ad_domain": "example.ru",
    "company": "Компания",
    "company_position": "Должность",
    "department": "Отдел",
    "active": true,
    "user_kind": "cts_user"
  }
}
```

**Поля `UserInfo`:**

| Поле               | Тип      | Описание |
|--------------------|----------|----------|
| `user_huid`        | string   | UUID пользователя |
| `name`             | string   | Имя |
| `emails`           | []string | Список email-адресов |
| `ad_login`         | string   | Логин Active Directory |
| `ad_domain`        | string   | Домен Active Directory |
| `company`          | string   | Компания |
| `company_position` | string   | Должность |
| `department`       | string   | Отдел |
| `active`           | bool     | Активен ли пользователь |
| `user_kind`        | string   | Тип пользователя |

**Реализация:** [`internal/botapi/client.go` — `GetUserByHUID`, `GetUserByEmail`, `GetUserByADLogin`](../../internal/botapi/client.go)

---

## 5. Входящие события (CTS -> Бот)

Платформа eXpress отправляет события на зарегистрированный URL бота методом `POST`. Основные типы:

| Событие           | Описание |
|-------------------|----------|
| `chat_created`    | Пользователь создал персональный чат с ботом |
| Команда (текст)   | Пользователь отправил текстовое сообщение / команду (например, `/help`) |
| Нажатие кнопки    | Пользователь нажал на inline-кнопку (bubble); в payload приходят `source_sync_id`, `metadata`, `data` |

---

## 6. Форматирование сообщений

### Markdown

Тело сообщения (`notification.body`) поддерживает Markdown-разметку:

```
**жирный текст**
```

Пример:

```json
{
  "group_chat_id": "<UUID>",
  "notification": {
    "status": "ok",
    "body": "Задача **Подготовить отчёт** успешно создана"
  }
}
```

### Упоминания (mentions)

В шаблонах сообщений поддерживается тег `<mention>`, который заменяется на упоминание пользователя. Используется, например, в Invite Bot через переменную `GREETING_MESSAGE_TML`.

### Inline-кнопки (Bubbles)

К сообщению можно прикрепить интерактивные кнопки через поле `bubbles`. Каждая кнопка содержит:

| Поле      | Тип    | Описание |
|-----------|--------|----------|
| `command` | string | Команда, отправляемая боту при нажатии |
| `label`   | string | Текст на кнопке |
| `data`    | object | Произвольные данные, передаваемые при нажатии (опционально) |

Кнопки группируются в строки (rows). Пример структуры:

```json
{
  "group_chat_id": "<UUID>",
  "notification": {
    "status": "ok",
    "body": "Выберите действие:"
  },
  "bubbles": [
    [
      {"command": "/approve", "label": "Да"},
      {"command": "/reject", "label": "Нет"}
    ],
    [
      {"command": "/help", "label": "Помощь"}
    ]
  ]
}
```

Каждый вложенный массив — отдельная строка кнопок.

### Metadata

К сообщению можно прикрепить произвольные данные через поле `metadata`. При нажатии на кнопку бот получает обратно:

- `source_sync_id` — UUID исходного сообщения
- `metadata` — данные, прикреплённые к сообщению
- `data` — данные из нажатой кнопки

```json
{
  "group_chat_id": "<UUID>",
  "notification": {
    "status": "ok",
    "body": "Список задач:"
  },
  "bubbles": [[{"command": "/details", "label": "Подробнее", "data": {"task_id": 42}}]],
  "metadata": {"page": 1, "total": 5}
}
```

### Редактирование сообщений

Отправленное сообщение можно отредактировать, используя `sync_id` из ответа. При редактировании обновляются только переданные поля — остальные остаются прежними. Чтобы убрать кнопки, нужно явно передать пустой массив `bubbles: []`.

---

## Сводная таблица

| Метод | Эндпоинт | Назначение |
|-------|----------|------------|
| GET   | `/api/v2/botx/bots/{bot_id}/token?signature={sig}` | Получение токена |
| POST  | `/api/v4/botx/notifications/direct` | Отправка сообщения / файла |
| GET   | `/api/v3/botx/chats/list` | Список чатов бота |
| GET   | `/api/v3/botx/users/by_huid?user_huid={huid}` | Поиск пользователя по HUID |
| GET   | `/api/v3/botx/users/by_email?email={email}` | Поиск пользователя по email |
| GET   | `/api/v3/botx/users/by_login?ad_login={login}&ad_domain={domain}` | Поиск пользователя по AD-логину |
