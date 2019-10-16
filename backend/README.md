# Koala POS Backend

## Build Requirements

- Go 1.13+
- Docker

## Build and development scripts

All build and dev helper scripts are implemented using
[Mage](https://magefile.org/). You can get a list of targets by running `go run
mage.go -l`. Run a Mage target with `go run mage.go TARGET`. All Mage targets
must be ran inside the backend folder.

## Building

Run the `buildInDocker` target.

```shell
cd $REPO/backend
go run mage.go buildInDocker
```

## Running

Use either the `runDev` or `runDevLogs` Mage target.

```shell
cd $REPO/backend
go run mage.go runDev
# or
go run mage.go runDevLogs
```

`runDevLogs` will display application logs from the database and API to the
terminal. I would recommend using this target for development. `runDev` will
start the containers and exit not showing any logs.

Once the application and database are started, the server will be available at
`http://localhost:8080` You can use `docker ps` to make sure both the
application and database containers started.

If you recompile the application, run `go run mage.go restartDev` to restart the
application container with the new binary.

Database data is stored in a Docker volume and won't be destroyed unless you run
`go run mage.go StopDevClean`. To stop the containers, use `go run mage.go
StopDev`.
