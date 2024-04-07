# GoBackendApp


### Setup

Setup Go Modules:
```bash
go mod init github.com/ShubhamTatvamasi/GoBackendApp
```

Setup GORM:
```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

### Postgres Docker

Setup Postgres Database:
```bash
docker run -d -p 5432:5432 \
  --name postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=simple_bank \
  postgres:16-alpine
```

Connect to Postgres Database:
```bash
PGPASSWORD=postgres psql -h 0.0.0.0 -U postgres -d simple_bank
```

List active sessions:
```sql
SELECT * FROM pg_stat_activity;
```

Remove Postgres Database:
```bash
docker rm -f postgres
```
