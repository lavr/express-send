# Users API

**Список методов**

| Метод | URL | Описание |
| --- | --- | --- |
| POST | /api/v3/botx/users/by_email | Поиск данных юзеров по их почтам |
| GET | /api/v3/botx/users/by_huid | Поиск данных юзера по его huid |
| GET | /api/v3/botx/users/by_login | Поиск данных юзера по его AD-логину и AD-домену |
| GET | /api/v3/botx/users/by_other_id | Поиск данных юзера по дополнительному идентификатору |
| GET | /api/v3/botx/users/users_as_csv | Получение списка пользователей на CTS |
| PUT | /api/v3/botx/users/update_profile | Обновление профиля юзера |

## Поиск данных юзеров по их почтам

Обрабатывается синхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **emails** [Array[String]] - почты юзеров

Пример параметров запроса:

```
{
    "emails": ["user1@cts.com", "user2@cts.com"]
  }
```

Пример успешного ответа:

```
{
    "status": "ok",
    "result": [
      {
        "user_huid": "6fafda2c-6505-57a5-a088-25ea5d1d0364",
        "ad_login": "ad_user_login1",
        "ad_domain": "cts.com",
        "name": "Bob",
        "company": "Bobs Co",
        "company_position": "Director",
        "department": "Owners",
        "emails": [
          "ad_user@cts.com", "user1@cts.com"
        ],
        "other_id": "some_id1",
        "user_kind": "cts_user",
        "active": true,
        "created_at": "2023-03-26T14:36:08.740618Z",
        "cts_id": "e0140f4c-4af2-5a2e-9ad1-5f37fceafbaf",
        "description": "Director in Owners dep",
        "ip_phone": 1271020,
        "manager": "Alice",
        "office": "SUN",
        "other_ip_phone": null,
        "other_phone": null,
        "public_name": "Bobby",
        "rts_id": "f46440a4-d930-58d4-b3f5-8110ab846ee3",
        "updated_at": "2023-03-26T14:36:08.740618Z",
      },
      {
        "user_huid": "a465f0f3-1354-491c-8f11-f400164295cb",
        "ad_login": "ad_user_login2",
        "ad_domain": "cts.com",
        "name": "Alice",
        "company": "Bobs Co",
        "company_position": "CEO",
        "department": "Owners",
        "emails": [
          "user2@cts.com"
        ],
        "other_id": "some_id2",
        "user_kind": "cts_user",
        "active": true,
        "created_at": "2023-03-28T14:36:08.740618Z",
        "cts_id": "e0140f4c-4af2-5a2e-9ad1-5f37fceafbaf",
        "description": "CEO",
        "ip_phone": 156,
        "manager": null,
        "office": "SUN",
        "other_ip_phone": null,
        "other_phone": null,
        "public_name": "Bobby",
        "rts_id": "f46440a4-d930-58d4-b3f5-8110ab846ee3",
        "updated_at": "2023-03-28T14:36:08.740618Z",
      }
    ]
  }
```

## Поиск данных юзера по его huid

Обрабатывается синхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **user_huid** [UUID] - huid юзера

Пример параметров запроса:

```
?user_huid=6fafda2c-6505-57a5-a088-25ea5d1d0364
```

Пример успешного ответа:

```
{
    "status": "ok",
    "result": {
      "user_huid": "6fafda2c-6505-57a5-a088-25ea5d1d0364",
      "ad_login": "ad_user_login1",
      "ad_domain": "cts.com",
      "name": "Bob",
      "company": "Bobs Co",
      "company_position": "Director",
      "department": "Owners",
      "emails": [
        "ad_user@cts.com", "user1@cts.com"
      ],
      "other_id": "some_id1",
      "user_kind": "cts_user",
      "active": true,
      "created_at": "2023-03-26T14:36:08.740618Z",
      "cts_id": "e0140f4c-4af2-5a2e-9ad1-5f37fceafbaf",
      "description": "Director in Owners dep",
      "ip_phone": 1271020,
      "manager": "Alice",
      "office": "SUN",
      "other_ip_phone": null,
      "other_phone": null,
      "public_name": "Bobby",
      "rts_id": "f46440a4-d930-58d4-b3f5-8110ab846ee3",
      "updated_at": "2023-03-26T14:36:08.740618Z",
    }
  }
```

Данные юзера не найдены (404):

```
{
    "status": "error",
    "reason": "user_not_found",
    "errors": [],
    "error_data": {}
  }
```

## Поиск данных юзера по его AD-логину и AD-домену

Обрабатывается синхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **ad_login** [String] - AD-логин пользователя

- **ad_domain** [String] - AD-домен пользователя

Пример запроса:

```
?ad_login=ad_user_login&ad_domain=cts.com
```

Пример успешного ответа:

```
{
    "status": "ok",
    "result": {
      "user_huid": "6fafda2c-6505-57a5-a088-25ea5d1d0364",
      "ad_login": "ad_user_login1",
      "ad_domain": "cts.com",
      "name": "Bob",
      "company": "Bobs Co",
      "company_position": "Director",
      "department": "Owners",
      "emails": [
        "ad_user@cts.com", "user1@cts.com"
      ],
      "other_id": "some_id1",
      "user_kind": "cts_user",
      "active": true,
      "created_at": "2023-03-26T14:36:08.740618Z",
      "cts_id": "e0140f4c-4af2-5a2e-9ad1-5f37fceafbaf",
      "description": "Director in Owners dep",
      "ip_phone": 1271020,
      "manager": "Alice",
      "office": "SUN",
      "other_ip_phone": null,
      "other_phone": null,
      "public_name": "Bobby",
      "rts_id": "f46440a4-d930-58d4-b3f5-8110ab846ee3",
      "updated_at": "2023-03-26T14:36:08.740618Z",
    }
  }
```

Данные юзера не найдены (404):

```
{
    "status": "error",
    "reason": "user_not_found",
    "errors": [],
    "error_data": {}
  }
```

## Поиск данных юзера по дополнительному идентификатору

Обрабатывается синхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **other_id** [String] - дополнительный идентификатор

Пример параметров запроса:

```
?other_id=some_id
```

Пример успешного ответа:

```
{
    "status": "ok",
    "result": {
      "user_huid": "6fafda2c-6505-57a5-a088-25ea5d1d0364",
      "ad_login": "ad_user_login1",
      "ad_domain": "cts.com",
      "name": "Bob",
      "company": "Bobs Co",
      "company_position": "Director",
      "department": "Owners",
      "emails": [
        "ad_user@cts.com", "user1@cts.com"
      ],
      "other_id": "some_id1",
      "user_kind": "cts_user",
      "active": true,
      "created_at": "2023-03-26T14:36:08.740618Z",
      "cts_id": "e0140f4c-4af2-5a2e-9ad1-5f37fceafbaf",
      "description": "Director in Owners dep",
      "ip_phone": 1271020,
      "manager": "Alice",
      "office": "SUN",
      "other_ip_phone": null,
      "other_phone": null,
      "public_name": "Bobby",
      "rts_id": "f46440a4-d930-58d4-b3f5-8110ab846ee3",
      "updated_at": "2023-03-26T14:36:08.740618Z",
    }
  }
```

Данные юзера не найдены (404):

```
{
    "status": "error",
    "reason": "user_not_found",
    "errors": [],
    "error_data": {}
  }
```

## Получение списка пользователей на CTS

Обрабатывается синхронно.

Список пользователей выдаётся в виде CSV-файла со следующим набором столбцов:

- **HUID** [UUID] - идентификатор пользователя

- **AD Login** [String] - AD-логин (Active Directory) пользователя

- **Domain** [String] - AD-домен пользователя

- **AD E-mail** [String] (Default: null) - AD-email пользователя

- **Name** [String] - имя пользователя

- **Sync source** [String] - источник синхронизации (`ad, admin, email, openid`)

- **Active** [Boolean] - флаг активности пользователя

- **Kind** [String] - тип пользователя (`cts_user, unregistered, botx`)

- **Company** [String] (Default: null) - имя компании пользователя

- **Department** [String] (Default: null) - отдел, в котором работает пользователь

- **Position** [String] (Default: null) - должность пользователя

- **Manager** [String] (Default: null) - руководитель

- **Manager HUID** [String] (Default: null) - HUID руководителя

**Заголовки:**

- authorization - "Bearer <token>"

**Параметры запроса:**

- **cts_user** [Boolean] (Default: true) - включить в результат пользователей с типом `cts_user`

- **unregistered** [Boolean] (Default: true) - включить в результат пользователей с типом `unregistered`

- **botx** [Boolean] (Default: false) - включить в результат пользователей с типом `botx`

Пример параметров запроса:

```
?cts_user=true&unregistered=false
```

Пример успешного ответа:

```
HUID,AD Login,Domain,AD E-mail,Name,Sync source,Active,Kind,Company,Department,Position
dbc8934f-d0d7-4a9e-89df-d45c137a851c,test_user_17,cts.example.com,,test_user_17,ad,true,cts_user,,,
5c1d8ec4-24c9-4dea-af13-075c52444f08,test11,cts.example.com,test11@domain.com,test11 edit,ad,true,cts_user,,,
2ffb8c03-380a-43c5-9caa-ae88d2f9032e,test-5,cts.example.com,test-5@domain.com,test-5,ad,true,cts_user,,,
```

Не выбран ни один из типов пользовалей (400):

```
{
  "status": "error",
  "reason": "no_user_kind_selected",
  "errors": [],
  "error_data": {}
}
```

## Обновление профиля юзера

Обрабатывается синхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **user_huid** [String] - идентификатор юзера

- **name** [String] (Default: skip) - имя юзера

- **public_name** [String] (Default: skip) - публичное имя юзера

- **avatar** [String] (Default: skip) - аватар пользователя в формате `data URL + base64 data (RFC 2397)`

- **company** [String] (Default: skip) - компания

- **company_position** [String] (Default: skip) - должность

- **description** [String] (Default: skip) - описание/заметки

- **department** [String] (Default: skip) - отдел

- **office** [String] (Default: skip) - адрес офиса

- **manager** [String] (Default: skip) - руководитель

Пример параметров запроса:

```
{
    "user_huid": "6fafda2c-6505-57a5-a088-25ea5d1d0364",
    "name": "John Bork",
    "public_name": "Johny B.",
    "avatar": "data:image/png;base64,eDnXAc1FEUB0VFEFctII3lRlRBcetROeFfduPmXxE/8=",
    "company": "Doge Co",
    "company_position": "Chief",
    "description": "Just boss",
    "department": "Commercy",
    "office": "Moscow",
    "manager": "Bob",
  }
```

Пример успешного ответа:

```
{
    "status": "ok",
    "result": true
  }
```

Юзер не найден (404):

```
{
    "status": "error",
    "reason": "user_not_found",
    "errors": [],
    "error_data": {
      "error_description": "User with specified id not found",
      "user_huid": "6fafda2c-6505-57a5-a088-25ea5d1d0364"
    }
  }
```

Неправильные данные для профиля (400):

```
{
    "status": "error",
    "reason": "invalid_profile",
    "errors": [],
    "error_data": {
      "errors": {"field": "invalid"},
      "error_description": "Invalid profile data",
      "user_huid": "6fafda2c-6505-57a5-a088-25ea5d1d0364"
    }
  }
```

Ошибка от сервиса `ad_phonebook` (503):

```
{
    "status": "error",
    "reason": "error_from_ad_phonebook_service",
    "errors": [],
    "error_data": {
      "reason": "some_error",
      "error_description": "AdPhonebook service returns error. Check BotX container logs (level :warn or upper) for more info.",
      "user_huid": "6fafda2c-6505-57a5-a088-25ea5d1d0364"
    }
  }
```

Непредвиденная ошибка (503):

```
{
    "status": "error",
    "reason": "unexpected_error",
    "errors": [],
    "error_data": {
      "error_description": "Got unexpected error. Check BotX container logs (level :error or upper) for more info.",
      "user_huid": "6fafda2c-6505-57a5-a088-25ea5d1d0364"
    }
  }
```
