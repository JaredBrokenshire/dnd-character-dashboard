## V0.0.8 - Race List `DATE`

### [Added]

- Added a new `races`, `traits`, `race_traits`, and `race_variants` tables with migrations and corresponding models
- Added a new race handler with a paginated list endpoint
- Added a race service and repo

## V0.0.7 - Class List 08/11/26

### [Added]

- Added a new `classes` table with a migration and corresponding model
- Added a new class handler with a paginated list endpoint
- Added a class service and repo
- Added paginate method to database layer
- Added file and class factories for creating test models

## V0.0.6 - File Storage 08/11/26

### [Added]

- Added local file storage with new dependencies service
- Added protected file endpoint with response objects

## V0.0.5 - Home Page 08/11/26

### [Changed]

- Updated the home page to include navigation cards

## V0.0.4 - Auth 08/11/26

### [Added]

- Created `users`, `roles` and `user_roles` tables with migrations
- Created user model and role model
- Added response wrapper for endpoint responses
- Added user, role and auth response objects
- Added auth register request object
- Added auth handler with register route, user and role repos
- Added auth service with JWT middleware
- Added auth register endpoint tests
- Added login endpoint with new service method
- Added refresh endpoint with new service method
- Added me endpoint to return the current user
- Added logout endpoint
- Added seeder to set inital users and roles
- Added frontend login/register pages with route protection
- Added vee-validate for form validation and FeatherIcons
- Added page layout with image background
- Added 404 page
- Added page header

### [Changed]

- Change github workflow to run on pr instead of push

## V0.0.3 - Server Logging 08/11/26

### [Added]

- Added logging package with middleware function

## V0.0.2 - Frontend Setup 08/11/26

### [Added]

- Initialised Vue2 project with tailwind, vue-router and ESLint
- Added temp home page
- Added API service to send http requests with axios

## V0.0.1 - Backend Setup 08/11/26

### [Added]

- Added GitHub CI workflow
- Added .dockerignore, .gitignore and .env.example
- Added README and CHANGELOG
- Added go main method in backend directory
- Added echo server with gorm DB
- Added custom validator to server
- Added router with middleware CORS config
- Added base test with test case helpers (TestServer, TestCase)
- Added config package
- Added db init method with repos
- Added services package with dependencies handler
- Added backend dockerfile and docker compose yml, with test file for starting a test db
- Added go migrations package

> - **[Added]** — New features, files, endpoints, components, or dependencies introduced in this version
> - **[Changed]** — Modifications to existing behaviour, UI, configuration, or structure that are not bug fixes
> - **[Fixed]** — Bug fixes and corrections to unintended behaviour
> - **[Removed]** — Features, files, endpoints, or dependencies that have been deleted