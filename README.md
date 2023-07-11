## Graceful shutdown golang example

This projects shows an example of how to handle a server "GracefulShutdown" without killing currently requests


```sh
go run main.go

netstat -lpan | grep -i 8080 # get program pid
## ex: tcp6       0      0 :::8080                 :::*                    LISTEN      510178/main
## PID = 510178

# Use low request time to see waiting the request before context deadline timeout
curl -v localhost:8080/block/5


# Use higher wait time request to see server shutting down without wait the request finish
curl -v localhost:8080/block/5

kill -s SIGTERM $PID # in this example 510178


```