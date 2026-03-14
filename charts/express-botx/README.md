# express-botx

Helm chart для деплоя [express-botx](https://github.com/lavr/express-botx) — HTTP-сервера для отправки сообщений в корпоративный мессенджер eXpress через BotX API.

Поддерживает вебхуки от Alertmanager и Grafana.

## Установка

```bash
helm install express-botx oci://ghcr.io/lavr/charts/express-botx
```

С кастомными значениями:

```bash
helm install express-botx oci://ghcr.io/lavr/charts/express-botx -f values.yaml
```

## Конфигурация

### Минимальный values.yaml

```yaml
config:
  bots:
    prod:
      host: express.company.ru
      id: "bot-uuid"
      secret: "bot-secret"
  chats:
    ops-alerts: "chat-uuid"
  server:
    listen: ":8080"
    base_path: /api/v1
    api_keys:
      - name: monitoring
        key: "api-key"
    alertmanager:
      default_chat_id: ops-alerts
    grafana:
      default_chat_id: ops-alerts
```

### Параметры

| Параметр | Описание | По умолчанию |
|----------|----------|--------------|
| `replicaCount` | Количество реплик | `1` |
| `image.repository` | Docker-образ | `lavr/express-botx` |
| `image.tag` | Тег образа | `appVersion` из Chart.yaml |
| `image.pullPolicy` | Pull policy | `IfNotPresent` |
| `config.bots` | Боты (host, id, secret) | `{}` |
| `config.chats` | Алиасы чатов (name → UUID) | `{}` |
| `config.cache.type` | Тип кэша: `file`, `vault`, `none` | `file` |
| `config.cache.ttl` | TTL кэша в секундах | `3600` |
| `config.server.listen` | Адрес для прослушивания | `":8080"` |
| `config.server.base_path` | Базовый путь API | `/api/v1` |
| `config.server.api_keys` | API-ключи (`[{name, key}]`) | `[]` |
| `config.server.alertmanager` | Настройки Alertmanager webhook | не задано |
| `config.server.grafana` | Настройки Grafana webhook | не задано |
| `existingSecret` | Имя существующего Secret с `config.yaml` | `""` |
| `service.type` | Тип сервиса | `ClusterIP` |
| `service.port` | Порт сервиса | `80` |
| `ingress.enabled` | Включить Ingress | `false` |
| `ingress.className` | Ingress class | `""` |
| `ingress.hosts` | Список хостов | `[]` |
| `ingress.tls` | TLS-настройки | `[]` |
| `resources.requests.cpu` | CPU request | `50m` |
| `resources.requests.memory` | Memory request | `64Mi` |
| `resources.limits.memory` | Memory limit | `128Mi` |
| `autoscaling.enabled` | Включить HPA | `false` |
| `extraEnv` | Дополнительные переменные окружения | `[]` |

### Секреты

Конфиг монтируется из Kubernetes **Secret** (не ConfigMap), т.к. содержит bot secret и API-ключи.

Для использования существующего секрета вместо генерации из `config.*`:

```yaml
existingSecret: my-express-botx-secret
```

Секрет должен содержать ключ `config.yaml` с полным YAML-конфигом.

### Ingress

```yaml
ingress:
  enabled: true
  className: nginx
  annotations:
    nginx.ingress.kubernetes.io/ssl-redirect: "true"
  hosts:
    - host: express-botx.company.ru
      paths:
        - path: /
          pathType: Prefix
  tls:
    - hosts:
        - express-botx.company.ru
      secretName: express-botx-tls
```

### Vault-интеграция

Для чтения секретов из Vault передайте `VAULT_TOKEN` через `extraEnv`:

```yaml
extraEnv:
  - name: VAULT_TOKEN
    valueFrom:
      secretKeyRef:
        name: vault-token
        key: token

config:
  bots:
    prod:
      host: express.company.ru
      id: "bot-uuid"
      secret: "vault:secret/data/express-botx#bot_secret"
```

## Эндпоинты

| Метод | Путь | Описание |
|-------|------|----------|
| `GET` | `/healthz` | Проверка здоровья (используется для probes) |
| `POST` | `{basePath}/send` | Отправка сообщения |
| `POST` | `{basePath}/alertmanager` | Приём вебхуков от Alertmanager |
| `POST` | `{basePath}/grafana` | Приём вебхуков от Grafana |

## Настройка Alertmanager

```yaml
# alertmanager.yml
receivers:
  - name: express
    webhook_configs:
      - url: http://express-botx.express-botx.svc/api/v1/alertmanager
        send_resolved: true
        http_config:
          bearer_token: "<api-key>"
```

## Настройка Grafana

Contact point → Webhook:
- **URL:** `http://express-botx.express-botx.svc/api/v1/grafana`
- **Authorization:** Bearer `<api-key>`
