run:
	go run ./cmd

# Build the application
build:
	go build -o bin/ecommerce ./cmd

# Run in development mode with live reload (if you have air installed)
dev:
	air -c .air.toml

# Clean build artifacts
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod tidy
	go mod download

