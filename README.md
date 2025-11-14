# Split-wise-go Starter Kit

This is a minimal "starter kit" repository for building a Go API with Chi, Pgx, and Postgres. It's designed to get students past the "blank canvas" stage and provide a solid foundation for their projects.

## Features

* **Go Backend:** Using `chi` for routing and `pgx` (v5) for Postgres.
* **Postgres Database:** Managed via `docker-compose`.
* **Database Migrations:** Using `golang-migrate`.
* **Simple Makefile:** Contains all common commands (`run`, `migrate-up`, etc.).
* **Environment Config:** Uses `.env` files for easy configuration.

## Prerequisites

Before you begin, you need to have the following tools installed on your system:

1.  [**Go**](https://go.dev/doc/install) (version 1.22+ recommended)
2.  [**Docker & Docker Compose**](https://docs.docker.com/get-docker/)
3.  [**golang-migrate CLI**](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

You can install the `golang-migrate` CLI with Go:

```sh
go install -v [github.com/golang-migrate/migrate/v4/cmd/migrate@latest](https://github.com/golang-migrate/migrate/v4/cmd/migrate@latest)
````

Alternatively, you can run `make install-migrate` from this repo.

**IMPORTANT:** Make sure that your Go bin directory (e.g., `$(go env GOPATH)/bin`) is in your system's `PATH` variable so you can run the `migrate` command.

## Quick Start

1.  **Clone the repository:**

    ```sh
    git clone [https://github.com/rengas/splitwise-go.git](https://github.com/rengas/splitwise-go.git)
    cd splitwise-go
    ```

2.  **Create your environment file:**
    Copy the example file to your own `.env` file.

    ```sh
    cp .env.example .env
    ```

    (The defaults in this file work perfectly with the `docker-compose.yml`.)

3.  **Start the Postgres database:**
    This command will start a Postgres container in the background.

    ```sh
    make docker-up
    ```

4.  **Run the database migrations:**
    This command will read the `.env` file, connect to the database, and run the SQL files in the `migrations/` directory.

    ```sh
    make migrate-up
    ```

    You should see an output indicating `000001_create_users_table.up.sql` was applied.

5.  **Run the application:**

    ```sh
    make run
    ```

    The server will start on `http://localhost:8080`.

6.  **Test the server:**
    Open a new terminal and test the `/ping` healthcheck endpoint:

    ```sh
    curl http://localhost:8080/ping
    ```

    If everything is working, you should see:

    ```json
    {"status": "ok", "message": "pong!"}
    ```

You are now ready to start building your API!

## Makefile Commands

* `make run`: Starts the Go web server (listens on port from `.env`).
* `make docker-up`: Starts the Postgres container via `docker-compose`.
* `make docker-down`: Stops and removes the Postgres container.
* `make migrate-new name=...`: Creates a new SQL migration file (e.g., `make migrate-new name=create_posts_table`).
* `make migrate-up`: Applies all pending "up" migrations.
* `make migrate-down`: Reverts the last "up" migration.
* `make install-migrate`: Installs the `golang-migrate` CLI tool for you.

<!-- end list -->

```
```