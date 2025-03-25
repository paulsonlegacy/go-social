# Load environment variables (Windows PowerShell method)
ENV_FILE := ./internal/config/.env
include $(ENV_FILE)
export $(shell powershell -Command "Get-Content $(ENV_FILE) | ForEach-Object { ($_ -match '^(.*?)=(.*)$') -and \"`$env:$($matches[1])='$($matches[2])'\" }")

# Variables
MIGRATIONS_PATH := "./internal/db/migrations"
GOOSE_BIN := $(shell where goose) 
# Why Use GOOSE_BIN Instead of Just goose?
# Works even if goose is installed in a non-standard location
# If multiple versions of goose exist, this avoids conflicts by using the one found in your system’s PATH.

# Run application
# Run using - make air
.PHONY: air
air:
	@air
	# Run using - make air

# Create a new migration file
# Run using - make migrate-create name=my_migration
.PHONY: migrate-create
migrate-create:
	@echo "Creating migration..."
	@$(GOOSE_BIN) -dir $(MIGRATIONS_PATH) create $(name) sql
	
# Run migrations (Up)
# Run using - make migrate-up
.PHONY: migrate-up
migrate-up:
	@echo "Running Goose migrations..."
	@$(GOOSE_BIN) -dir $(MIGRATIONS_PATH) $(DBDRIVER) "$(DBURL)" up

# Run migrations (Down)
# Run using - make migrate-down
.PHONY: migrate-down
migrate-down:
	@echo "Rolling back migrations..."
	@$(GOOSE_BIN) -dir $(MIGRATIONS_PATH) $(DBDRIVER) "$(DBURL)" down

# Show migration status
# Run using - make migrate-status
.PHONY: migrate-status
migrate-status:
	@echo "Fetching migration status..."
	@$(GOOSE_BIN) -dir $(MIGRATIONS_PATH) $(DBDRIVER) "$(DBURL)" status
