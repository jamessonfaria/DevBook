# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

DevBook API is a small Go REST API backed by MySQL. Identifiers, comments, and commit messages are in Portuguese — match that convention when adding code.

## Commands

```bash
go run .              # start the API using values from .env
go build .             # compile and catch type errors
go test ./...          # run all tests (no _test.go files exist yet)
gofmt -w main.go src   # format before committing
```

Database schema is defined in `sql/sql.sql`; apply it to a local MySQL instance before running the API. Copy `.env.example` to `.env` and set `DB_USUARIO`, `DB_SENHA`, `DB_NOME`, and `API_PORT`. Never commit real credentials.

## Architecture

Request flow follows a straightforward layered pattern, wired together per-request (no shared connection pool or DI container):

1. **`main.go`** calls `config.Carregar()` then `router.Gerar()` and starts `http.ListenAndServe`.
2. **`src/config`** loads `.env` via `godotenv` and builds the MySQL DSN (`config.StringConexaoBanco`) and port (`config.Porta`) as package-level vars.
3. **`src/router` / `src/router/rotas`** builds a `gorilla/mux` router. Routes are declared as `Rota` structs (URI, HTTP method, handler func, `RequerAutenticacao` flag) in per-resource files like `rotas/usuarios.go`, aggregated in `rotas.go`'s `Configurar`. Note: `RequerAutenticacao` is currently declared but not enforced anywhere — there is no auth middleware yet.
4. **`src/controllers`** holds HTTP handlers (e.g. `usuarios.go`). Each handler opens its own DB connection via `banco.Conectar()`, `defer db.Close()`s it, instantiates a repository, and writes the result with `src/respostas`. There is no shared request context or middleware chain.
5. **`src/repositorios`** contains raw-SQL data access structs (e.g. `usuarios` in `repositorios/usuarios.go`) built with `NovoRepositorio<Entity>(db)`, using prepared statements via `database/sql`.
6. **`src/modelos`** defines domain structs with a `Preparar(etapa string) error` method that validates then formats/trims fields. `etapa` distinguishes flows (e.g. `"cadastro"` vs `"edicao"`) where validation rules differ (e.g. password required only at signup).
7. **`src/banco`** wraps `sql.Open`/`Ping` for MySQL (`github.com/go-sql-driver/mysql`) connection creation only — no pooling/reuse across requests.
8. **`src/respostas`** centralizes JSON responses: `respostas.JSON(w, status, data)` and `respostas.Error(w, status, err)`.

When adding a new resource, mirror this exact chain: model in `modelos/`, repository in `repositorios/`, handler in `controllers/`, route entry in `router/rotas/<resource>.go` registered inside `rotas.Configurar`.

## Coding Conventions

- Format with `gofmt` (tabs); don't hand-align fields.
- Package names: lowercase, short (`config`, `router`, `repositorios`).
- Exported identifiers: `PascalCase`; unexported: `camelCase`.
- Match existing Portuguese naming for handlers/repo methods (e.g. `CriarUsuario`, `NovoRepositorioUsuarios`).

## Testing Guidelines

Place tests beside the code they cover using `*_test.go` filenames (e.g. `src/controllers/usuarios_test.go`). Prefer table-driven tests for handlers, validation, and repository logic. For DB-related tests, isolate setup data clearly and avoid coupling to a developer's local schema beyond the `.env` configuration.

## Commit Guidelines

Commits use short, imperative summaries in Portuguese (e.g. `implementado respostas`, `validacoes`). Keep each commit focused on one change.
