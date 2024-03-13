sqlup:
	migrate -path internal/migrate -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"  up
sqldown:
	migrate -path internal/migrate -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"  down
.PHONY: sqlup sqldown