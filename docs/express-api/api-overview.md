# Общее описание

## Создание чат-бота

Чат-бот создается в консоли администратора корпоративного сервера.

**Описание полей**

| Параметр | Тип данных | Описание |
| --- | --- | --- |
| id | UUID | Уникальный идентификатор в системе BotX. Используется для идентификации в BotX и приложении чат-бота |
| app_id | string | Уникальный текстовый идентификатор чат-бота |
| url | string | Ссылка на API чат-бота |
| name | string | Имя чат-бота |
| description | string | Описание |
| enabled | boolean | Включен/выключен чат-бот |
| status_message | string | Статус-сообщение |
| secret_key | string | Секретный ключ, генерируется в момент создания чат-бота |
| proto_version | integer | Используемая версия протокола (BotX -> Bot) |

**Пример чат-бота**

- **id**: "dcfa5a7c-7cc4-4c89-b6c0-80325604f9a4"

- **app_id**: "trello"

- **url**: "https://bot.com/api/v1/botx_trello"

- **name**: "Trello Bot"

- **description**: "Бот для работы с Trello-досками"

- **enabled**: true

- **status_message**: "It works!"

- **secret_key**: "secret"

- **proto_version**: 4

## Редактирование чат-бота

Отредактировать свойства чат-бота можно на его странице в консоли администратора.

**Доступные свойства**

-
-
-

-
-
-
-

| Параметр | Тип данных | Описание |
| --- | --- | --- |
| allowed_data | all\|commands\|none(default: commands) | Типы данных, которые может принимать чат-бот:all- чат-бот принимает все сообщения, отправленные в чат;commands- чат-бот принимает сообщение, только если его упомянули/заменшили;none- чат-бот не принимает сообщений |
| allow_chat_creating | true\|false(default: false) | Позволить чат-боту создавать чаты |
| show_in_catalog | true\|false(default:true) | Отображать чат-бота в каталоге чат-ботов |
| communication_availability | corporate\|trust\|local\|all(default: corporate) | Тип пользователей, которые могут взаимодействовать с чат-ботом:corporate- корп. пользователь с любого CTS;trust- корп. пользователь с трастового CTS;local- корп. пользователь с локального CTS;all- любой пользователь (в том числе с RTS) |
| use_botx_ca_cert | true\|false(default: false) | Использовать SSL CA сертификат BotX при отправке запроса к чат-боту.Для использования BotX SSL CA сертификата необходимо загрузить его в панели администратора на странице настроек сервера (/settings/server).Должна быть загружена полная цепочка выдающих серверов, т. е. сертификат следующего вида:-----BEGIN CERTIFICATE-----Intermediate CA-----END CERTIFICATE----------BEGIN CERTIFICATE-----Root CA-----END CERTIFICATE-----Для использования Botx SSL CA сертификата необходимо установить у чат-бота аттрибут “Версия протокола” в значение 3 или выше.После смены сертификата требуется перезагрузка контейнера.После отключения свойства (если чат-бот раннее использовал его) требуется перезагрузка контейнера. |
| use_pds_token | true\|false(default: false) | Еслиtrue, то при каждом обращение к чат-боту генерируется JWT-токен, который подписывается ПДС ключом. Полученный JWT-токен передается при каждой отправке команды/запроса чат-боту в HTTP-заголовке pds_token |
| pds_key | string | Ключ, которым будет подписываться PDS JWT-токен. Ключ представляет из себя валидный Private RSA-ключ. Подпись выполняется с использованием RS256 алгоритма |
| use_open_id_access_token | true\|false(default: false) | Еслиtrue, то при обращение к чат-боту пользователем передается open_id токен пользователя в заголовкеopen_id_access_token |

## Разработка чат-ботов

Перед тем как приступить к разработке чат-бота, используя наши инструменты, ознакомьтесь со следующей документацией:

1. [Что такое чат-боты и SmartApp](../)

2. [BotX API](botx-api/)

3. [Bot API v4](bot-api/)

4. SDK [pybotx](https://github.com/ExpressApp/pybotx) и [async-box](https://github.com/ExpressApp/async-box) (шаблон бота).

5. Пример простого бота на основе `async-box` - [todo-bot](https://github.com/ExpressApp/todo-bot).

При разработке бота вам могут понадобиться виджеты (многие из них уже реализованы в библиотеке [pybotx-widgets](https://github.com/ExpressApp/pybotx-widgets)). Для удобной работы с состоянием пользователя используйте библиотеку [pybotx-fsm](https://github.com/ExpressApp/pybotx-fsm).

### Библиотеки

- [pybotx](https://github.com/ExpressApp/pybotx)

- [async-box](https://github.com/ExpressApp/async-box)

- [pybotx-fsm](https://github.com/ExpressApp/pybotx-fsm)

- [pybotx-smartapp-rpc](https://github.com/ExpressApp/pybotx-smartapp-rpc)

- [pybotx-smart-logger](https://github.com/ExpressApp/pybotx-smart-logger)

- [pybotx-smartapp-smart-logger](https://github.com/ExpressApp/pybotx-smartapp-smart-logger)

- [smartapp-bridge](https://github.com/ExpressApp/smartapp-bridge)

- [smartapp-sdk](https://github.com/ExpressApp/smartapp-sdk)

### Примеры ботов

- [next-feature-bot](https://github.com/ExpressApp/next-feature-bot)

- [todo-bot](https://github.com/ExpressApp/todo-bot)

- [weather-smartapp](https://github.com/ExpressApp/weather-smartapp)

- [next-feature-smartapp](https://github.com/ExpressApp/next-feature-smartapp)

- [next-feature-smartapp-frontend](https://github.com/ExpressApp/next-feature-smartapp-frontend)
