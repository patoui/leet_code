# Leet Code

Repository of my work on leet code questions and solutions

## Running the solutions

Change into the problems directory and run `go run main.go`, as an example:

```
cd 28_find_the_index_of_the_first_occurrence_in_a_string/
go run main.go
```

### Running the profiler

**Pre-requisites**
- Go bin added to path. Example on Linux, add `export PATH = $PATH:$HOME/go/bin`
- `pprof` must be installed: `go install github.com/google/pprof@latest`

Steps:
- change into the problems directory, example `cd 28_find_the_index_of_the_first_occurrence_in_a_string/`
- run the command to generate the benchmark output `go test -cpuprofile cpu1.prof -bench .`
- see visual output `pprof -http=:8080 cpu1.prof`

This should have opened a browser with the callgraph