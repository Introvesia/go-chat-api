# go-chat-api

Chat API server with Golang for Introvesia Workspace

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

1. Uncomment this code at `main.go`
    ```go
    libs.AutoMigrate()
    ```
2. Run the server

## Check the module is uninstalled

```bash
go list -m all | grep <module>
```

## Uninstall the module

```bash
go mod tidy
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