# fiber-app

A minimal REST API built with [Fiber](https://gofiber.io/) (Go web framework), structured
for easy extension.

## Project structure

```
fiber-app/
├── main.go              # entrypoint, middleware setup
├── config/               # environment-based configuration
│   └── config.go
├── models/                # data types
│   └── user.go
├── handlers/               # request handlers (business logic)
│   └── user_handler.go
├── routes/                 # route registration
│   └── routes.go
├── go.mod
└── go.sum
```

## Requirements

- Go 1.22+

## Run locally

```bash
go mod tidy
go run main.go
```

The server starts on `:3000` by default. Set `PORT` to override:

```bash
PORT=8080 go run main.go
```

## Endpoints

| Method | Path                  | Description        |
|--------|------------------------|---------------------|
| GET    | `/health`               | Health check        |
| GET    | `/api/v1/users`         | List all users      |
| GET    | `/api/v1/users/:id`     | Get a user by ID    |
| POST   | `/api/v1/users`         | Create a user       |
| PUT    | `/api/v1/users/:id`     | Update a user       |
| DELETE | `/api/v1/users/:id`     | Delete a user       |

### Example requests

```bash
curl -X POST localhost:3000/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Ada Lovelace","email":"ada@example.com"}'

curl localhost:3000/api/v1/users
```

## Notes

- The `handlers.UserHandler` currently stores users in an in-memory map guarded by a mutex.
  Swap this out for a real database (e.g. via `gorm` or `pgx`) by giving `UserHandler` a
  repository/store dependency instead.
- Middleware included: `recover` (panic recovery), `logger` (request logging), `cors`.

## Build a binary

```bash
go build -o fiber-app .
./fiber-app
```
