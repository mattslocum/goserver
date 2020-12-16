# goserver
First try at Golang. Pretty fun.

Developed with GoLang 1.15

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

## Building
Because cloud native apps are awesome, the final deployable as a docker image.

```shell script
$ docker build -t goserver . 
```
Run locally:
```shell script
docker run -p 8080:8080 --rm -i goserver:latest
```

