# Resource Management
[![CircleCI](https://circleci.com/gh/thinhlvv/resource-management/tree/master.svg?style=svg&circle-token=53c4910bab358e892b7fc5856bbf1b2d6837ea18)](https://circleci.com/gh/thinhlvv/resource-management/tree/master)

API server to manage Resources .

## Commands

```bash
# Make copy of the environment for database in development.
$ cp .env.sample .env

# Make copy of the environment for database in staging.
$ cp .env.sample .env.staging

# Install all dependencies.
$ make install

# Start docker services.
$ docker-compose up -d

# Run migration files (locally), default is set to development.
$ make migrate

# Rollback migration version
$ make rollback

# Run migration files (on staging).
$ make migrate ENV=staging

# Run test.
$ make test

# Start the development server.
$ make start

# Cleanup local database.
$ make clean

# Review code.
$ make review
```

## Installation

Go version : 1.12.6

1. Run database for application and testing:
```bash
docker-compose up -d
```

Make sure you got two docker database:
```bash
docker ps 
```

2. Install dependencies:
```bash
make install
```

3. Migrate database:
```bash
make migrate
```

Troubleshoot:

`goose: command not found` - make sure you export $PATH=$GOPATH/bin

4. Start server
```bash
make start
```
