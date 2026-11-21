# D&D Character Dashboard

A full-stack web application for managing Dungeons & Dragons character sheet information.

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

For the frontend, create `frontend/.env` from the example:

```bash
cp frontend/.env.example frontend/.env
```

Edit `VUE_APP_API_BASE_URL` to point at your API (default: `http://api.localdnd.com:7788`).

### 2. Start the stack

```bash
docker compose up --build
```

| Service  | Default URL              |
|----------|--------------------------|
| API      | `http://localhost:7788`  |
| MySQL    | `http://localhost:56741` |

The API container mounts the `backend/` source directory as a volume and uses **CompileDaemon** for live reload — source
changes are picked up without rebuilding the image.

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
├── .dockerignore
├── .env.example                        # Root environment variable template
├── .gitignore
├── CHANGELOG.md
├── docker-compose.yml                  # Full stack: api + database
├── docker-compose-test.yml             # Test database container only
├── README.md
│
├── backend/
│   ├── api/
│   │   ├── server.go                   # Server struct (Echo + DB + Repos + Services)
│   │   ├── handlers/                   # API Endpoint handlers
│   │   ├── requests/                   # Request structs
│   │   ├── repsonses/                  # Response structs
│   │   └── routes/                     # Route registration
│   ├── cmd/api/api.go                  # Entry point — loads env, wires server, starts Echo
│   ├── config/                         # Configuration settings
│   ├── db/
│   │   ├── db.go                       # MySQL connect, run migrations, seed, Paginate() helper
│   │   ├── migrations/                 # Versioned migration files + registration list
│   │   ├── models/                     # DB models
│   │   └── repositories/               # Data-access layer
│   ├── docker/
|   |   └── Dockerfile                  # Alpine Go image with CompileDaemon + swag
│   ├── pkg/
│   │   ├── closer/                     # Test DB closer
│   │   ├── go-migrations/              # Vendored go-migrations fork
│   │   ├── logging/                    # zerolog server logger
│   │   ├── middleware/                 # JWT auth middleware, logging middleware
|   |   └── validation/                 # Custom struct validator
│   ├── services/                       # Application/Business logic
│   └── tests/                          # Table-driven integration tests
│       ├── dependency_tests            # Dependency Tests
│       ├── helpers/                    # Test helpers (TestServer, TestCase and mock DB connection)
|       └── mocks/                      # Mock services
│
└── frontend/
    ├── .browserslistrc                 # Browser config
    ├── .env.example                    # Frontend environment variable template
    ├── .eslintrc.js                    # ESLint config
    ├── babel.config.js                 # Babel config
    ├── jsconfig.json                   # JS config
    ├── package.json                    # npm scripts: serve, build, lint
    ├── postcss.config.js               # PostCSS config
    ├── tailwind.config.js              # TailwindCSS config
    ├── vue.config.js                   # Vue config
    ├── public/                         # index.html & favicon.ico
    └── src/
        ├── App.vue                     # Entrypoint component
        ├── global-components.js        # Global components
        ├── main.js                     # Entrypoint JS
        ├── @core/                      # Core components from external libraries
        ├── assets/                     # CSS and Image assets
        ├── components/                 # Reusable Vue components
        ├── layouts/                    # Reusable Vue layouts
        ├── libs/                       # 3rd Party libraries
        ├── router/                     # Vue Router
        ├── services/                   # Axios API client + per-resource service modules
        ├── store/                      # Vuex modules
        └── views/                      # Page Views
```

---

## API Endpoints

All routes are registered under the root path (`/`). Every protected route requires a `Authorization: Bearer <token>` 
header.

**Auth**

| Method  | Path             | Description                               |
|---------|------------------|-------------------------------------------|
| POST    | `/auth/register` | Create a new user account                 |
| POST    | `/auth/login`    | Login, returns JWT access + refresh token |
| POST    | `/auth/refresh`  | Exchange a refresh token for new tokens   |
| GET     | `/auth/me`       | Get the authenticated user's profile      |
| POST    | `/auth/logout`   | Logout and invalidate the session         |

**File**

| Method | Path               | Description                  |
|--------|--------------------|------------------------------|
| GET    | `/files/:filepath` | Retrieve a file from storage |

**Class**

| Method | Path           | Params                                                                                                                                                         | Description                                     |
|--------|----------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------|
| GET    | `/classes`     | - page (Query) = Page number to return (0-indexed)<br/>- page_size (Query) = Number of items to return per page<br/> - search (Query) = Filter classes by name | Get a paginated list of classes ordered by name |
| GET    | `/classes/:id` | - id (Path) = The class ID to return                                                                                                                           | Get a class by its ID                           |

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