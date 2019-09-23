# Koala POS Backend

## Build Requirements

- Go 1.13+
- Make
- Docker

## Building

Change directory to the backend folder and run `make build-in-docker`.

```shell
cd $REPO/backend
make build-in-docker
```

The command will download the golang Docker image and build the application
inside. The executable will be in the bin directory.

## Running Dev Server

To run the application in development mode, cd into the docker folder and run
`docker-compose up -d`.

```shell
cd $REPO/backend/docker
docker-compose up -d
```

Once the application and database are started, the server will be available
at `http://localhost:8080'. You can use `docker ps` to make sure both the
application and database containers started.

If you recompile the application, run `docker-compose restart pos-backed` to
restart the application container with the new binary.

Database data is stored in a Docker volume and won't be destroyed unless you run
`docker-compose down`. To stop the containers, use `docker-compose stop`.
