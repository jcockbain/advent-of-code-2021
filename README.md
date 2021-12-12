# 🎄 advent-of-code-2021 🎄

![Go](https://github.com/jcockbain/advent-of-code-2021/workflows/Go/badge.svg)
![Go Report Card](https://goreportcard.com/badge/github.com/jcockbain/advent-of-code-2021)

Solutions to 2021 Advent of code. 

## Summary 

[Advent of Code](https://adventofcode.com/) is an annual advent-calendar of programming puzzles. Here are my 2021 solutions in Golang. 

## Running the Code

To fetch the input, and create a template dir for each day: 

```shell
./new_day 2021 {day1}
```

To then run the solutions: 

```go
// using day1 as an example
cd day01

// run the binary
go run main.go

// run tests
go test

//benchmark solution
go test --bench=BenchmarkMain

```

## Runtimes

The results are found using a `BenchmarkMain` benchmark in each solution. This table is generated by running the `benchmark.sh` script.

|  DAY  | RUNTIME |
|-------|---------|
|     1 | 101µs   |
|     2 | 715µs   |
|     3 | 511µs   |
|     4 | 17.7ms  |
|     5 | 42.8ms  |
|     6 | 69.8µs  |
|     7 | 34.3ms  |
|     8 | 3.04ms  |
|     9 | 8.05ms  |
|    10 | 1.45ms  |
|    11 | 6.57ms  |
|    12 | 365ms   |
| Total | 480ms   |
