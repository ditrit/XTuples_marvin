# xtuples

```bash
# download the dependencies and run the server
go mod tidy
go mod download
go run cmd/api/main.go

# run with a custom config path using flags
go run cmd/api/main.go --conf "path/to/config"
# or run with a custom config path using env vars (useful when running in docker or tests with `go test ./...`)
MY_PROJECT_CONF_PATH="path/to/config" go run cmd/api/main.go
```

### Init project structure

```bash
├── cmd                 # holds the entrypoint for the backend server
├── internal            
│   └── api            
│       ├── modules     # holds the generated modules
│       └── server      # holds all routes and the function which sets up the server
│       └── response    # holds a predefined response function and response messages
├── pkg
│   ├── app             # holds functions which setup the environment for the application
│   ├── convert         # holds util functions for converting types ( used in the settings package )
│   ├── settings        # holds util functions for loading .env vars and storing them in the settings.Environment variable
│   │   ├── cli         # package for dealing with cli flags / arguments
│   │   ├── database    # holds functions for connecting to the db
│   └── validate        # holds functions which check if passed down structs pass their validation tags
└── client              # holds the generated fetch functions for python and typescript
```

#### Update on change

```bash
# https://github.com/cespare/reflex
reflex -r '\.go' -s -- sh -c "go run cmd/api/main.go"
```

#### Renaming generated sql files from `.sql.gen.txt` to `.sql`

```bash
# run from the root of the project
for i in ./db/sql/*.sql.gen.txt;  do mv "$i" "${i/.sql.gen.txt}.sql"; done
```