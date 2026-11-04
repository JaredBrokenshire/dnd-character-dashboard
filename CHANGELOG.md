## v0.1.2 `DATE`

### [Added]

- Added classes list endpoint with new handler, service and repo

### [Changed]

- Updated the character classes and subclasses fields to be pointers for consistency

### [Fixed]

- Fixed factory error handling

## v0.1.1 04/11/26

### [Added]

- Added a debounced search bar to the character list page, allowing users to filter their characters by name without triggering an API call on every keystroke
- Added `lodash` as a frontend dependency to provide the `debounce` utility used by the character search
- Added timestamps and `user_id` foreign key to the `Campaign` model, associating each campaign with an owning user (DM)
- Added `SlimUserResponse` and constructor
- Added new users, characters, races, subclasses and profile pictures to the seeders
- Added campaign routes and campaign service to frontend
- Added campaign list page and campaign card component

### [Changed]

- Combined the `uri` and `query` log fields in the request logging middleware into a single `uri` field formatted as `path?query`, reducing log verbosity and making request URLs easier to read at a glance
- Added left padding and larger text (`xl:text-lg`) to the `InputWrapper` component so search inputs align consistently with the surrounding UI at wider screen sizes
- Restructured the character list page to always render the search bar and page container regardless of loading state, moving the empty state card inside the main layout
- Updated the campaign list endpoint to return all campaigns visible to the authenticated user — campaigns where the user is the owner (`campaigns.user_id`) or has a character assigned to the campaign (`characters.user_id`)
- Updated asset seeder methods to take the model's id as the initial filename and use a random string for created file model

### [Fixed]

- Fixed CI checks running on push to production/development

### [Removed]

- Removed redundant CampaignsCard component from `/frontends/views/home/sections`

---

## v0.1.0 03/11/26

### [Added]

- Created project structure with backend and frontend directories, `.gitignore`, `.dockerignore`, `.env.example`, and Docker Compose configuration
- Built Go REST API using the Echo v4 framework with JWT authentication (register, login, refresh, logout, and current-user endpoints)
- Added character endpoints to list and retrieve characters with their associated race, classes, subclasses, campaign, and profile picture
- Added campaign list endpoint
- Added authenticated static file serving endpoint for character portraits and class logos
- Implemented GORM data models for User, Character, Campaign, Class, Subclass, Race, Role, and File
- Implemented versioned database migration system with 13 migrations covering all tables and relationships
- Implemented database seeders for roles, races, classes, subclasses, campaigns, characters, character profile pictures, and class logos; seeds a default admin user on first startup
- Added repository layer for users, characters, campaigns, and roles with a centralised `Repos` aggregator
- Added service layer for auth, characters, campaigns, and users with a `Dependencies` injection container
- Added JWT middleware for protecting routes and a request logging middleware using zerolog
- Added custom Echo request validator backed by `go-playground/validator`
- Added local filesystem file storage abstraction used by the file handler and seeders
- Added `SysConfig` singleton loaded from environment variables with coded defaults covering rate limiting, bcrypt cost, brute-force protection, and CORS origins
- Added `Paginate` helper on the DB package to handle `page` / `page_size` query parameters
- Added Swagger documentation generated via `swaggo/swag`, served at `/swagger/*`
- Added GitHub Actions CI pipeline running the backend build and full test suite against a real MySQL 8.0 instance, plus a frontend install and production build check
- Added table-driven integration test suite covering auth, character, campaign, and file endpoints with a shared `TestServer`, `TestCase` runner, and mock implementations for services and file storage
- Added local file storage unit tests and `try` package unit tests
- Built Vue 2 frontend with Vuex state management, Vue Router with auth guards (`requiresAuth`, `requiresAdmin`, `guestOnly`), and Tailwind CSS styling
- Added login and registration pages with form validation via `vee-validate`
- Added character list page displaying character cards with portrait, name, level, race, and class information
- Added home page with campaign and character summary cards
- Added 404 error page
- Added Axios API client with per-resource service modules for auth, characters, and files
- Added `InputWrapper`, `Card`, and `PageHeader` reusable Vue components
- Added ACL integration via `@casl/ability` and `@casl/vue`
- Added toast notification support via `vue-toastification`
- Added Docker Compose configuration for the full stack (API + MySQL) and a separate test database container
- Added Dockerfile using Alpine Go with `CompileDaemon` for live-reload during development and `swag init` on container start

### [Changed]

- Updated README with accurate project structure, full API endpoint reference, complete environment variable table, and corrected quick-start instructions

---

> - **[Added]** — New features, files, endpoints, components, or dependencies introduced in this version
> - **[Changed]** — Modifications to existing behaviour, UI, configuration, or structure that are not bug fixes
> - **[Fixed]** — Bug fixes and corrections to unintended behaviour
> - **[Removed]** — Features, files, endpoints, or dependencies that have been deleted
