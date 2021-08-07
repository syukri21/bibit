## Requirements

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Feature

- Clean code
- Docker and docker-compose
- ORM with gorm
- Unit testing

## Setup

### Building and running the dockerized codebase

1. Clone the repository.

1. Build all containers using `docker-compose build --no-cache`.

1. Run all containers using `docker-compose up`.

   **Notes**:
   When running docker compose, it will create new volume(s). If change(s) are made to `package.json`, it won't be detected by the container.
   I recommend:

    - To remove the volume `docker volume prune`, or
    - Update package manually by doing `docker exec -it <container id> sh` and `npm install`

   I use the `--no-cache` flag so that new npm packages gonna be installed.

   **Tips**:

    - Use `--build` in `docker-compose` to force update the docker image created, e.g. `docker-compose up --build`

1. The docker and docker compose will setup all requirements on the fly, and please provide coffee as this may take a few minutes.

1. If have finished it, that means you can use this server with all databases and data ready to use.
### Cleaning up

1. When you're done, `Ctrl-C` in the main `docker-compose up` window to terminate the running processes.

1. Run `docker-compose down` to stop and remove containers.

