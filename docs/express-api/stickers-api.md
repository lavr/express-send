# Stickers API

**Список методов**

| Метод | URL | Описание |
| --- | --- | --- |
| GET | /api/v3/botx/stickers/packs | Получение списка наборов стикеров |
| GET | /api/v3/botx/stickers/packs/:pack_id | Получение набора стикеров |
| GET | /api/v3/botx/stickers/packs/:pack_id/stickers/:sticker_id | Получение стикера из набора стикеров |
| POST | /api/v3/botx/stickers/packs | Создание набора стикеров |
| POST | /api/v3/botx/stickers/packs/:pack_id/stickers | Добавление стикера в набор стикеров |
| PUT | /api/v3/botx/stickers/packs/:pack_id | Редактирование набора стикеров |
| DELETE | /api/v3/botx/stickers/packs/:pack_id | Удаление набора стикеров |
| DELETE | /api/v3/botx/stickers/packs/:pack_id/stickers/:sticker_id | Удаление стикера из набора стикеров |

## Получение списка наборов стикеров

Обрабатывается синхронно.

Для перемещения по списку используется пагинация.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **user_huid** [UUID] (Default: null) - уникальный идентификатор создателя набора(ов) стикеров

- **limit** [Integer] (Default: null) - максимальное количество наборов в списке

- **after** [String] (Default: null) - base64-строка с мета-информацией для перемещения по списку. Генерируется на стороне сервера. Содержится в ответе на каждый запрос в объекте `pagination` (см. пример успешного ответа).

Пример запроса:

```
?user_huid=84a12e71-3efc-5c34-87d5-84e3d9ad64fd&limit=10&after=ABCEoS5xPvxcNIfVhOPZrWT9AA0ACnVwZGF0ZWRfYXQAf____QAAAAA=
```

Пример успешного ответа (200):

```
{
    "status": "ok",
    "result": {
        "packs": [
             {
               "id": "26080153-a57d-5a8c-af0e-fdecee3c4435",
               "name": "Sticker Pack",
               "preview": "/uploads/sticker_pack/26080153-a57d-5a8c-af0e-fdecee3c4435/9df3143975ad4e6d93bf85079fbb5f1d.png?v=1614781425296",
               "public": true,
               "stickers_count": 2,
               "stickers_order": [
                 "a998f599-d7ac-5e04-9fdb-2d98224ce4ff",
                 "25054ac4-8be2-5a4b-ae00-9efd38c73fb7"
               ],
               "inserted_at": "2020-11-28T12:56:43.672163Z",
               "updated_at": "2021-02-18T12:52:31.571133Z",
               "deleted_at": null,
             }
        ],
        "pagination": {
            "after": "ABAmCAFTpX1ajK8O_ezuPEQ1AA0ACnVwZGF0ZWRfYXQAf____gAAAAA="
       }
    }
  }
```

## Получение набора стикеров

Обрабатывается синхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **pack_id** [UUID] - уникальный идентификатор набора стикеров

Пример запроса:

```
/api/v3/botx/stickers/packs/26080153-a57d-5a8c-af0e-fdecee3c4435
```

Пример успешного ответа (200):

```
{
    "status": "ok",
    "result": {
      "id": "26080153-a57d-5a8c-af0e-fdecee3c4435",
      "name": "Sticker Pack",
      "public": true,
      "preview": "/uploads/sticker_pack/26080153-a57d-5a8c-af0e-fdecee3c4435/b4577728162f4d9ea2b35f25f9f0dcde.png?v=1626137130775",
      "stickers_order": [
        "528c3953-5842-5a30-b2cb-8a09218497bc",
        "75bb24c9-7c08-5db0-ae3e-085929e80c54"
      ],
      "stickers": [
        {
          "id": "528c3953-5842-5a30-b2cb-8a09218497bc",
          "emoji": "😀",
          "link": "/uploads/sticker_pack/26080153-a57d-5a8c-af0e-fdecee3c4435/e7a73cf1b6164b15bc46f21aacfa734f.png?v=1626137621017",
          "inserted_at": "2020-12-28T12:56:43.672163Z",
          "updated_at": "2020-12-28T12:56:43.672163Z",
          "deleted_at": null
        },
        {
          "id": "75bb24c9-7c08-5db0-ae3e-085929e80c54",
          "emoji": "🤔",
          "link": "/uploads/sticker_pack/26080153-a57d-5a8c-af0e-fdecee3c4435/b4577728162f4d9ea2b35f25f9f0dcde.png?v=1626137130775",
          "inserted_at": "2020-12-28T12:56:43.672163Z",
          "updated_at": "2020-12-28T12:56:43.672163Z",
          "deleted_at": null
        }
      ],
      "inserted_at": "2020-12-28T12:56:43.672163Z",
      "updated_at": "2020-12-28T12:56:43.672163Z",
      "deleted_at": null
    }
  }
```

## Получение стикера из набора стикеров

Обрабатывается синхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **pack_id** [UUID] - уникальный идентификатор набора стикеров

- **sticker_id** [UUID] - уникальный идентификатор стикера в наборе стикеров

Пример запроса:

```
/api/v3/botx/stickers/packs/26080153-a57d-5a8c-af0e-fdecee3c4435/stickers/528c3953-5842-5a30-b2cb-8a09218497bc
```

Пример успешного ответа (200):

```
{
    "status": "ok",
    "result": {
        "id": "528c3953-5842-5a30-b2cb-8a09218497bc",
        "emoji": "😀",
        "link": "/uploads/sticker_pack/26080153-a57d-5a8c-af0e-fdecee3c4435/1a036cec0df042898540824717c5ac32.png?v=1607014055820",
        "preview": "/uploads/sticker_pack/26080153-a57d-5a8c-af0e-fdecee3c4435/c9bfcd8eb4224f7a98de28bdcbffc2b2.png?v=1610438587839"
    }
  }
```

## Создание набора стикеров

Обрабатывается синхронно.

Создает только непубличные наборы стикеров (набор виден только на CTS с ботом).

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **name** [String] - название набора стикеров

- **user_huid** [UUID] (Default: null) - huid создателя набора

Пример запроса:

```
{
    "name": "Sticker Pack"
}
```

Пример успешного ответа (201):

```
{
    "status": "ok",
    "result": {
        "id": "26080153-a57d-5a8c-af0e-fdecee3c4435",
        "name": "Sticker Pack",
        "public": false,
        "preview": null,
        "stickers": [],
        "stickers_order": [],
        "inserted_at": "2021-07-10T00:27:55.616703Z",
        "updated_at": "2021-07-10T00:27:55.616703Z",
        "deleted_at": null
    }
  }
```

## Добавление стикера в набор стикеров

Обрабатывается синхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **emoji** [String] - эмодзи стикера

- **image** [String] - изображение стикера в base64-формате

Пример запроса:

```
{
    "emoji": "🤔",
    "image": "data:image/png;base64,iVBORw0KGgoAAAANSUhE",
}
```

Пример успешного ответа (201):

```
{
    "status": "ok",
    "result": {
        "id": "75bb24c9-7c08-5db0-ae3e-085929e80c54",
        "emoji": "🤔",
        "link": "/uploads/sticker_pack/26080153-a57d-5a8c-af0e-fdecee3c4435/b4577728162f4d9ea2b35f25f9f0dcde.png?v=1626137130775",
        "inserted_at": "2020-12-28T12:56:43.672163Z",
        "updated_at": "2020-12-28T12:56:43.672163Z",
        "deleted_at": null
    }
  }
```

## Редактирование набора стикеров

Обрабатывается синхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **name** [String] - название набора стикеров

- **preview** [UUID] - уникальный идентификатор стикера из набора, выбранного в качестве превью

- **stickers_order** [Array[UUID]] - список идентификаторов стикеров набора в порядке их отображения. Для изменения порядка необходимо передавать весь список идентификаторов, в противном случае параметр игнорируется

Пример запроса:

```
{
    "name": "Sticker Pack 2.0",
    "preview": "75bb24c9-7c08-5db0-ae3e-085929e80c54",
    "stickers_order": ["75bb24c9-7c08-5db0-ae3e-085929e80c54", "528c3953-5842-5a30-b2cb-8a09218497bc"]
}
```

Пример успешного ответа (200):

```
{
    "status": "ok",
    "result": {
      "id": "26080153-a57d-5a8c-af0e-fdecee3c4435",
      "name": "Sticker Pack 2.0",
      "public": true,
      "preview": "/uploads/sticker_pack/26080153-a57d-5a8c-af0e-fdecee3c4435/b4577728162f4d9ea2b35f25f9f0dcde.png?v=1626137130775",
      "stickers_order": [
        "75bb24c9-7c08-5db0-ae3e-085929e80c54",
        "528c3953-5842-5a30-b2cb-8a09218497bc"
      ],
      "stickers": [
        {
          "id": "75bb24c9-7c08-5db0-ae3e-085929e80c54",
          "emoji": "🤔",
          "link": "/uploads/sticker_pack/26080153-a57d-5a8c-af0e-fdecee3c4435/b4577728162f4d9ea2b35f25f9f0dcde.png?v=1626137130775",
          "inserted_at": "2020-12-28T12:56:43.672163Z",
          "updated_at": "2020-12-28T12:56:43.672163Z",
          "deleted_at": null
        },
        {
          "id": "528c3953-5842-5a30-b2cb-8a09218497bc",
          "emoji": "😀",
          "link": "/uploads/sticker_pack/26080153-a57d-5a8c-af0e-fdecee3c4435/e7a73cf1b6164b15bc46f21aacfa734f.png?v=1626137621017",
          "inserted_at": "2020-12-28T12:56:43.672163Z",
          "updated_at": "2020-12-28T12:56:43.672163Z",
          "deleted_at": null
        }
      ],
      "inserted_at": "2020-12-28T12:56:43.672163Z",
      "updated_at": "2021-07-22T13:26:41.562143Z",
      "deleted_at": null
    }
  }
```

## Удаление набора стикеров

Обрабатывается асинхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**

- **pack_id** [UUID] - уникальный идентификатор набора стикеров

Пример запроса:

```
/api/v3/botx/stickers/packs/26080153-a57d-5a8c-af0e-fdecee3c4435
```

Пример успешного ответа (202):

```
{
    "status": "ok",
    "result": "delete_sticker_pack_pushed"
  }
```

## Удаление стикера из набора стикеров

Обрабатывается асинхронно.

**Заголовки:**

- authorization - "Bearer <token>"

- content-type - "application/json"

**Параметры запроса:**
-  **pack_id** [UUID] - уникальный идентификатор набора стикеров
-  **sticker_id** [UUID] - уникальный идентификатор стикера из набора стикеров

Пример запроса:

```
/api/v3/botx/stickers/packs/26080153-a57d-5a8c-af0e-fdecee3c4435/stickers/528c3953-5842-5a30-b2cb-8a09218497bc
```

Пример успешного ответа (202):

```
{
    "status": "ok",
    "result": "delete_sticker_from_pack_pushed"
  }
```
