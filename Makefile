sqlc:
	- sqlc generate -f ./config/sqlc.yaml
run:
	- go run cmd/main.go

migrate-create:
	- migrate create -ext sql -dir internal/constant/query/schemas $(args)