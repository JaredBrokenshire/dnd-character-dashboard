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
├── .env.example                        # Root environment variable template
├── docker-compose.yml                  # Full stack: api + database
├── docker-compose-test.yml             # Test database container only
│
├── backend/
│   ├── cmd/api/api.go                  # Entry point — loads env, wires server, starts Echo
│   ├── api/
│   │   ├── server.go                   # Server struct (Echo + DB + Repos + Services)
│   │   └── routes/                     # Route registration
│   ├── db/
│   │   ├── db.go                       # MySQL connect, run migrations, seed, Paginate() helper
│   │   ├── migrations/                 # Versioned migration files + registration list
│   │   └── repositories/               # Data-access layer
│   ├── docker/
|   |   └── Dockerfile                  # Alpine Go image with CompileDaemon + swag
│   ├── pkg/
│   │   ├── middleware/                 # JWT auth middleware, logging middleware
│   │   ├── logging/                    # zerolog server logger
│   │   └── go-migrations/              # Vendored go-migrations fork
│   └── tests/                          # Table-driven integration tests
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
        ├── main.js                     # Entrypoint JS
        ├── assets/                     # CSS and Image assets
        ├── components/                 # Reusable Vue components
        ├── router/                     # Vue Router
        ├── services/                   # Axios API client + per-resource service modules
        ├── store/                      # Vuex modules
        └── views/                      # Page Views
```

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