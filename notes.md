# Notes

- The default mux from the net/http library i.e http.NewServeMux() does not support dynamic URLs and wildcards hence the need for some external mux/router library.

- Gorilla mux library is depreciated and go-chi is used for this project as a modern choice.

- Air Verse is installed and used for hot reloads - github.com/air-verse/air@latest

- Edit .air.toml file to suit project (directory structure)

- Air Verse is initialized using the command - `air init` and boots project using - `air`

- github.com/lib/pq is the driver for postgres and interfaces with the "database/sql"

- "database/sql" 