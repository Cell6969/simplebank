# Simple Bank Application with Golang

## create migration file
```shell
migrate create -ext sql -dir db/migration -seq init_schema
```

## run migration file
```shell
migrate -path db/migration -database "postgresql://root:root@localhost:5432/gosimplebank?sslmode=disable" -verbose up
```

## rollback migration file
```shell
migrate -path db/migration -database "postgresql://root:root@localhost:5432/gosimplebank?sslmode=disable" -verbose down
```

## run cover test
```shell
go test -v -cover ./...
```