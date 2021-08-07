## Requirements

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Feature

- Clean code
- Docker and docker-compose
- ORM with gorm
- Unit testing
- Rest API
- gRPC

## Setup

### Building and running the dockerized codebase

1. Clone the repository.
1. Build all containers using `docker-compose build --no-cache`.
1. Run all containers using `docker-compose up`.
1. The docker and docker compose will setup all requirements on the fly, and please provide coffee as this may take a few minutes.
1. If have finished it, that means you can use this server with all databases and data ready to use.


### Building and running without Docker
1. install `golang:1.15`, `postgresql`, `protoc or protobuf`
1. run in cmd `make dev`

### Cleaning up

1. When you're done, `Ctrl-C` in the main `docker-compose up` window to terminate the running processes.

1. Run `docker-compose down` to stop and remove containers.

### Generate GRPC Protobuf
```bazaar
protoc --go_out=plugins=grpc:. src/models/proto/*.proto
```
