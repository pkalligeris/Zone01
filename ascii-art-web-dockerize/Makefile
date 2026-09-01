.PHONY: all build test clean run run-examples fmt

BINARY_NAME := ascii-art
SOURCE_DIR := ./cmd/ascii-art

all: build

build:
	go build -o $(BINARY_NAME) $(SOURCE_DIR)

test:
	go test ./...

clean:
	rm -f $(BINARY_NAME)

# Run with default argument "Hello". Override with: make run ARGS="something else"
ARGS ?= "Hello"
run:
	go run $(SOURCE_DIR) $(ARGS)

run-examples:
	go run $(SOURCE_DIR) "Hello"
	go run $(SOURCE_DIR) "wfjeh"
	go run $(SOURCE_DIR) "Hello\nWorld"
	go run $(SOURCE_DIR) "12345"
	@echo "--- Standard ---"
	@go run $(SOURCE_DIR) "Hello"
	@echo "\n--- Banner: Shadow ---"
	@go run $(SOURCE_DIR) "Hello" shadow
	@echo "\n--- Banner: Thinkertoy ---"
	@go run $(SOURCE_DIR) "Hello" thinkertoy
	@echo "\n--- Color: Red ---"
	@go run $(SOURCE_DIR) --color=red "Hello"
	@echo "\n--- Color: Green Substring ---"
	@go run $(SOURCE_DIR) --color=green "ll" "Hello"
	@echo "\n--- Align: Right ---"
	@go run $(SOURCE_DIR) --align=right "Hello"
	@echo "\n--- Align: Center ---"
	@go run $(SOURCE_DIR) --align=center "Hello"
	@echo "\n--- Align: Justify ---"
	@go run $(SOURCE_DIR) --align=justify "Hello World"
	@echo "\n--- Error: No Arguments ---"
	-@go run $(SOURCE_DIR)
	@echo "\n--- Error: Missing Color Flag '=' ---"
	-@go run $(SOURCE_DIR) --color red "Hello"
	@echo "\n--- Error: Missing Output Flag '=' ---"
	-@go run $(SOURCE_DIR) --output result.txt "Hello"
	@echo "\n--- Error: Missing Align Flag '=' ---"
	-@go run $(SOURCE_DIR) --align right "Hello"

fmt:
	go fmt ./...