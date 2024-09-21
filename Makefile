APP_NAME := Backend
SRC := cmd/main.go

# Build the application
build:
	go build -o $(APP_NAME) $(SRC)

# Run the application
run: build
	./$(APP_NAME)