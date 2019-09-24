# Koala POS Backend

## Build Requirements

- Go 1.13+
- Docker

## Building

Change directory to the backend folder and run `go run mage.go`.

```shell
cd $REPO/backend
go run mage.go // Runs the build target in magefile.go
```

### Building for Docker

Run the `buildInDocker` target.

```shell
go run mage.go buildInDocker
```

## Running Dev Server

Make sure to build the application for Docker deployment first.

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
