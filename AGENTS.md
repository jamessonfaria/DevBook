# Repository Guidelines

## Project Structure & Module Organization
This repository is a small Go API. The entrypoint is `main.go`, which loads environment configuration and starts the HTTP server. Application code lives under `src/`: `controllers/` handles HTTP requests, `router/` and `router/rotas/` define routes, `repositorios/` contains database access, `modelos/` defines domain structs, `banco/` manages DB connections, `config/` loads `.env`, and `respostas/` centralizes API responses. SQL assets live in `sql/sql.sql`.

## Build, Test, and Development Commands
Use the standard Go toolchain from the repository root:

```bash
go run .
go build .
go test ./...
gofmt -w main.go src
```

`go run .` starts the API using values from `.env`. `go build .` compiles the service and catches type errors. `go test ./...` runs all package tests; there are no `_test.go` files yet, so add tests as new behavior is implemented. `gofmt -w` should be run before commits.

## Coding Style & Naming Conventions
Follow idiomatic Go formatting with tabs via `gofmt`; do not align fields manually. Keep package names lowercase and short (`config`, `router`, `repositorios`). Exported identifiers use `PascalCase`; unexported helpers use `camelCase`. Match the current Portuguese naming in handlers and repository methods, for example `CriarUsuario` and `NovoRepositorioUsuarios`.

## Testing Guidelines
Place tests beside the code they cover using `*_test.go` filenames, such as `src/controllers/usuarios_test.go`. Prefer table-driven tests for handlers, validation, and repository logic. For database-related code, isolate setup data clearly and avoid coupling tests to a developer’s local schema beyond the expected `.env` configuration.

## Commit & Pull Request Guidelines
Recent commits use short, imperative summaries in Portuguese, for example `router, rotas e controllers usuarios` and `implementado respostas`. Keep commits focused on one change. Pull requests should include a concise description, note any API or database changes, list test commands run, and include sample request/response payloads when endpoint behavior changes.

## Security & Configuration Tips
Copy `.env.example` to `.env` and set `DB_USUARIO`, `DB_SENHA`, `DB_NOME`, and `API_PORT` locally. Do not commit real credentials. Review SQL changes in `sql/` with code changes so schema and repository behavior stay aligned.
