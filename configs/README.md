# Configs

## Priority

- `production.env`
- `development.env`

## Example

```.env
AUTH_SECRET=secret
AUTH_ISSUER=hikilist
AUTH_ACCESS_LIFE_TIME=24
AUTH_REFRESH_LIFE_TIME=168

DB_USER=golang-template-service
DB_PASSWORD=golang-template-service
DB_DBNAME=golang-template-service
DB_HOST=127.0.0.1
DB_PORT=5432

SERVER_READ_TIMEOUT=10000
SERVER_WRITE_TIMEOUT=10000
SERVER_IDLE_TIMEOUT=30000
SERVER_READ_HEADER_TIMEOUT=2000
```
