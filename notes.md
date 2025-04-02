# Notes

- Ensure to add Go bin folder to PATH ($GOPATH/bin) for global installs and commands:

```powershell 
$env:Path += ";$(go env GOPATH)\bin"
```

To make this change persist across sessions:

```powershell
[System.Environment]::SetEnvironmentVariable("Path", $env:Path + ";$(go env GOPATH)\bin", [System.EnvironmentVariableTarget]::User)
```

This permanently adds $(go env GOPATH)\bin to the user's PATH.

- The DEFAULT mux from the net/http library i.e http.NewServeMux() does not support dynamic URLs and wildcards hence the need for some external mux/router library.

- Gorilla mux library is depreciated and go-chi is used for this project as a modern choice.

- Air Verse is installed and used for hot reloads - github.com/air-verse/air@latest

- Edit .air.toml file to suit project (directory structure)

- Air Verse is initialized using the command - `air init` and boots project using - `air`

- github.com/lib/pq is the driver for postgres and interfaces with the "database/sql"

- "database/sql" is an abstraction 

- docker was used for installing and running postgres instances

- There are 2 migrations libraries preferred for Go backends - Goose and Golang-migrate

- For Golang-migrate:

- golang-migrate is installed using: go get -u github.com/golang-migrate/migrate/v4

- Create migration command: migrate create -seq -ext sql -dir ./internal/db/migrations create_users

- Run migration command: migrate -database "DB_URL" ./internal/db/migrations migrations up

- For Goose (which is used for this project)

- Create migration command: goose -dir internal/db/migrations create create_users_table sql

- Run migration command: goose -dir internal/db/migrations mysql "root:@tcp(localhost:3306)/go_social" up

- To undo the last migration: goose -dir internal/db/migrations mysql "root:@tcp(localhost:3306)/go_social" down

- To know migration status: goose -dir internal/db/migrations mysql "root:@tcp(localhost:3306)/go_social" status

- Makefile is added to store easily run commands

- Install make on Windows using chocolatey - choco install make

- Struct validation was done using "github.com/go-playground/validator/v10" library

- install using go get github.com/go-playground/validator/v10