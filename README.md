# Rate API

A simple API for retrieving currency exchange rates.

## Build and Run

Build the Docker image by running `docker-compose build` in the root directory of the project.

Run the Docker container by running `docker-compose up` in the root directory of the project.

## Endpoints

### GET /api/v1/assets

Returns a list of all available assets.

### GET /api/v1/assets/:code

Returns a single asset by code.

## Development

The application is written in Go and uses the following libraries:

### Gin

Gin is a web framework written in Go. It is used for building web applications and APIs.

### godotenv

godotenv is a library for loading environment variables from a .env file.

### pgxpool

pgxpool is a library for connecting to a PostgreSQL database. It provides a connection pool and supports transactions.

## Tests

Tests are written using Go's built-in testing library. Tests can be run by executing `go test` in the root directory of the project.

## Deployment

The application can be deployed to any environment that supports Docker. See the section above for more information about the Docker compose file.

## Rate Limiting
The API uses a rate limiter to control traffic and prevent abuse.