Various simple code to demonstrate concurrency in Golang

To run simple clock, run the golang code and connect to server with `nc`
```sh
$ go run simple-clock.go
$ nc localhost 8000
```

Run with multiple clients
```sh
$ go run simple-clock.go
$ go run netcat.go
$ go run netcat.go
$ go run netcat.go
```
