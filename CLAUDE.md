# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common Development Commands

| Task | Command |
|------|---------|
| **Build the binary** | `go build ./cmd/server` |
| **Run the server (development)** | `go run ./cmd/server` |
| **Run all unit/integration tests** | `go test ./...` |
| **Run a single test** | `go test ./... -run ^TestName$` (replace `TestName` with the exact test function name) |
| **Run linter** | `golangci-lint run` (assumes the project has a `.golangci.yml` configuration) |
| **Generate Ent schema & code** | `go generate ./internal/data/ent` |
| **Apply database migrations** | `go run ./cmd/server migrate` (the `migrate` sub‑command is implemented in `cmd/server/main.go`) |
| **Load environment variables** | `source config/config.env` (contains the Firebase service key path and other env vars) |
| **Run Docker compose (dev)** | `docker-compose -f docker-compose.dev.yaml up` |
| **Run Docker compose (prod)** | `docker-compose -f docker-compose.prod.yaml up` |

*Tip:* Most commands can be run from the repository root. Use `make` targets if a `Makefile` is present (e.g., `make build`, `make test`).

---

## High‑Level Architecture Overview

```
cmd/
 └─ server/
      main.go            ← entry point, sets up the HTTP server and wires dependencies
internal/
 ├─ app/                 ← top‑level application bootstrap (creates the Gin engine)
 ├─ cache/               ← in‑memory or Redis cache abstractions
 ├─ controller/          ← HTTP handlers (one file per resource, e.g., auth, user, notification)
 ├─ cron/                ← scheduled background jobs (e.g., cleanup, batch notifications)
 ├─ data/
 │   ├─ datasource/      ← low‑level DB connection handling
 │   └─ ent/             ← Ent ORM generated code (entities, queries, migrations)
 ├─ firebase/            ← Firebase client wrapper (FCM, Auth, etc.)
 ├─ handler/             ← reusable request/response helpers (JSON binding, error handling)
 ├─ helper/              ← utility functions (e.g., token generation, hashing)
 ├─ job/                 ← long‑running jobs executed by the cron system
 ├─ mapper/              ← conversion between DB models and API DTOs
 ├─ middleware/          ← Gin middleware (auth, logging, CORS, permission checks)
 ├─ permission/          ← RBAC definitions and permission checks
 ├─ route/               ← API route registration (grouped by feature)
 ├─ server/              ← server configuration (port, TLS, graceful shutdown)
 └─ service/             ← business‑logic services used by controllers

docs/
 ├─ swagger.yaml         ← OpenAPI spec for the REST API
 └─ docs.go              ← HTTP handler that serves the Swagger UI
pkg/
 └─ (shared libraries, if any)
vendor/
 └─ third‑party dependencies (managed by Go modules)
```

### Core Flow

1. **Startup** (`cmd/server/main.go`)  
   - Loads environment variables from `config/config.env`.  
   - Initializes the Ent client (`internal/data/ent`) and Firebase client (`internal/firebase`).  
   - Creates the Gin engine via `internal/app/app.go`.  
   - Registers middleware (`internal/middleware`) and routes (`internal/route`).  
   - Starts the HTTP server.

2. **Request Handling**  
   - Incoming HTTP requests are routed to a controller in `internal/controller`.  
   - Controllers invoke corresponding **services** in `internal/service` to perform business logic.  
   - Services interact with the **data layer** (`internal/data/ent` for relational data, `internal/firebase` for push notifications).  
   - Results are mapped to API response structs by `internal/mapper` and sent back via the `handler` helpers.

3. **Background Jobs** (`internal/cron`)  
   - A cron scheduler is started during boot (`internal/cron/cron.go`).  
   - Jobs are defined in `internal/job` and may use the same services and data layer as HTTP handlers.

4. **Authentication & Authorization**  
   - JWT validation is performed by middleware in `internal/middleware/auth.go`.  
   - Permission checks are centralized in `internal/permission` and referenced by route definitions.

5. **Firebase Integration**  
   - The Firebase client (`internal/firebase`) wraps the Firebase Admin SDK for tasks such as sending FCM messages and managing user tokens.

### Important Files to Inspect First

| Purpose | Path |
|---------|------|
| Server entry point | `cmd/server/main.go` |
| Application bootstrap (Gin engine) | `internal/app/app.go` |
| Route registration | `internal/route/router.go` (or similar) |
| Auth middleware | `internal/middleware/auth.go` |
| Ent schema (generated) | `internal/data/ent/schema/*.go` |
| Firebase service wrapper | `internal/firebase/*.go` |
| Example controller | `internal/controller/user.controller.go` |
| Swagger spec | `docs/swagger.yaml` |

---

## Cursor / Copilot Rules (if present)

- No specific `.cursor` or `.cursorrules` directories were found.  
- No `.github/copilot-instructions.md` file was detected.

---

## README Highlights

- The repository is a **simple Go + Firebase** sample project.  
- An architectural diagram is stored at `asset/fcm_design.png`.  
- The `Makefile` (if present) provides shortcuts for building, testing, and running the server.

---

### Usage Tips for Future Claude Instances

- When adding new features, follow the existing **controller → service → data** pattern.  
- Register new routes in the appropriate group within `internal/route`.  
- Extend the Ent schema and run `go generate ./internal/data/ent` to regenerate ORM code.  
- Update `docs/swagger.yaml` and run `make swagger` (if a make target exists) to keep API docs in sync.

---

*End of CLAUDE.md*