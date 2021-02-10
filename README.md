## CodelyTV - Go HTTP API - Hexagonal Architecture

This repository contains the code examples used on the CodelyTV course.

### Requirements

- Go v1.15+
- MySQL (see below).

### Contents

This project has been designed as a single Go module with multiple applications.
Each folder contains a completely functional application (binary) that can be executed isolated.

Each folder corresponds to one of the course lessons / videos:
1. [`01-01-your-first-http-endpoint`](./01-01-your-first-http-endpoint) - Nuestro primer endpoint HTTP en Go
1. [`01-02-using-gin`](./01-02-using-gin) - Usando Gin: nuestro primer framework
1. [`01-03-architectured-healthcheck`](./01-03-architectured-healthcheck) - Arquitecturando nuestro health check
1. [`02-01-post-course-endpoint`](./02-01-post-course-endpoint) - Implementando el endpoint de creación de curso
1. [`02-02-repository-injection`](./02-02-repository-injection) - Inyectando nuestro repositorio
1. [`02-03-controller-test`](./02-03-controller-test) - Testeando nuestro endpoint
1. [`02-04-domain-validations`](./02-04-domain-validations) - Añadiendo validaciones a nuestro dominio
1. [`03-01-mysql-repository-implementation`](./03-01-mysql-repository-implementation) - Implementando nuestro repositorio para MySQL
1. [`03-02-repository-test`](./03-02-repository-test) - Testeando nuestro repositorio
1. [`04-01-application-service`](./04-01-application-service) - Refactorizando el endpoint para extraer el Application Service
1. [`04-02-application-service-test`](./04-02-application-service-test) - Testeando el Application Service
1. [`04-03-command-bus`](./04-03-command-bus) - Unificando nuestros casos de uso: Command Bus
1. [`05-01-graceful-shutdown`](./05-01-graceful-shutdown) - Graceful shutdown
1. [`05-02-timeouts`](./05-02-timeouts) - Timeouts en operaciones asíncronas: repositorio
1. [`06-01-http-middlewares`](./06-01-http-middlewares) - Usando middlewares HTTP en Go
1. [`06-02-time-parse-in-go`](./06-02-time-parse-in-go) - El secreto mejor guardado de Go y sus fechas
1. [`06-03-gin-middlewares`](./06-03-gin-middlewares) - Implementando el middleware de recuperación de errores en Gin

### Usage

To execute the application from any lesson, just run:

```sh
export COURSE_LESSON=02-04-domain-validations; go run $COURSE_LESSON/cmd/api/main.go 
```

Replacing `COURSE_LESSON` value by any of the available ones.

#### Simple examples

Some lessons only contain a single `main.go` file with a few lines of code.
To run one of those lessons, just run:

```sh
export COURSE_LESSON=01-01-your-first-http-endpoint; go run $COURSE_LESSON/main.go 
```

#### MySQL & Docker

From `02-01-post-course-endpoint` on, the application on each directory relies
on a MySQL database. So, to simplify its execution, we've added a
`docker-compose.yaml` file with a MySQL container already set up.

To run it, just execute:

```sh
docker-compose up -d 
```

You can also use your own MySQL instance. Note that those applications
expects a MySQL instance to be available on `localhost:3306`,
identified by `codely:codely` and with a `codely` database.

To set up your database, you can execute the `schema.sql` file
present on the `sql` directory. It's automatically loaded if
you use the provided `docker-compose.yaml` file.

#### Tests

To execute all tests, just run:

```sh
go test ./... 
```

To execute only the tests present in one of the lessons, run:

```sh
go test ./02-04-domain-validations/... 
```