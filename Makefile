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

# Create a new migration file
.PHONY: migrate-create
migrate-create:
	@$(GOOSE_BIN) -dir $(MIGRATIONS_PATH) create $(name) sql
	# Run using - make migrate-create name=my_migration

# Run migrations (Up)
.PHONY: migrate-up
migrate-up:
	@echo "Running Goose migrations..."
	@$(GOOSE_BIN) -dir $(MIGRATIONS_PATH) $(DBDRIVER) "$(DBURL)" up
	# Run using - make migrate-up

# Run migrations (Down)
.PHONY: migrate-down
migrate-down:
	@echo "Rolling back migrations..."
	@$(GOOSE_BIN) -dir $(MIGRATIONS_PATH) $(DBDRIVER) "$(DBURL)" down
	# Run using - make migrate-down

# Show migration status
.PHONY: migrate-status
migrate-status:
	@echo "Fetching migration status..."
	@$(GOOSE_BIN) -dir $(MIGRATIONS_PATH) $(DBDRIVER) "$(DBURL)" status
	# Run using - make migrate-status
