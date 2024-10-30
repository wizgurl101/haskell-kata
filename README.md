# Learning Go

Repo containing my solutions to leetcode. Goal is to learn Go (the programming language, not the boardgame \|>.<|/) :D

### How to run a go program:

```
go run .
```

### How to run a unit test

cd into a folder and use the following command:

```
go test .
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
