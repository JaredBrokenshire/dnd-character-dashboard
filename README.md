# D&D Character Dashboard

A full-stack web application for managing Dungeons & Dragons character sheet information. Users can register an account, view their characters (with race, class, subclass, campaign, and portrait), and browse campaigns.

- **Backend**: Go 1.26 · Echo v4 · GORM v1 · MySQL 8.0
- **Frontend**: Vue 2 · Vuex · Vue Router · Tailwind CSS · Axios
- **Infrastructure**: Docker Compose (API + database containers, live-reload via CompileDaemon)
- **CI**: GitHub Actions (build + test on every push to `main` and `development`)

---

## Quick Start

### 1. Configure environment

```bash
cp .env.example .env
```

Edit `.env` — at minimum change `JWT_SECRET`, `ADMIN_EMAIL`, and `ADMIN_PASSWORD` before running.

For the frontend, create `frontend/.env` from the example:

```bash
cp frontend/.env.example frontend/.env
```

Edit `VUE_APP_API_BASE_URL` to point at your API (default: `http://api.localdnd.com:7788`).

### 2. Start the stack

```bash
docker compose up --build
```

| Service | Default URL |
|---------|-------------|
| API | `http://localhost:7788` |
| MySQL | `localhost:56741` |

The API container mounts the `backend/` source directory as a volume and uses **CompileDaemon** for live reload — source changes are picked up without rebuilding the image.

### 3. Start the frontend dev server

```bash
cd frontend
npm install
npm run serve
```

---

## Project Structure

```
dnd-character-dashboard/
├── .env.example                        # Root environment variable template
├── docker-compose.yml                  # Full stack: api + database
├── docker-compose-test.yml             # Test database container only
│
├── backend/
│   ├── cmd/api/api.go                  # Entry point — loads env, wires server, starts Echo
│   ├── api/
│   │   ├── server.go                   # Server struct (Echo + DB + Repos + Services)
│   │   ├── handlers/                   # HTTP handlers
│   │   ├── routes/                     # Route registration
│   │   ├── requests/                   # Request DTOs
│   │   └── responses/                  # Response DTOs
│   ├── config/
│   │   └── config.go                   # SysConfig singleton — env-backed, defaults in code
│   ├── db/
│   │   ├── db.go                       # MySQL connect, run migrations, seed, Paginate() helper
│   │   ├── models/                     # GORM models
│   │   ├── migrations/                 # Versioned migration files + registration list
│   │   ├── repositories/               # Data-access layer
│   │   └── seeders/                    # Admin user + reference data seeders
│   ├── services/                       # Business logic
│   ├── pkg/
│   │   ├── closer/                     # Helper method to close the environment down
│   │   ├── middleware/                 # JWT auth middleware, logging middleware
│   │   ├── logging/                    # zerolog server logger
│   │   ├── validation/                 # Custom Echo validator (go-playground/validator)
│   │   └── go-migrations/              # Vendored go-migrations fork
│   ├── tests/                          # Table-driven integration tests
│   └── docker/Dockerfile               # Alpine Go image with CompileDaemon + swag
│
└── frontend/
    ├── .env.example                    # Frontend environment variable template
    ├── package.json                    # npm scripts: serve, build, lint
    ├── tailwind.config.js
    └── src/
        ├── router/                     # Vue Router with auth guards (requiresAuth, requiresAdmin, guestOnly)
        ├── views/                      # Page Views
        ├── store/                      # Vuex modules (auth, etc.)
        ├── services/                   # Axios API client + per-resource service modules
        └── components/                 # Shared UI components and layouts
```

---

## API Endpoints

All routes are registered under the root path (`/`). Every protected route requires a `Authorization: Bearer <token>` header.

**Auth** — rate-limited

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| POST | `/auth/register` | — | Create a new user account |
| POST | `/auth/login` | — | Login, returns JWT access + refresh token |
| POST | `/auth/refresh` | — | Exchange a refresh token for new tokens |
| GET | `/auth/me` | ✓ | Get the authenticated user's profile |
| POST | `/auth/logout` | ✓ | Logout and invalidate the session |

**Class**

| Method | Path             | Auth | Description                   |
|--------|------------------|------|-------------------------------|
| GET | `/classes`    | ✓ | Get paginated list of classes |

**Race**

| Method | Path         | Auth | Description                 |
|--------|--------------|------|-----------------------------|
| GET | `/races`     | ✓ | Get paginated list of races |
| GET | `/races/:id` | ✓ | Get race by id              |

**File**

| Method | Path         | Auth | Description                                       |
|--------|--------------|------|---------------------------------------------------|
| GET | `/:filepath` | ✓ | Get file from local storage at location: filepath |

---

## Environment Variables

Copy `.env.example` to `.env` and set the values for your environment.

| Variable | Default / Example | Description |
|----------|-------------------|-------------|
| `ENVIRONMENT` | `development` | Runtime mode (`production` or `development`) |
| `HOST` | `localhost` | Application hostname |
| `PORT` | `7788` | Port the Go server listens on inside the container |
| `DB_DRIVER` | `mysql` | Database driver |
| `DB_HOST` | `dnd_mysql` | Database hostname (matches Docker Compose service name) |
| `DB_PORT` | `3306` | Database port |
| `DB_NAME` | `dnd` | Database name |
| `DB_USER` | `local_user` | Database username |
| `DB_PASSWORD` | `password` | Database password |
| `EXPOSE_PORT` | `7788` | Host → API container port mapping |
| `EXPOSE_DB_PORT` | `56741` | Host → MySQL container port mapping |
| `JWT_SECRET` | *(must be changed)* | Secret used to sign JWT tokens |
| `ADMIN_EMAIL` | `admin@example.com` | Seeded admin account email |
| `ADMIN_PASSWORD` | `Admin1234!` | Seeded admin account password |
| `CORS_ORIGINS` | `http://localdnd.com:8080` | Comma-separated list of allowed CORS origins |

Additional settings can be tuned via environment variables and are read by `backend/config/config.go` (all have safe defaults):

| Variable | Default | Description |
|----------|---------|-------------|
| `LOG_MIDDLEWARE` | `true` | Enable request logging |
| `HASH_COST` | `10` | bcrypt hash cost |
| `LOGIN_RATE_LIMIT` | `10` | Max login requests per second |
| `BRUTE_FORCE_LIMIT` | `5` | Max failed login attempts before lockout |
| `BRUTE_FORCE_DURATION` | `15m` | Lockout window duration |

---

## Running Backend Tests

Tests use a real MySQL instance (same as the Docker Compose `dnd_mysql` service, or the test DB container).

Start the test database:

```bash
docker compose -f docker-compose-test.yml up -d
```

Then run the test suite:

```bash
cd backend
go test ./tests/... -v -timeout 120s
```

### Test Architecture

Tests are table-driven, using a `[]helpers.TestCase` slice per endpoint. Each `TestCase` supports:

| Field | Purpose |
|-------|---------|
| `Name` | Sub-test label |
| `Method` / `URL` / `QueryParams` | Request to fire |
| `Body` | JSON request body (marshalled automatically) |
| `AccessToken` | Sent as `Authorization: Bearer <token>` |
| `ExpectedStatus` | Required HTTP status code |
| `ExpectedBody` | Substrings that must appear in the response |
| `AbsentBody` | Substrings that must NOT appear |
| `DatabaseChecks` | Post-request assertions against the DB |
| `ExpectedCallback` | Full custom assertion on the raw response bytes |
| `Setup` | Fixture setup run before the request |
| `Teardown` | Fixture cleanup run after assertions |

---

## Frontend Setup

```bash
cd frontend
npm install
npm run serve    # Dev server with hot-reload
npm run build    # Production build → frontend/dist/
npm run lint     # ESLint with auto-fix
```

The frontend reads `VUE_APP_API_BASE_URL` from `frontend/.env` to know where to send API requests.

---

## Adding Migrations

1. Create a new file in `backend/db/migrations/list/`.
2. Register the new struct in `backend/db/migrations/process/`.

Migrations run automatically on server startup.

---

## Initial Admin Account

On first startup the seeder creates an admin user from:

```
ADMIN_EMAIL=admin@example.com
ADMIN_PASSWORD=Admin1234!
```

Change these in `.env` before deploying to a shared or production environment.
