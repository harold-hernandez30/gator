# Overview - Gator

Gator is a CLI RSS feed service. It allows users to add feeds the want to be fed.

## Usage

### Example

```bash
go run . feeds
```

### Available Commands

#### register

Adds a user

#### reset

Delete all users

#### users

Get all users

#### agg

Long running tasks that listens for feeds

#### addfeed

Add a feed

##### params

feedName - name of the feed
feedUrl - url of the feed

#### feeds

List a feed

#### follow

Follow a feed

#### unfollow

Unfollow a feed

#### browse

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
