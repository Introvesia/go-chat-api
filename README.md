# go-chat-api

Container management API with Golang

## Prerequisites

- Golang installed

## Preparation for migration tools

Install migration tools:

```bash
brew install golang-migrate
```

Create a new migration:

```bash
migrate create -ext sql -dir migrations -seq <name>
```

## Migrate database

```bash
go run migrate.go
```

## Check the module is uninstalled

```bash
go list -m all | grep <module>
```

## Uninstall the module

```bash
go mod tidy
```

## Login

```bash
curl -X POST http://localhost:8082/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin", "password":"password123"}'
```

## List of channels

```bash
curl -X POST http://localhost:8082/channels \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json"
```

## Hot reload

Install `air`:

```bash
go install github.com/air-verse/air@latest
```

Add this code after edit profile use command `nano ~/.bashrc`:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Run `main.go` with this command:

```bash
air
```