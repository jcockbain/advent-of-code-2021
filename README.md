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

// benchmark solution
go test --bench=BenchmarkMain

```

## Runtimes

The results are found using a `BenchmarkMain` benchmark in each solution. This table is generated by running the `benchmark.sh` script.

|  DAY  | RUNTIME |
|-------|---------|
|     1 | 115µs   |
|     2 | 673µs   |
|     3 | 549µs   |
|     4 | 20.6ms  |
|     5 | 36.8ms  |
|     6 | 71.1µs  |
|     7 | 37.9ms  |
|     8 | 2.93ms  |
|     9 | 9.18ms  |
|    10 | 1.44ms  |
|    11 | 6.62ms  |
|    12 | 441ms   |
|    13 | 1.81ms  |
|    14 | 971µs   |
|    15 | 1.55s   |
|    16 | 345µs   |
|    17 | 2.09ms  |
|    18 | 4.78s   |
|    19 | 25.9s   |
|    20 | 1.09s   |
|    21 | 5.33s   |
| Total | 39.2s   |
