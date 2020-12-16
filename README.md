# goserver
First try at Golang. Pretty fun.

Developed with GoLang 1.15

### Routing
All routes have a matching directory structure in /routes.
I made a master router.go file that coordinates all the routing.
I'm tempted to try a different routing structure where the http router is passed into each route handler. That way the file owns both the route location as well as the logic.

## Running Locally
```shell script
$ go run main.go
```

## Endpoints
POST /hash
* data: password=&lt;password>
* Example: `curl --data "password=hunter2" http://localhost:8080/hash`
* Return: key to retrieve the password hash when it is ready

GET /hash/&lt;id>
* id: The key returned from creating a new password hash.
* Example: curl http://localhost:8080/hash/1
* It may take about 5 seconds before the hash is available.

/stats
* Returns the stats for the hash count and average timing in microseconds

/shutdown
* Graceful shutdown of the app.
* The app also accepts SIGTERM Interrupt

## Env Variables
LOG_LEVEL
* DEBUG | INFO | WARN | ERROR
* Case sensitive. Defaults to Debug.

PORT
* Defaults to 8080

## Run Tests
```shell script
go test ./...
```

## Building
Because cloud native apps are awesome, the final deployable as a docker image.

```shell script
$ docker build -t goserver . 
```
Run the build locally:
```shell script
docker run -p 8080:8080 --rm -i goserver:latest
```

