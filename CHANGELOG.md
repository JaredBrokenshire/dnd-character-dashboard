## V0.0.2 - Auth

### [Added]

- Created roles and users DB tables, with user_roles link table
- Added response wrapper for api endpoints
- Added custom request validator package
- Added auth, role and user response objects
- Added auth request objects
- Created role and user repositories
- Created auth and user services, added to dependencies service
- Created auth handler with register, login, refresh, me and logout endpoints
- Added system config
- Created test server and test case helpers with mock db connection
- Created mock auth service for tests
- Added test running methods to base test
- Added tests for auth endpoints
- Added database seeder with default roles and users
- Added auth views (register and login) to frontend
- Created page header component and image background layout
- Added frontend route protection using auth service and store

## V0.0.1 - Project Initialisation

### [Added]

- Created CHANGELOG with update header formats
- Created dockerignore
- Initialised go.mod file
- Created API entrypoint
- Created echo server
- Created database init with env variables
- Added server logging package
- Added middleware package with logging and JWT Auth config
- Added server router with CORS and middleware configs
- Added db to server
- Added server initialisation to entrypoint
- Created database repositories and added them to server
- Added expose ports on the host machine for the server and database
- Added docker compose config to start database and api server
- Added go-migrations package and added migrations to db
- Added CI github workflow
- Added PR template
- Initialised Vue2 frontend with VueX store, VueRouter, TailwindCSS and ESLint
- Added frontend api service with axios
- Added base test in tests directory
- Added test database via docker compose

### [Changed]

- Updated default gitignore to better reflect tech stack


> - **[Added]** — New features, files, endpoints, components, or dependencies introduced in this version
> - **[Changed]** — Modifications to existing behaviour, UI, configuration, or structure that are not bug fixes
> - **[Fixed]** — Bug fixes and corrections to unintended behaviour
> - **[Removed]** — Features, files, endpoints, or dependencies that have been deleted