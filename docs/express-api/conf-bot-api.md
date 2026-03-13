# Conferences Bot API

**Список методов**

| Метод | URL | Описание |
| --- | --- | --- |
| POST | /api/v1/conference_bot/conferences | Создание конференции |
| POST | /api/v1/conference_bot/conferences/<call_id> | Изменение конференции |
| DELETE | /api/v1/conference_bot/conferences/<call_id> | Удаление конференции |

## Общие сведения

Conferences Bot API позволяет разработчику взаимодействовать с конференциями внутри мессенджера путем их создания, редактирования и удаления.

### Форматы дат

В запросах/ответах используется `iso8601` в формате с микросекундами и таймзоной UTC.

Пример: `2022-10-31T10:10:46.275304Z`

### Аутентификация

Аутентификация API должна производиться с использованием [HTTP Basic аутентификации](https://www.rfc-editor.org/rfc/rfc7617) на стороне веб-сервера либо Bearer-аутентификации.

### Формат ответа API

**Успешный запрос:**

- status: <статус результата запроса>

- result: <результат запроса>

```
{
    "status": "ok",
    "result": "data"
}
```

**Неудачный запрос:**

- status: <статус результата запроса>

- reason: <причина ошибки>

- error_data: <объект с данными ошибки>

- errors: <массив ошибок, более подробно раскрывающий причину ошибки>

```
{
    "status": "error",
    "reason": "invalid_attributes",
    "error_data": {"call_id": "uuid"},
    "errors": ["call_id is required and can't be blank"]
}
```

## Создание конференции

**Заголовки:**

- accept - "application/json"

- content-type - "application/json"

**Статус ответа**

201

**Поля запроса:**

- **contact_type** (default: "email") - тип идентификаторов участников (`"email"/"ad_login"`, optional)

- **name** (default: "conference") - имя конференции (optional)

- **members** - список `email/ad_login` участников (required)

- **admins** - список `email/ad_login` админов (required)

- **creator** - `email/ad_login` создателя конференции (required)

- **start_at** (default: now) - время начала (iso8601, optional)

- **end_at** (default: null) - время завершения (iso8601, optional)

- **link**

  - **link_type** - тип ссылки (`public|trusts|corporate|server`, required)

  - **access_code** (default: null) - код доступа в конференцию (optional)

**Поля ответа:**

- **name** - имя

- **call_id** - UUID - уникальный идентификатор конференции

- **members** - список найденых участников

  - **email/ad_login** - `email/ad_login` пользователя (зависит от `contact_type` запроса)

  - **user_huid** - huid пользователя

- **start_at** - `iso8601` начала звонка

- **end_at** - `iso8601` завершения звонка

- **link**

  - **link_id** - ID ссылки на звонок

  - **link_type** - тип ссылки

  - **link** - ссылка на звонок

  - **access_code** - код доступа

**Возможные ошибки**

Дата завершения конференции раньше, чем текущее время (400):

- reason: end_at_earlier_than_now

Дата начала позже, чем дата окончания (400):

- reason: start_at_later_than_end_at

Некорректные параметры (400):

- reason: invalid_attributes

Создатель не найден (400):

- reason: creator_not_found

По `email/ad_login` не было найдено ни одного пользователя (400):

- reason: no_users_found

Несколько пользователей найдено по одному `email` (503):

- reason: multiple_users_found_by_email

Несколько пользователей найдено по одному `ad_login` (503):

- reason: multiple_users_found_by_ad_login

## Изменение конференции

Метод используется для изменения параметров конференции.

**Заголовки:**

- accept - "application/json"

- content-type - "application/json"

**Статус ответа**

200

**Поля запроса:**

- **contact_type** (default: "email") - тип идентификаторов участников (`"email"/"ad_login"`, optional)

- **name** - имя конференции (optional)

- **members** - список `email/ad_login` участников (optional)

- **admins** - список `email/ad_login` админов (optional)

- **actor** - `email/ad_login` админа, изменяющего конференцию (required)

- **start_at** - время начала (iso8601, optional)

- **end_at** - время завершения (iso8601, optional)

- **link**

  - **link_type** - тип ссылки (optional)

  - **access_code** - код доступа в конференцию (optional)

**Поля ответа:**

- **name** - имя

- **members** - список найденых участников

  - **email/ad_login** - `email/ad_login` пользователя (зависит от `contact_type` запроса)

  - **user_huid** - huid пользователя

- **start_at** - `iso8601` начала звонка

- **end_at** - `iso8601` завершения звонка

- **link**

  - **link_id** - ID ссылки на звонок

  - **link_type** - тип ссылки

  - **link** - ссылка на звонок

  - **access_code** - код доступа

**Возможные ошибки**

Дата завершения конференции раньше, чем текущее время (400):

- reason: end_at_earlier_than_now

Дата начала позже, чем дата окончания (400):

- reason: start_at_later_than_end_at

Изменяемая конференция не найдена (404):

- reason: conference_not_found

`actor` не найден (400):

- reason: actor_not_found

Некорректные параметры (400):

- reason: invalid_attributes

По `email/ad_login` не было найдено ни одного пользователя (400):

- reason: no_users_found

`actor` не является администратором конференции (400):

- reason: not_permitted

Несколько пользователей найдено по одному `email` (503):

- reason: multiple_users_found_by_email

Несколько пользователей найдено по одному `ad_login` (503):

- reason: multiple_users_found_by_ad_login

## Удаление конференции

**Статус ответа**

204

**Поля запроса:**

- **actor** - `email/ad_login` админа, удаляющего конференцию (required)

- **contact_type** (default: "email") - тип идентификатора `actor` (`"email"/"ad_login"`, optional)

**Возможные ошибки**

`actor` не найден (400):

- reason: actor_not_found

Удаляемая конференция не найдена (404):

- reason: conference_not_found
