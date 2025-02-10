# go-container-api

Container management API with Golang

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