# Halucigenia Backend

A clean, production-ready REST API for managing bookmarks.

Built with Go, Gin, PostgreSQL, and Hexagonal Architecture principles.

------------------------------------------------------------------------

## 🚀 Tech Stack

-   Go 1.25+
-   Gin (HTTP framework)
-   PostgreSQL
-   dbmate (database migrations)
-   Docker / Docker Compose
-   Heroku-ready deployment

------------------------------------------------------------------------

## 🏗 Architecture

This project follows a clean / hexagonal architecture:

    internal/
    ├── app/            # Application use cases (business logic)
    ├── adapters/
    │   ├── http/       # HTTP transport layer (Gin handlers)
    │   └── postgres/   # PostgreSQL repositories
    ├── platform/
    │   ├── config/     # Configuration layer (env + yaml)
    │   └── db/         # Database connection setup

Flow:

    HTTP → Handler → Service → Repository → PostgreSQL

------------------------------------------------------------------------

## 📦 Features

-   Create bookmarks
-   List bookmarks
-   Delete bookmarks
-   Clean separation of concerns
-   Graceful shutdown
-   Environment-based configuration
-   Production-ready setup

------------------------------------------------------------------------

## 🗄 Database Schema

### bookmarks

  Column       Type        Constraints
  ------------ ----------- ------------------------
  id           UUID        Primary Key
  title        TEXT        NOT NULL
  url          TEXT        NOT NULL
  created_at   TIMESTAMP   NOT NULL DEFAULT now()

### Example SQL

``` sql
CREATE TABLE bookmarks (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
```

------------------------------------------------------------------------

## 🔧 Configuration

The application reads configuration in this order:

1.  Environment variables
2.  Config file (optional)
3.  Default values

### Required Environment Variables

    DATABASE_URL=postgresql://user:password@host:port/dbname?sslmode=disable
    PORT=3000
    ENV=production

On Heroku, `DATABASE_URL` is automatically provided.

------------------------------------------------------------------------

## 🐳 Running Locally with Docker

``` bash
docker compose up -d
```

Postgres will be available at:

    postgresql://postgres:example@localhost:5432/mypocket?sslmode=disable

------------------------------------------------------------------------

## 🧱 Database Migrations (dbmate)

Create a new migration:

``` bash
make migrate-new name=create_bookmarks_table
```

Run migrations:

``` bash
make migrate-up
```

Rollback:

``` bash
make migrate-down
```

------------------------------------------------------------------------

## ▶ Running the Server

``` bash
go run main.go
```

Server starts at:

    http://localhost:3000

------------------------------------------------------------------------

## 🌐 API Endpoints

### Create Bookmark

    POST /bookmarks

Body:

``` json
{
  "title": "Google",
  "url": "https://google.com"
}
```

------------------------------------------------------------------------

### List Bookmarks

    GET /bookmarks

------------------------------------------------------------------------

### Delete Bookmark

    DELETE /bookmarks/:id

------------------------------------------------------------------------

## 🛑 Graceful Shutdown

The application listens for:

-   SIGINT
-   SIGTERM

On shutdown:

-   Stops accepting new requests
-   Finishes in-flight requests
-   Closes database connections cleanly

------------------------------------------------------------------------

## 🚀 Deployment (Heroku)

1.  Set environment variables:

```{=html}
<!-- -->
```
    heroku config:set ENV=production

2.  Push:

```{=html}
<!-- -->
```
    git push heroku main

Heroku automatically provides:

    DATABASE_URL
    PORT

------------------------------------------------------------------------

## 🧪 Future Improvements

-   Authentication (JWT)
-   Tags
-   Pagination
-   Structured logging
-   Observability (OpenTelemetry)
-   Rate limiting

------------------------------------------------------------------------

## 📄 License

GPL-3.0
