sqlc:
	sqlc generate
run:
	go run main.go

migratedb:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/bd?sslmode=disable" -verbose up

migratedowndb:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/bd?sslmode=disable" -verbose down

.PHONY: sqlc run migratedb migratedowndb