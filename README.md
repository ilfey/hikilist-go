# Golang API template

## Structure

[`/api/controllers`](/api/controllers/README.md) - Директория с пакетами контроллеров

`/api/router` - Пакет с роутером

[`/configs`](/configs/README.md) - Директория для конфигов

`/internal/app` - Пакет для сборки DI контейнера

[`/internal/config`](/internal/config/README.md) - Пакет с глобальным конфигом приложения

[`/internal/entities`](/internal/entities/README.md) - Директория с пакетами сущностей (моделей бд)

[`/internal/models`](/internal/models/README.md) - Директория с пакетами моделей (CreateModel и тд.)

[`/internal/repositories`](/internal/repositories/README.md) - Директория с репозиториями

[`/internal/services`](/internal/services/README.md) - Директория с сервисами

`/server` - Пакет сервера

## Setup

```sh
go run setup.go
```
