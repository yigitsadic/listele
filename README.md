# What is this?

This is an example project to implement unit-tests, interfaces, database integration tests with docker.

Application serves records found on `people` table as JSON list.

To run `docker-compose up`

`curl http://localhost:3035`

Run tests:

```
go test ./...
```

Run tests with database integration tests:

```
RUN_INTEGRATION_TESTS=YES go test ./...
```
