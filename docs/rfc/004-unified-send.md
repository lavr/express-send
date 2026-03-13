# RFC-004: Унификация команды send

- **Статус:** Draft
- **Дата:** 2026-03-13

## Контекст

Сейчас отправка разделена на две субкоманды:

```bash
express-bot send message "текст"
express-bot send file --caption "подпись" ./report.pdf
```

Но BotX API (`POST /api/v4/botx/notifications/direct`) — это **один** эндпоинт, который принимает текст и файл в одном запросе. Разделение на две команды создаёт искусственное ограничение: нельзя отправить файл с текстом в одном вызове, нельзя использовать opts или metadata.

Полная структура запроса API:

```json
{
  "group_chat_id": "...",
  "notification": {
    "status": "ok",
    "body": "текст сообщения",
    "metadata": {"key": "value"},
    "opts": {"silent_response": true},
    "bubble": [[{"command": "/cmd", "label": "Кнопка"}]],
    "keyboard": [[{"command": "/help", "label": "Помощь"}]],
    "mentions": [{"mention_id": "...", "mention_type": "user", "mention_data": {...}}]
  },
  "file": {
    "file_name": "card.png",
    "data": "data:image/png;base64,..."
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

## Предложение

Заменить `send message` и `send file` на одну команду `send`. Обратная совместимость не требуется.

### Использование

```bash
# Текст (как раньше)
express-bot send "Деплой завершён"
express-bot send --from report.txt
echo "OK" | express-bot send

# Файл
express-bot send --file ./report.pdf
express-bot send --file ./report.pdf --file-name custom-name.pdf

# Файл из stdin
cat image.png | express-bot send --file - --file-name image.png

# Текст + файл (одним запросом)
express-bot send --file ./report.pdf "Отчёт за март"
echo "Отчёт за март" | express-bot send --file ./report.pdf

# Статус
express-bot send --status error "Сборка упала"

# Silent (без push-уведомления получателю)
express-bot send --silent "тихое сообщение"

# Stealth mode (сообщение видно только отправителю)
express-bot send --stealth "видно только мне"

# Force DND (доставить, даже если у получателя DND)
express-bot send --force-dnd "срочно!"

# Metadata (произвольный JSON)
express-bot send --metadata '{"ticket_id": 42}' "Тикет создан"
```

### Флаги

| Флаг | Описание | По умолчанию |
|------|----------|:---:|
| `--file PATH` | Путь к файлу (или `-` для stdin) | — |
| `--file-name NAME` | Имя файла (обязательно при `--file -`) | basename от `--file` |
| `--from PATH` | Прочитать текст сообщения из файла | — |
| `--status STATUS` | Статус уведомления: `ok` или `error` | `ok` |
| `--silent` | Не показывать push-уведомление (`notification.opts.silent_response`) | `false` |
| `--stealth` | Stealth mode — сообщение видно только боту (`opts.stealth_mode`) | `false` |
| `--force-dnd` | Доставить даже в режиме DND (`opts.notification_opts.force_dnd`) | `false` |
| `--no-notify` | Не отправлять уведомление совсем (`opts.notification_opts.send = false`) | — |
| `--metadata JSON` | Произвольный JSON для `notification.metadata` | — |

### Источники текста (приоритет)

1. `--from FILE` — из файла
2. Позиционные аргументы — `express-bot send "Hello"`
3. stdin — `echo "Hello" | express-bot send`

Текст не обязателен, если есть `--file` (файл без подписи).

### Что НЕ включаем в P0

Поля `bubble`, `keyboard`, `mentions` — сложные JSON-структуры, неудобные для CLI-флагов. Их реализация откладывается:

- **bubble/keyboard** — потребуют либо `--bubble-json FILE`, либо отдельный DSL. Решение — в отдельном RFC.
- **mentions** — аналогично, можно добавить `--mention HUID` который подставит mention-синтаксис в body. Отдельный RFC.
- **recipients** — фильтрация получателей внутри группового чата. Редкий use-case, откладываем.

## Изменения

### `internal/botapi/client.go`

Единый метод `Send()` заменяет `SendNotification()` и `UploadFile()`:

```go
type SendRequest struct {
    GroupChatID  string            `json:"group_chat_id"`
    Notification *SendNotification `json:"notification,omitempty"`
    File         *SendFile         `json:"file,omitempty"`
    Opts         *SendOpts         `json:"opts,omitempty"`
}

type SendNotification struct {
    Status   string           `json:"status"`
    Body     string           `json:"body"`
    Metadata json.RawMessage  `json:"metadata,omitempty"`
    Opts     *NotificationOpts `json:"opts,omitempty"`
}

type NotificationOpts struct {
    SilentResponse bool `json:"silent_response,omitempty"`
}

type SendFile struct {
    FileName string `json:"file_name"`
    Data     string `json:"data"` // data:mime;base64,...
}

type SendOpts struct {
    StealthMode      bool              `json:"stealth_mode,omitempty"`
    NotificationOpts *DeliveryOpts     `json:"notification_opts,omitempty"`
}

type DeliveryOpts struct {
    Send     *bool `json:"send,omitempty"`
    ForceDND bool  `json:"force_dnd,omitempty"`
}

func (c *Client) Send(ctx context.Context, req *SendRequest) error
```

Старые методы `SendNotification()` и `UploadFile()` удаляются.

### `internal/cmd/send.go`

Команда `send` без субкоманд. Собирает `SendRequest` из флагов:

```go
func runSend(args []string, deps Deps) error
```

Логика:
1. Парсит флаги + global flags
2. `config.Load()` + `RequireChatID()`
3. Читает текст (from/args/stdin) — опционально
4. Читает файл (--file) — опционально
5. Валидация: нужен хотя бы текст или файл
6. Собирает `SendRequest`, отправляет
7. Retry при 401

### `internal/cmd/sendfile.go`

Удаляется.

### `internal/cmd/cmd.go`

В диспетчере `Run()`: `case "send"` → `runSend()` напрямую (без субкоманд).

## Приоритизация

| Приоритет | Задача | Сложность |
|-----------|--------|-----------|
| P0 | Единый `Send()` в client.go | Средняя |
| P0 | Команда `send` без субкоманд | Средняя |
| P0 | `--file`, `--file-name`, `--from`, `--status` | Низкая |
| P0 | `--silent`, `--stealth`, `--force-dnd`, `--no-notify` | Низкая |
| P0 | `--metadata` (raw JSON) | Низкая |
| P1 | `--bubble-json`, `--keyboard-json` (из файла) | Средняя |
| P2 | `--mention HUID` (синтаксический сахар) | Средняя |

## Файлы

| Действие | Файл |
|----------|------|
| MODIFY | `internal/botapi/client.go` — `SendRequest`, `Send()`, удаление `SendNotification()`/`UploadFile()` |
| MODIFY | `internal/cmd/send.go` — единая команда `send` |
| DELETE | `internal/cmd/sendfile.go` |
| MODIFY | `internal/cmd/cmd.go` — `send` без диспетчера субкоманд |
| MODIFY | `README.md` |

## Проверка

1. `express-bot send "Hello"` — текст
2. `express-bot send --from report.txt` — текст из файла
3. `echo "OK" | express-bot send` — текст из stdin
4. `express-bot send --file ./report.pdf` — файл без текста
5. `express-bot send --file ./report.pdf "Отчёт"` — файл + текст
6. `cat img.png | express-bot send --file - --file-name img.png` — файл из stdin
7. `express-bot send --status error "Fail"` — статус error
8. `express-bot send --silent "тихо"` — silent_response
9. `express-bot send --stealth "скрыто"` — stealth_mode
10. `express-bot send --metadata '{"id":1}' "test"` — metadata
11. `express-bot send` (без текста и файла) — ошибка
12. `express-bot send --file - ` (stdin + terminal + без --file-name) — ошибка
