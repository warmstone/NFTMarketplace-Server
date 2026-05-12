# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build & Development Commands

| Command | Purpose |
|---------|---------|
| `make build` | Build binary to `./main` |
| `make run` | Run the server directly |
| `make watch` | Live reload via [air](https://github.com/air-verse/air) (auto-installs if missing) |
| `make test` | Run all tests (`go test ./... -v`) |
| `make itest` | Run integration tests only (`go test ./internal/database -v`) |
| `make docker-run` | Start PostgreSQL container via Docker Compose |
| `make docker-down` | Stop PostgreSQL container |
| `make clean` | Remove built binary |

## Architecture

**Go 1.26** HTTP server for an NFT marketplace backend.

```
cmd/api/main.go          → Entry point. Creates server, handles graceful shutdown (SIGINT/SIGTERM, 5s timeout).
internal/server/          → HTTP layer. Gin router with CORS (allows localhost:5173), registers routes.
internal/database/        → PostgreSQL access via pgx. Singleton `Service` interface with Health() and Close().
```

- **Server** (`internal/server/server.go`): Reads `PORT` from env, wires up the `database.Service`, builds an `*http.Server` with Gin as the handler. Read/write/idle timeouts are 10s/30s/60s.
- **Routes** (`internal/server/routes.go`): `GET /` returns `{"message":"Hello World"}`, `GET /health` proxies to `db.Health()`.
- **Database** (`internal/database/database.go`): Package-level vars read env vars prefixed `NFTMAREKTPLACE_DB_*` (DATABASE, USERNAME, PASSWORD, PORT, HOST, SCHEMA). `New()` is a singleton — reuses the connection on subsequent calls. Connects via `pgx` driver with SSL disabled. `Health()` pings with a 1s timeout and returns connection pool stats.
- **Graceful shutdown** (`cmd/api/main.go`): Catches SIGINT/SIGTERM, calls `server.Shutdown()` with a 5s deadline, then signals `done`.

## Environment Variables

Loaded via `godotenv/autoload` (`.env` file). Key variables:

- `PORT` — server listen port
- `NFTMAREKTPLACE_DB_DATABASE`, `NFTMAREKTPLACE_DB_USERNAME`, `NFTMAREKTPLACE_DB_PASSWORD`
- `NFTMAREKTPLACE_DB_HOST`, `NFTMAREKTPLACE_DB_PORT`, `NFTMAREKTPLACE_DB_SCHEMA`

## Testing

- Unit tests use standard `testing` + `httptest` for handlers.
- **Database integration tests** use [testcontainers-go](https://github.com/testcontainers/testcontainers-go) to spin up a real PostgreSQL container. `TestMain` in `database_test.go` starts the container before tests run and tears it down afterward. The test overrides the package-level connection vars to point at the container.

# CLAUDE.md

Behavioral guidelines to reduce common LLM coding mistakes. Merge with project-specific instructions as needed.

**Tradeoff:** These guidelines bias toward caution over speed. For trivial tasks, use judgment.

## 1. Think Before Coding

**Don't assume. Don't hide confusion. Surface tradeoffs.**

Before implementing:
- State your assumptions explicitly. If uncertain, ask.
- If multiple interpretations exist, present them - don't pick silently.
- If a simpler approach exists, say so. Push back when warranted.
- If something is unclear, stop. Name what's confusing. Ask.

## 2. Simplicity First

**Minimum code that solves the problem. Nothing speculative.**

- No features beyond what was asked.
- No abstractions for single-use code.
- No "flexibility" or "configurability" that wasn't requested.
- No error handling for impossible scenarios.
- If you write 200 lines and it could be 50, rewrite it.

Ask yourself: "Would a senior engineer say this is overcomplicated?" If yes, simplify.

## 3. Surgical Changes

**Touch only what you must. Clean up only your own mess.**

When editing existing code:
- Don't "improve" adjacent code, comments, or formatting.
- Don't refactor things that aren't broken.
- Match existing style, even if you'd do it differently.
- If you notice unrelated dead code, mention it - don't delete it.

When your changes create orphans:
- Remove imports/variables/functions that YOUR changes made unused.
- Don't remove pre-existing dead code unless asked.

The test: Every changed line should trace directly to the user's request.

## 4. Goal-Driven Execution

**Define success criteria. Loop until verified.**

Transform tasks into verifiable goals:
- "Add validation" → "Write tests for invalid inputs, then make them pass"
- "Fix the bug" → "Write a test that reproduces it, then make it pass"
- "Refactor X" → "Ensure tests pass before and after"

For multi-step tasks, state a brief plan:
```
1. [Step] → verify: [check]
2. [Step] → verify: [check]
3. [Step] → verify: [check]
```

Strong success criteria let you loop independently. Weak criteria ("make it work") require constant clarification.

---

**These guidelines are working if:** fewer unnecessary changes in diffs, fewer rewrites due to overcomplication, and clarifying questions come before implementation rather than after mistakes.
