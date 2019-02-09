# HOWTO

``` shell
go test -coverprofile cover.out

curl localhost:8080/hello

Ctrl^C to shutdown server.

go tool cover -html=cover.out
```
You'll see that only the main and helloHandler code has been marked as covered.
