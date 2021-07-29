GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=server_scratch_bin

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

migrate/up:
	migrate -source file://db/migrations -database 'mysql://root@tcp(127.0.0.1:3306)/server_scratch_dev' up

migrate/down:
	migrate -source file://db/migrations -database 'mysql://root@tcp(127.0.0.1:3306)/server_scratch_dev' down

migrate/reset: migrate/down migrate/up