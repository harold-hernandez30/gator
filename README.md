# Overview - Gator

## Postgress CLI

### Connect to Postgresql

```bash
sudo -u postgres psql
```

### Connect to database

```bash
\c gator
```

#### Show Table relations

```bash
\dt
```

## Migrate Up/Down

### Migrate up

To run migrate up manually

```bash
goose postgres "postgres://postgres:postgres@localhost:5432/gator" up
```

### Migrate down

```bash
goose postgres "postgres://postgres:postgres@localhost:5432/gator" down
```

## SQLC

### Config file

[SQLC.yaml](./sqlc.yaml)

#### schema

Our migration script

#### queries

Where our SQL queries are stored in this project

##### Run queries

Run from the root of the project

```bash
sqlc generate
```

#### engine

`postgresql` - our database engine

#### gen -> go -> out

The output from our [Run Queries](#run-queries). Declared with internal/database

## Dependencies

### Google

```bash
go get github.com/google/uuid
```

### Driver

```bash
go get github.com/lib/pq
```
