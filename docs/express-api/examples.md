# Примеры использования BotX API

## Отправка сообщения

**[Метод "Отправка директ-нотификации v4"](../../api/botx-api/notifications-api/#отправка-директ-нотификации-v4)

При отправке сообщения используются UUID цели (личный чат, групповой чат, канал) и содержимое сообщения. Бот должен быть одним из участников чата.

> **Warning**
> Внимание!
> Для отправки сообщения в канал бот должен быть его администратором.

Пример:

```
POST https://{express-server}/api/v4/botx/notifications/direct

{
  "group_chat_id": "dec60c05-77b7-0d78-159e-b4fbee4d48f6",
  "notification": {
    "body": "It works!"
  }
}
```

![Рисунок_4](../development-and-debugging-images/04_development-and-debugging.png)

Пример ответа метода:

```
HTTP/1.1 202 Accepted

{
    "result": {
        "sync_id": "0f75c697-0dd1-58d0-99bf-ecfe765caee4"
    },
    "status": "ok"
}
```

Пример коллбэка:

```
POST https://{bot}/notification/callback

{
    "error_data": {
        "error_description": "Chat with specified id not found",
        "group_chat_id": "5d3dfdc2-86fc-5c0e-a8ac-7eee214d9c41"
    },
    "errors": [],
    "reason": "chat_not_found",
    "status": "error",
    "sync_id": "0f75c697-0dd1-58d0-99bf-ecfe765caee4"
}
```

### Отправка сообщения определённым участникам чата

Процедура будет рассмотрена на конкретным сценарии:

В групповом чате есть сообщение с кнопкой, нажать на которую может только определённая группа пользователей (группы определены внутри бота). Если на кнопку нажмёт человек, у которого нет соответствующих прав, чат-бот должен отправить ему сообщение об ошибки. Такое сообщение будет видно всем и в результате будет создана угроза спам-атаки.

Для решения этой проблемы достаточно указать в получателях сообщения того пользователя, которому оно предназначено. Остальные участники сообщение не увидят.

Пример:

```
POST https://{express-server}/api/v4/botx/notifications/direct

{
  "group_chat_id": "dec60c05-77b7-0d78-159e-b4fbee4d48f6",
  "recipients": ["83fbf1c7-f14b-5176-bd32-ca15cf00d4b7"],
  "notification": {
    "body": "Personal message"
  }
}
```

### Скрытие пользовательских сообщений из истории

Для безопасной передачи чат-боту конфиденциальных данных используется режим конфиденциальности или флаг `silent_message`.

[](https://express.ms/android_tutorial.pdf)
> **Note**
> Примечание
> Подробнее о режиме конфиденциальности можно прочитать в
> руководстве пользователя
> .

Как только пользователь получит сообщение от бота с таким флагом, все последующие сообщения от него будут скрыты. При получении сообщения без данного флага скрытие сообщений отключено.

![Рисунок_5](../development-and-debugging-images/05_development-and-debugging.png)

Пример:

```
POST https://{express-server}/api/v4/botx/notifications/direct

{
  "group_chat_id": "dec60c05-77b7-0d78-159e-b4fbee4d48f6",
  "notification": {
    "body": "Enter your password (don't worry, it will be hidden):",
    "opts": {
      "silent_response": true
    }
  }
}
```

### Отправка сообщения с кнопками

Кнопки в сообщениях бывают двух типов: бабблы (`bubble`) и кнопки клавиатуры (`keyboard`).

![Рисунок_6](../development-and-debugging-images/06_development-and-debugging.png)

В отличие от бабблов, кнопки клавиатуры скрываются после нажатия. Их можно снова раскрыть, нажав на иконку клавиатуры в сообщении (правый верхний угол).

Расположение кнопок в сообщении описывается с помощью двумерного массива (матрица). Т. е. каждый вложенный массив представляет собой ряд кнопок.

При нажатии на кнопку, происходит отправка сообщения боту (сообщение видно в чате). В это сообщение входит:

- команда, встроенная в кнопку (может содержать аргументы, например `/command foo bar`);

- произвольный JSON-объект `data`, также встроенный в кнопку;

- произвольный JSON-объект `metadata`, встроенный в сообщение.

``````
> **Note**
> Примечание
> Объект
> data
> предназначен для хранения информации, индивидуальной для кнопки. Например, ID ответа. Объект
> metadata
> хранит общую для сообщения информацию (например, ID вопроса, в котором выбирается ответ). Если ситуация требует дублировать одинаковую информацию в каждой кнопке, рекомендуется использовать
> metadata
> .

````
> **Warning**
> Внимание!
> data
> и
> metadata
> позволяют хранить произвольный JSON-объект. Так как большие объекты в сообщениях будут тормозить мессенджер, желательно хранить там только идентификаторы объектов в БД. Суммарный размер запроса в BotX не должен превышать 1 M.

Пример создания двух кнопок:

```
POST https://{express-server}/api/v4/botx/notifications/direct

{
  "group_chat_id": "dec60c05-77b7-0d78-159e-b4fbee4d48f6",
  "notification": {
    "body": "The time has come to make a choice, Mr. Anderson:",
    "metadata": {"foo": "bar"},
    "bubble": [
      [
        {
          "command": "/choose blue",
          "label": "Blue pill",
          "data": {"baz": "quux"}
        },
        {
          "command": "/choose red",
          "label": "Red pill"
        }
      ]
    ]
  }
}
```

````
> **Note**
> Примечание
> Кнопки клавиатуры обладают той же функциональностью и структурой, что и бабблы. В пример выше можно поменять ключ
> bubble
> на
> keyboard
> и тип кнопок изменится.

### Скрытое нажатие на кнопку

В большинстве случаев отправка сообщения, дублирующего команду, при нажатии на кнопку неуместна. Это особенно актуально для интерактивных виджетов. Чтобы скрыть команду, используйте `"silent": true` в описании кнопки:

```
POST https://{express-server}/api/v4/botx/notifications/direct

{
  "group_chat_id": "dec60c05-77b7-0d78-159e-b4fbee4d48f6",
  "notification": {
    "body": "Message with silent button",
    "bubble": [
      [
        {
          "command": "/foo",
          "label": "Foo",
          "silent": true,
        }
      ]
    ]
  }
}
```

### Изменение ширины отдельных кнопок

Изменение ширины кнопок может потребоваться, например, для создания чек-листа:

![Рисунок_7](../development-and-debugging-images/07_development-and-debugging.png)

Размер задаётся в виде веса. В данном примере вес кнопок с отметкой составляет 1, а вес кнопок с вариантами – 10.

Пример:

```
POST https://{express-server}/api/v4/botx/notifications/direct

{
  "group_chat_id": "dec60c05-77b7-0d78-159e-b4fbee4d48f6",
  "notification": {
    "body": "Make a choice:",
    "bubble": [
      [
        {
          "command": "/choose 1",
          "label": "☐",
          "opts": {
            "h_size": 1
          }
        },
        {
          "command": "/choose 1",
          "label": "First variant",
          "opts": {
            "h_size": 10
          }
        }
      ],
      [
        {
          "command": "/choose 2",
          "label": "☑",
          "opts": {
            "h_size": 1
          }
        },
        {
          "command": "/choose 2",
          "label": "Second variant",
          "opts": {
            "h_size": 10
          }
        }
      ]
    ]
  }
}
```

### Автовыбор ширины кнопок

По умолчанию чат-бот размещает много однотипных кнопок в одной строке. Результат представлен ниже:

![Рисунок_8](../development-and-debugging-images/08_development-and-debugging.png)

Если удобство или сценарий взаимодействия требуют, чтобы клиент сам распределил кнопки, используется опция `buttons_auto_adjust`.

![Рисунок_9](../development-and-debugging-images/09_development-and-debugging.png)

Пример:

```
POST https://{express-server}/api/v4/botx/notifications/direct

{
  "group_chat_id": "dec60c05-77b7-0d78-159e-b4fbee4d48f6",
  "notification": {
    "body": "Auto-adjusted buttons:",
    "bubble": [
      [
        {"command": "/foo", "label": "01"},
        {"command": "/foo", "label": "02"},
        {"command": "/foo", "label": "03"},
        {"command": "/foo", "label": "04"},
        {"command": "/foo", "label": "05"},
        {"command": "/foo", "label": "06"},
        {"command": "/foo", "label": "07"},
        {"command": "/foo", "label": "08"},
        {"command": "/foo", "label": "09"},
        {"command": "/foo", "label": "10"},
        {"command": "/foo", "label": "11"},
        {"command": "/foo", "label": "12"},
        {"command": "/foo", "label": "13"},
        {"command": "/foo", "label": "14"},
        {"command": "/foo", "label": "15"},
        {"command": "/foo", "label": "16"},
        {"command": "/foo", "label": "17"},
        {"command": "/foo", "label": "18"},
        {"command": "/foo", "label": "19"},
        {"command": "/foo", "label": "20"}
      ]
    ],
    "opts": {
      "buttons_auto_adjust": true
    }
  }
}
```

### Всплывающее уведомление при нажатии на кнопку

![Рисунок_10](../development-and-debugging-images/10_development-and-debugging.png)

Функция вызова всплывающего уведомления в чате при нажатии на кнопку.

Пример:

```
POST https://{express-server}/api/v4/botx/notifications/direct

{
  "group_chat_id": "dec60c05-77b7-0d78-159e-b4fbee4d48f6",
  "notification": {
    "body": "Button with alert:",
    "bubble": [
      [
        {
          "command": "/foo",
          "label": "Foo",
          "opts": {
              "show_alert": true,
              "alert_text": "This is alert!",
              "silent": true
          }
        }
      ]
    ]
  }
}
```

Если не передать `alert_text`, но передать `show_alert: true`, будет использован текст уведомления по-умолчанию:

![Рисунок_11](../development-and-debugging-images/11_development-and-debugging.png)

### Упоминание пользователя

Чтобы пользователь обратил внимание на сообщение в групповом чате, можно использовать упоминание. Упоминание "пробивает" отключенные уведомления в чате. Также, через упоминание можно поделиться контактом пользователя (или бота). С точки зрения пользователя, упоминания делятся на 3 вида:

- упоминание пользователя + упоминание всех в чате (посылает пользователю пуш, что его упомянули в чате);

- упоминание-ссылка на пользователя (для передачи контакта, не посылает пуш об упоминании);

- упоминание-ссылка на групповой чат или канал (если чат или канал закрыт, переход по ссылке ничего не даст).

Упоминание состоит из двух частей:

- данные, которые находятся в свойстве `mentions`;

- ID в [специальном формате](../../api/botx-api/notifications-api/#отправка-директ-нотификации-v4), который встраивается в тело сообщения (закрепляя за ним определённую позицию).

![Рисунок_12](../development-and-debugging-images/12_development-and-debugging.png)

Пример:

```
POST https://{express-server}/api/v4/botx/notifications/direct

{
  "group_chat_id": "dec60c05-77b7-0d78-159e-b4fbee4d48f6",
  "notification": {
    "body": "Mention of myself: @{mention:123e4567-e89b-12d3-a456-426655440000}",
    "mentions": [
      {
        "mention_type": "user",
        "mention_id": "123e4567-e89b-12d3-a456-426655440000",
        "mention_data": {
          "user_huid": "5d3dfdc2-86fc-5c0e-a8ac-7eee214d9c41"
        }
      }
    ]
  }
}
```

### Отправка файла в сообщении

Перед отправкой файл должен быть закодирован в формат [base64](https://ru.wikipedia.org/wiki/Base64) и использовать [data url](https://datatracker.ietf.org/doc/html/rfc2397). Ограничения по отправке файлов доступны в [спецификации платформы](../../api/botx-api/#ограничения-api).

![Рисунок_13](../development-and-debugging-images/13_development-and-debugging.png)

Пример:

```
POST https://{express-server}/api/v4/botx/notifications/direct

{
  "group_chat_id": "dec60c05-77b7-0d78-159e-b4fbee4d48f6",
  "file": {
      "data": "data:text/plain;base64,aHR0cHM6Ly93d3cueW91dHViZS5jb20vd2F0Y2g/dj1kUXc0dzlXZ1hjUQ==",
      "file_name": "hello.txt"
  }
}
```

> **Note**
> Примечание
> Также можно отправлять файл вместе с текстом сообщения, кнопками и т. д.

### Отправка сообщения только в режиме конфиденциальности

Подробнее о режиме конфиденциальности можно прочитать в [руководстве пользователя](https://express.ms/android_tutorial.pdf).

Между включением режима конфиденциальности и его непосредственной активацией существует небольшая задержка. Используйте флаг `stealth_mode`, чтобы гарантировать отправку сообщений только после активации режима конфиденциальности. В этом случае, если на момент отправки сообщения режим конфиденциальности не активирован, приложение выдаст ошибку:

```
POST https://{bot}/notification/callback

{
    "error_data": {
        "bot_id": "5d3dfdc2-86fc-5c0e-a8ac-7eee214d9c41",
        "error_description": "Stealth mode disabled in specified chat",
        "group_chat_id": "dec60c05-77b7-0d78-159e-b4fbee4d48f6"
    },
    "errors": [],
    "reason": "stealth_mode_disabled",
    "status": "error",
    "sync_id": "76f679bb-b215-5d90-beb3-a7e6adb95812"
}
```

Пример:

```
POST https://{express-server}/api/v4/botx/notifications/direct

{
  "group_chat_id": "dec60c05-77b7-0d78-159e-b4fbee4d48f6",
  "notification": {
    "body": "Stealth-mode only message"
  },
  "opts": {
    "stealth_mode": true
  }
}
```

### Пропустить отправку пуш-уведомления

Если вы хотите "тихо" отправить сообщение, не требующего срочного внимания пользователя, вы можете передать флаг, который пропустит отправку пуша на устройство пользователя.

Пример:

```
POST https://{express-server}/api/v4/botx/notifications/direct

{
  "group_chat_id": "dec60c05-77b7-0d78-159e-b4fbee4d48f6",
  "notification": {
    "body": "Message without push"
  },
  "opts": {
    "notification_opts": {
      "send": false
    }
  }
}
```

### Игнорирование отключенных уведомлений в чате

Если сценарий требует привлечь внимание пользователя, используется флаг, "пробивающий" отключенные уведомления. Упоминание пользователя также игнорирует отключенные уведомления.

Пример:

```
POST https://{express-server}/api/v4/botx/notifications/direct

{
  "group_chat_id": "dec60c05-77b7-0d78-159e-b4fbee4d48f6",
  "notification": {
    "body": "Very important message"
  },
  "opts": {
    "notification_opts": {
      "force_dnd": true
    }
  }
}
```

## Отправка внутреннего сообщения другим ботам

**[Метод "Отправка внутренней бот нотификации"](../../api/botx-api/notifications-api/#отправка-внутренней-бот-нотификации)

Чат-боты не могут обмениваться обычными сообщениями. Однако, есть внутренние сообщения, которые можно использовать для взаимодействий между чат-ботами. В этом случае мессенджер выступает в качестве шины сообщений.

Важные детали о внутренних сообщениях:

-
внутренние сообщения невидимы для пользователей и нигде не сохраняются;

-
чат-бот может отправить внутреннее сообщение только в групповой чат, где должны быть собраны получатели;

-
если при отправке внутреннего сообщения не нашлось ни одного получателя, метод отправки вернёт ошибку.

Отправка внутренних сообщений использует callback (подробнее см. [отправка сообщений](#отправка-сообщения)).

## Редактирование сообщений

**[Метод "Редактирование события"](../../api/botx-api/events-api/#редактирование-события)

Чат-боту доступна функция редактирования своих сообщений. Такая функциональность позволяет делать интерактивные виджеты с кнопками, которые обновляются при нажатии. В отличие от обновления сообщения пользователем, обновлённое сообщение чат-бота не получает статус "изменено".

![Рисунок_14](../development-and-debugging-images/14_development-and-debugging.png)

Для обновления сообщения используется его `sync_id` и часть, которая должна быть обновлена.

Пример:

```
POST https://{express-server}/api/v3/botx/events/edit_event

{
    "sync_id": "cfeabf62-458d-559a-864a-0b0dd5b03cc5",
    "payload": {
        "body": "Edited!"
    }
}
```

![Рисунок_15](../development-and-debugging-images/15_development-and-debugging.png)

## Ответ на сообщение (reply)

**[Метод "Ответ на сообщение (reply)"](../../api/botx-api/events-api/#ответ-на-сообщение-reply)

![Рисунок_16](../development-and-debugging-images/16_development-and-debugging.png)

Для отправки reply используется `sync_id` сообщения, к которому прикрепится ответ, и содержимое ответа.

Пример:

```
POST https://{express-server}/api/v3/botx/events/reply_event

{
    "source_sync_id": "72315b74-0049-5304-b5a6-f2a037d48d3e",
    "reply": {
        "body": "Replied!"
    }
}
```

## Получение статуса сообщения

**[Метод "Получение статуса сообщения"](../../api/botx-api/events-api/#получение-статуса-сообщения)

Чтобы получить сведения о том, кто получи и\или прочитал сообщение, используется проверка статуса сообщения. Главное условие для правильного срабатывания метода - устройство должно быть подключено к сети Интернет и приложение должно быть как минимум запущено в фоне:

![Рисунок_17](../development-and-debugging-images/17_development-and-debugging.png)

Пример:

```
GET https://{express-server}/api/v3/botx/events/eaf8f559-13cc-5b6a-8ad5-9a4eba4ab9e4/status
```

```
HTTP/1.1 200 OK

{
    "result": {
        "group_chat_id": "1a2aedff-602b-5e23-b40d-fe15612a2de6",
        "read_by": [
            {
                "read_at": "2022-02-24T11:02:40.699118Z",
                "user_huid": "83fbf1c7-f14b-5176-bd32-ca15cf00d4b7"
            }
        ],
        "received_by": [
            {
                "received_at": "2022-02-24T11:02:40.697361Z",
                "user_huid": "83fbf1c7-f14b-5176-bd32-ca15cf00d4b7"
            }
        ],
        "sent_to": [
            "83fbf1c7-f14b-5176-bd32-ca15cf00d4b7",
            "a61be545-e15f-5a09-bfb8-beb57fbf4233"
        ]
    },
    "status": "ok"
}
```

## Отправка "печатает…"

**[Метод "Отправка typing"](../../api/botx-api/events-api/#отправка-typing)

Чат-боту доступно специальное событие, предназначенное для имитации статуса набора текста (такой статус появляется в заголовке чата, если пользователь набирает сообщение).

![Рисунок_18](../development-and-debugging-images/18_development-and-debugging.png)

Пример:

```
POST https://{express-server}/api/v3/botx/events/typing

{
  "group_chat_id": "dec60c05-77b7-0d78-159e-b4fbee4d48f6"
}
```

Чтобы остановить отправку "печатает…" нужно использовать [метод "Отправка stop_typing"](../../api/botx-api/events-api/#отправка-stop_typing).

```
POST https://{express-server}/api/v3/botx/events/stop_typing

{
  "group_chat_id": "dec60c05-77b7-0d78-159e-b4fbee4d48f6"
}
```

## Получение списка чатов бота

**[Метод "Получение списка чатов бота"](../../api/botx-api/chats-api/#получение-списка-чатов-бота)

Для реализации сценария может потребоваться список всех чатов конкретного чат-бота. Например, поиск чата с определённым пользователем. Для решения данных задач рекомендуется использовать БД бота.

Пример:

```
GET https://{express-server}/api/v3/botx/chats/info?group_chat_id=15fca806-d73b-0bc5-31b6-66e26676ca43
```

```
HTTP/1.1 200 OK

{
    "result": {
        "chat_type": "chat",
        "creator": "48c155c4-51c7-57cb-991a-180c473b5602",
        "description": null,
        "group_chat_id": "15fca806-d73b-0bc5-31b6-66e26676ca43",
        "inserted_at": "2022-03-11T09:57:37.237426Z",
        "members": [
            {
                "admin": false,
                "server_id": "4230a78e-1305-5a87-a9bc-fe6eaaa2247a",
                "user_huid": "5d3dfdc2-86fc-5c0e-a8ac-7eee214d9c41",
                "user_kind": "botx"
            },
            {
                "admin": true,
                "server_id": "4230a78e-1305-5a87-a9bc-fe6eaaa2247a",
                "user_huid": "48c155c4-51c7-57cb-991a-180c473b5602",
                "user_kind": "cts_user"
            }
        ],
        "name": "Personal Chat",
        "updated_at": "2022-03-11T09:57:37.237426Z",
        "shared_history": False,
    },
    "status": "ok"
}
```

## Получение информации о чате

**[Метод "Получение информации о чате"](../../api/botx-api/chats-api/#получение-информации-о-чате)

Более подробный вариант предыдущего метода, раскрывающий больше информации об участниках чата.

Пример:

```
GET https://{express-server}/api/v3/botx/chats/info?group_chat_id=15fca806-d73b-0bc5-31b6-66e26676ca43
```

```
HTTP/1.1 200 OK

{
    "result": {
        "chat_type": "chat",
        "creator": "48c155c4-51c7-57cb-991a-180c473b5602",
        "description": null,
        "group_chat_id": "15fca806-d73b-0bc5-31b6-66e26676ca43",
        "inserted_at": "2022-03-11T09:57:37.237426Z",
        "members": [
            {
                "admin": false,
                "server_id": "4230a78e-1305-5a87-a9bc-fe6eaaa2247a",
                "user_huid": "5d3dfdc2-86fc-5c0e-a8ac-7eee214d9c41",
                "user_kind": "botx"
            },
            {
                "admin": true,
                "server_id": "4230a78e-1305-5a87-a9bc-fe6eaaa2247a",
                "user_huid": "48c155c4-51c7-57cb-991a-180c473b5602",
                "user_kind": "cts_user"
            }
        ],
        "name": "Personal Chat",
        "updated_at": "2022-03-11T09:57:37.237426Z",
        "shared_history": False,
    },
    "status": "ok"
}
```

## Добавление пользователей в чат

**[Метод "Добавление юзеров в чат"](../../api/botx-api/chats-api/#добавление-юзеров-в-чат)

![Рисунок_19](../development-and-debugging-images/19_development-and-debugging.png)

Пример:

```
POST https://{express-server}/api/v3/botx/chats/add_user

{
  "group_chat_id": "1a2aedff-602b-5e23-b40d-fe15612a2de6",
  "user_huids": ["a61be545-e15f-5a09-bfb8-beb57fbf4233"]
}
```

## Удаление пользователей из чата

**[Метод "Удаление юзеров из чата"](../../api/botx-api/chats-api/#удаление-юзеров-из-чата)

![Рисунок_20](../development-and-debugging-images/20_development-and-debugging.png)

Пример:

```
POST https://{express-server}/api/v3/botx/chats/remove_user

{
  "group_chat_id": "1a2aedff-602b-5e23-b40d-fe15612a2de6",
  "user_huids": ["a61be545-e15f-5a09-bfb8-beb57fbf4233"]
}
```

## Повышение участников чата до администраторов

**[Метод "Добавление администратора в чат"](../../api/botx-api/chats-api/#добавление-администратора-в-чат)

![Рисунок_21](../development-and-debugging-images/21_development-and-debugging.png)

Пример:

```
POST https://{express-server}/api/v3/botx/chats/add_admin

{
  "group_chat_id": "1a2aedff-602b-5e23-b40d-fe15612a2de6",
  "user_huids": ["a61be545-e15f-5a09-bfb8-beb57fbf4233"]
}
```

## Включение режима конфиденциальности в чате

**[Метод "Включение стелс-режима в чате"](../../api/botx-api/chats-api/#включение-стелс-режима-в-чате)

[](https://express.ms/android_tutorial.pdf)
> **Note**
> Примечание
> Подробнее о режиме конфиденциальности можно прочитать в
> руководстве пользователя
> .

![Рисунок_22](../development-and-debugging-images/22_development-and-debugging.png)

Пример:

```
POST https://{express-server}/api/v3/botx/chats/stealth_set

{
  "group_chat_id": "dec60c05-77b7-0d78-159e-b4fbee4d48f6",
  "burn_in": 10,
  "expire_in": 30,
  "disable_web": false
}
```

## Отключение режима конфиденциальности в чате

**[Метод "Отключение стелс-режима в чате"](../../api/botx-api/chats-api/#отключение-стелс-режима-в-чате)

[](https://express.ms/android_tutorial.pdf)
> **Note**
> Примечание
> Подробнее о режиме конфиденциальности можно прочитать в
> руководстве пользователя
> .

Пример:

```
POST https://{express-server}/api/v3/botx/chats/stealth_disable

{
    "group_chat_id": "dec60c05-77b7-0d78-159e-b4fbee4d48f6"
}
```

## Создание чата

**[Метод "Создание чата"](../../api/botx-api/chats-api/#создание-чата)

Перед эксплуатацией метода создания чата (личный, групповой, канал) нужно учитывать следующие нюансы:

- предварительно у чат-бота в панели администратора активируется параметр **allow_chat_creating** (параметр дает доступ к методу создания чата конкретному чат-боту);

- создателем чата (и администратором) будет являться бот, инициировавший запрос;

- чат-бот добавляется в список участников чата, даже если он явно не указан в списке при отправке запроса;

- опция `shared_history` позволяет создать чат с серверным ключом. Благодаря этой опции пользователи чата видят предыдущую переписку;

- если личный чат с пользователем уже существует, при попытке создать его снова сообщения об ошибки не будет. Метод вернет id существующего чата.

- после удаления личного чата с чат-ботом, данные о чате все равно остаются на сервере. Чат-бот может обращаться к этому чату, не прибегая повторно к методу создания чата.

![Рисунок_23](../development-and-debugging-images/23_development-and-debugging.png)

Пример:

```
POST https://{express-server}/api/v3/botx/chats/create

{
    "chat_type": "group_chat",
    "name": "Test chat",
    "members": ["83fbf1c7-f14b-5176-bd32-ca15cf00d4b7"]
}
```

## Закрепление сообщения в чате

**[Метод "Закрепление сообщения в чате"](../../api/botx-api/chats-api/#закрепление-сообщения-в-чате)

![Рисунок_24](../development-and-debugging-images/24_development-and-debugging.png)

Чат-бот может закрепить сообщение в чате в следующих случаях:

- если сообщение еще не закреплено кем-то другим;

- если чат-бот является администратором чата. В таком случае статус сообщение (закреплено\не закреплено) не имеет значения.

Для закрепления сообщения в канале чат-бот должен быть администратором вне зависимости от статуса сообщения (закреплено  не закреплено).

Пример:

```
POST https://{express-server}/api/v3/botx/chats/pin_message

{
  "chat_id": "dec60c05-77b7-0d78-159e-b4fbee4d48f6",
  "sync_id": "ae7bee6a-b47c-5dc8-9a43-5ec490b4f281"
}
```

## Открепление сообщения в чате

**[Метод "Открепление сообщения в чате"](../../api/botx-api/chats-api/#открепление-сообщения-в-чате)

![Рисунок_25](../development-and-debugging-images/25_development-and-debugging.png)

Чат-бот может открепить любое сообщение, если является администратором чата.

Пример:

```
POST https://{express-server}/api/v3/botx/chats/unpin_message

{
  "chat_id": "dec60c05-77b7-0d78-159e-b4fbee4d48f6"
}
```

## Поиск пользователя

-
**[Метод "Поиск данных юзеров по его почтам"](../../api/botx-api/users-api/#поиск-данных-юзеров-по-их-почтам)

-
**[Метод "Поиск данных юзера по его huid"](../../api/botx-api/users-api/#поиск-данных-юзера-по-его-huid)

-
**[Метод "Поиск данных юзера по его AD-логину и AD-домену"](../../api/botx-api/users-api/#поиск-данных-юзера-по-его-ad-логину-и-ad-домену)

-
**[Метод "Поиск данных юзера по дополнительному идентификатору"](../../api/botx-api/users-api/#поиск-данных-юзера-по-дополнительному-идентификатору)

Для поиска пользователя используются следующие атрибуты:

1.
`huid` (human id). Это уникальный идентификатор пользователя в мессенджере.
   Пример:

```
GET https://{express-server}/api/v3/botx/users/by_huid?user_huid=83fbf1c7-f14b-5176-bd32-ca15cf00d4b7

{
    "result": {
        "ad_domain": "ccsteam.ru",
        "ad_login": "********",
        "company": "Express, Unlimited Production",
        "company_position": "********",
        "department": "********",
        "emails": [
            "********@ccsteam.ru"
        ],
        "name": "********",
        "user_huid": "83fbf1c7-f14b-5176-bd32-ca15cf00d4b7",
        "user_kind": "cts_user"
    },
    "status": "ok"
}
```

2.
`ad` (логин и домен из Active Directory)

```
GET https://{express-server}/api/v3/botx/users/by_login?ad_login=********&ad_domain=********
```

> **Note**
> Примечание
> У всех методов поиска пользователя ответ имеет одинаковый формат.

3.
`email`:

```
GET https://{express-server}/api/v3/botx/users/by_email?email=********
```

4.
`other_id`:

```
GET https://{express-server}/api/v3/botx/users/by_other_id?other_id=********
```

## Обновление профиля пользователя

**[Метод "Обновление профиля юзера"](../../api/botx-api/users-api/#обновление-профиля-юзера)

Чат-бот может менять информацию о пользователе, которая отображается в его профиле в eXpress. Обновление происходит по uuid пользователя. Обновить можно как все поля в профиле, так и какие-то определенные. Пример:

```
PUT https://{express-server}/api/v3/botx/users/update_profile

{
  "user_huid": "9b0db6fb-115c-57db-ba68-8329f990b083",
  "department": "New awesome department"
}
```

## Получение списка пользователей на CTS

**[Метод "Получение списка пользователей на CTS"](../../api/botx-api/users-api/#получение-списка-пользователей-на-cts)

Чат-бот может вернуть список пользователей на своем CTS в формате CSV. Можно так же указать, каких пользователей включать в список: `cts_user`, `unregistered`, `botx`. Допускается выбрать как один тип, так и все возможные.

```
GET https://{express-server}/api/v3/botx/users/users_as_csv?cts_user=true&unregistered=false

HUID,AD Login,Domain,AD E-mail,Name,Sync source,Active,Kind,Company,Department,Position
dbc8934f-d0d7-4a9e-89df-d45c137a851c,test_user_17,cts.example.com,,test_user_17,ad,true,cts_user,,,
5c1d8ec4-24c9-4dea-af13-075c52444f08,test11,cts.example.com,test11@domain.com,test11 edit,ad,true,cts_user,,,
2ffb8c03-380a-43c5-9caa-ae88d2f9032e,test-5,cts.example.com,test-5@domain.com,test-5,ad,true,cts_user,,,
```

## Отправка SmartApp-события

**[Метод "Отправка SmartApp-события"](../../../../smartapps/developer-guide/smartapp-api/#отправка-smartapp-события)

События используются в качестве ответов на запросы от фронтенда SmartApp и как уведомления SmartApp (может быть получено только если SmartApp открыт в настоящий момент). Событие считается запросом, если заполнено поле `ref`. Если нет – событие считается уведомлением.

``[](../../api/botx-api/#длина-запроса)
> **Warning**
> Внимание!
> В поле
> data
> можно передать произвольный JSON-объект, но размер запроса в BotX не должен превышать
> 1 M
> .

Пример:

```
POST https://{express-server}/api/v3/botx/smartapps/event

{
  "ref": "6c4e232e-ce47-4141-9502-998bc0062523",
  "data": {"foo": "bar"},
  "group_chat_id": "8e81b23d-9d01-0934-0d99-9410d621353b",
  "opts": {},
  "smartapp_api_version": 1,
  "smartapp_id": "0d7a43fa-6c4a-5842-b0ab-5e051921e18c"
}
```

## Отправка SmartApp push-уведомления

**[Метод "Отправка SmartApp-нотификации"](../../../../smartapps/developer-guide/smartapp-api/#отправка-smartapp-нотификации)

Push-уведомления используются чтобы обратить внимание пользователя на событие в SmartApp. SmartApp с событием откроется после нажатия на уведомление. Например, это может быть уведомление о назначенной задаче.

SmartApp подгружает информацию о событии сразу после открытия.

```
POST https://{express-server}/api/v3/botx/smartapps/notification

{
    "body": "Test",
    "group_chat_id": "dec60c05-77b7-0d78-159e-b4fbee4d48f6",
    "opts": {},
    "smartapp_api_version": 1,
    "smartapp_counter": 1
}
```

## Получение списка SmartApp на CTS

Чат-бот может вернуть список SmartApp на своем CTS. В результате запроса бот получит список SmartApp и версию этого списка (поле `phonebook_version`), которую можно использовать как необязательный параметр при последующих запросах. Пример:

```
GET https://{express-server}/api/v3/botx/smartapps/list

{
    "result": {
        "phonebook_version": 516853,
        "smartapps": [
            {
                "app_id": "service_desk_smartapp",
                "avatar": "********",
                "avatar_preview": "********",
                "enabled": true,
                "id": "02298156-3469-540e-9670-ef0f83c45f63",
                "name": "Service Desk App"
            },
            {
                "app_id": "email_smart_app",
                "avatar": "********",
                "avatar_preview": "********",
                "enabled": true,
                "id": "ca4a67df-8e09-5db9-8c98-f745e8dde72b",
                "name": "Email App"
            }
        ]
    },
    "status": "ok"
}
```
