# express-send

CLI-утилита для отправки сообщений в корпоративный мессенджер eXpress через BotX API.

Работает в стиле Unix: принимает текст из аргумента, файла или stdin, конфигурируется через файл, переменные окружения и флаги.

## Установка

```bash
go install github.com/lavr/express-send@latest
```

Или сборка из исходников:

```bash
git clone https://github.com/lavr/express-send.git
cd express-send
go build -o express-send .
```

## Использование

```bash
# Текст как аргумент
express-send "Сборка #42 прошла успешно"

# Несколько слов без кавычек тоже работают
express-send Deploy finished successfully

# Из файла
express-send --message-from report.txt

# Из stdin (пайплайн)
echo "Deploy OK" | express-send
cat changelog.md | express-send

# Всё через флаги
express-send --host express.company.ru --bot-id UUID --secret KEY --chat-id UUID "Hello"
```

При успехе утилита завершается молча (exit 0). Ошибки выводятся в stderr (exit 1).

## Конфигурация

Параметры загружаются слоями, каждый следующий перекрывает предыдущий:

1. **YAML-файл** (`~/.config/express-send/config.yaml`)
2. **Переменные окружения**
3. **Флаги командной строки**

### Файл конфигурации

По умолчанию: `~/.config/express-send/config.yaml`

```yaml
host: express.company.ru
bot_id: 054af49e-5e18-4dca-ad73-4f96b6de63fa
secret: my-bot-secret
chat_id: 1a2b3c4d-5e6f-7890-abcd-ef1234567890

cache:
  type: file                  # none | file | vault
  file_path: /tmp/tokens.json # только для type: file
  ttl: 3600                   # время жизни токена в секундах
```

Путь к файлу можно указать явно: `--config /path/to/config.yaml`

### Переменные окружения

| Переменная | Описание |
|---|---|
| `EXPRESS_HOST` | Хост сервера eXpress |
| `EXPRESS_BOT_ID` | UUID бота |
| `EXPRESS_SECRET` | Секрет бота |
| `EXPRESS_CHAT_ID` | UUID чата |
| `EXPRESS_CACHE_TYPE` | Тип кэша: `none`, `file`, `vault` |
| `EXPRESS_CACHE_FILE_PATH` | Путь к файлу кэша токенов |
| `EXPRESS_CACHE_VAULT_URL` | URL Vault-сервера |
| `EXPRESS_CACHE_VAULT_PATH` | Путь в Vault KV v2 |
| `EXPRESS_CACHE_TTL` | TTL кэша в секундах |

### Флаги

```
--host          хост сервера eXpress
--bot-id        UUID бота
--secret        секрет бота (литерал, env:VAR или vault:path#key)
--chat-id       UUID целевого чата
--message-from  прочитать сообщение из файла
--config        путь к файлу конфигурации
--no-cache      отключить кэширование токена
```

## Секреты

Значение `--secret` (и поле `secret` в конфиге) поддерживает три формата:

```bash
# Литеральное значение
express-send --secret "my-secret-key" "Hello"

# Из переменной окружения
express-send --secret env:EXPRESS_BOT_SECRET "Hello"

# Из HashiCorp Vault (KV v2)
express-send --secret "vault:secret/data/express#bot_secret" "Hello"
```

Для Vault необходимы переменные `VAULT_ADDR` и `VAULT_TOKEN`.

## Кэширование токенов

Утилита получает токен бота при каждом запуске. Чтобы избежать лишних запросов, токен можно кэшировать.

### Файловый кэш

```yaml
cache:
  type: file
  file_path: ~/.cache/express-send/tokens.json  # опционально, есть значение по умолчанию
  ttl: 3600
```

### Vault кэш

Хранит токены в HashiCorp Vault KV v2:

```yaml
cache:
  type: vault
  vault_url: https://vault.example.com
  vault_path: secret/data/express-send/tokens
  ttl: 3600
```

Требует переменную `VAULT_TOKEN`. Таймаут на операции с Vault — 5 секунд.

### Отключение кэша

```bash
express-send --no-cache "Hello"
```

Или в конфиге: `cache.type: none` (значение по умолчанию).

## Как это работает

1. Загрузка конфигурации (YAML + env + флаги)
2. Чтение сообщения (файл / аргумент / stdin)
3. Разрешение секрета (литерал / env / Vault)
4. Подпись HMAC-SHA256: `HMAC(key=secret, msg=bot_id)` — hex uppercase
5. Получение токена: `GET /api/v2/botx/bots/{bot_id}/token?signature={sig}`
6. Отправка: `POST /api/v4/botx/notifications/direct` с `Authorization: Bearer {token}`
7. При ответе 401 — автоматический повтор с новым токеном (один раз)

## Структура проекта

```
express-send/
  main.go                       # точка входа: флаги, конфиг, основной flow
  internal/
    config/config.go            # Config struct, Load() — YAML/env/flag layering
    secret/secret.go            # Resolve(): литерал / env:VAR / vault:path#key
    auth/auth.go                # BuildSignature(), GetToken()
    token/                      # кэширование токенов
      cache.go                  #   интерфейс Cache + NoopCache
      file.go                   #   файловый кэш
      vault.go                  #   Vault KV v2 кэш
    sender/sender.go            # Send() — POST notification
    input/input.go              # ReadMessage() — arg/file/stdin
```

## Тестирование

```bash
go test ./...
```

Тесты используют `httptest.Server` для HTTP-вызовов, `t.TempDir()` для файловых операций и `t.Setenv()` для переменных окружения.
