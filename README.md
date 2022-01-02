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
|     1 | 89.8µs  |
|     2 | 568µs   |
|     3 | 399µs   |
|     4 | 9.21ms  |
|     5 | 26.2ms  |
|     6 | 59.4µs  |
|     7 | 29.5ms  |
|     8 | 2.16ms  |
|     9 | 5.26ms  |
|    10 | 1.08ms  |
|    11 | 4.8ms   |
|    12 | 280ms   |
|    13 | 1.31ms  |
|    14 | 744µs   |
|    15 | 1.36s   |
|    16 | 239µs   |
|    17 | 2.05ms  |
|    18 | 2.42s   |
|    19 | 15s     |
|    20 | 440ms   |
|    21 | 728ms   |
|    22 | 4.42ms  |
|    23 | 1.19s   |
|    24 | 15.2µs  |
|    25 | 1.56s   |
| Total | 23s     |
