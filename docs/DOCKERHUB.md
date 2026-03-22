# express-botx

[![GitHub](https://img.shields.io/github/v/release/lavr/express-botx)](https://github.com/lavr/express-botx/releases/latest)
![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)

CLI and HTTP server for sending messages to [eXpress](https://express.ms) corporate messenger via BotX API.

Accepts webhooks from Alertmanager and Grafana, supports async delivery via RabbitMQ/Kafka, works as a CLI tool or HTTP service.

## Features

- **Send messages** from CLI, scripts, CI/CD pipelines
- **API requests** — arbitrary eXpress BotX API calls from CLI with auto-authentication
- **HTTP server** with API for sending messages and receiving webhooks
- **Alertmanager & Grafana** — built-in webhook endpoints for monitoring
- **Async queues** — RabbitMQ or Kafka for reliable delivery
- **Secrets** — environment variables and HashiCorp Vault support
- **Kubernetes-ready** — Docker, Helm chart, binary

## Image variants

| Tag | Base | Description |
|-----|------|-------------|
| `latest`, `<version>` | Alpine 3.21 | Default image |
| `latest-rootless`, `<version>-rootless` | scratch | Minimal image, runs as non-root (UID 65534) |

Both variants are built for `linux/amd64` and `linux/arm64`.

## Quick Start

### Send a message
```bash
docker run --rm -v ./config.yaml:/config.yaml \
  lavr/express-botx send --config /config.yaml "Hello!"
```

### Run HTTP server

```
docker run -d -p 8080:8080 -v ./config.yaml:/config.yaml \
  lavr/express-botx serve --config /config.yaml
```

## HTTP Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/healthz` | Health check |
| `POST` | `/api/v1/send` | Send message |
| `POST` | `/api/v1/alertmanager` | Alertmanager webhook |
| `POST` | `/api/v1/grafana` | Grafana webhook |

All POST endpoints require `Authorization: Bearer <key>`.


## Documentation

Full documentation: [github.com/lavr/express-botx](https://github.com/lavr/express-botx)

## License

MIT
