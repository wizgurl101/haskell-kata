# Learning Go

### How to run a go program:

```
go run .
```

### How to run all unit tests

```
go test ./...
```

### How to run unit tests with a check for race condition

```
go test -race ./...
```

### How to run tests with generating a unit test coverage profile

```
go test ./... -coverprofile=coverage.out
```

### Use this command to display coverage info in terminal

```
go tool cover -func=coverage.out
```

### Use this command to display coverage info in HTML format in a browser

```
go tool cover -html=coverage.out
```
