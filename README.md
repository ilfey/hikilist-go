# Golang API template

## Setup

```sh
go run setup.go
```

## Migrations

### Make migrations

```sh
atlas migrate diff --env gorm
```

### Migrate

```sh
atlas schema apply --env gorm --url "<dsn>"
```

## Entities

### Schema

```mermaid
erDiagram
    animes {
        uint id pk

        string title
        string description
        string poster
        uint episodes
        uint episodes_released
        uint mal_id
        uint shiki_id

        datetime created_at
        datetime updated_at
        datetime deleted_at
    }

    animes_related {
        uint anime_id pk,fk
        uint related_id pk,fk
    }

    animes }o--o{ animes_related : ""

    users {
        uint id pk

        string username
        string password

        datetime created_at
        datetime updated_at
        datetime deleted_at
    }
```
