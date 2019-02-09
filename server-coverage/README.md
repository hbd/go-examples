# HOWTO

Get coverage of handlers. Could be used in getting code coverage for integration and system tests.

``` shell
go test -coverprofile cover.out

curl localhost:8080/hello

Ctrl^C to shutdown server.

go tool cover -html=cover.out
```

You'll see that only the main and helloHandler code has been marked as covered.
