run:
	go run cmd/main.go
swag:
	swag init --parseInternal --parseGoList -g cmd/main.go 
.PHONY: sqlup sqldown run