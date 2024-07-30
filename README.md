# Task 6

## How to run?

> go to the root folder and run a command (in task_6 folder):
```bash
go run cmd/main/main.go
```

## How to test benchmarks?

> go to the root folder and run a command (in task_6 folder): 
```bash
go test -bench=. internal/tests/mutex_test.go
# and 
go test -bench=. internal/tests/perf_test.go
```