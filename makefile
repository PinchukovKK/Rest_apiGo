DB_DSN := "postgres://postgres:123GO123d@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down 

migrate-force:
	$(MIGRATE) force 20250212105252
	
run:
	go run cmd/app/main.go