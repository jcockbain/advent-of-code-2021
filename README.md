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
|     3 | 574µs   |
|     4 | 16.3ms  |
|     5 | 40.1ms  |
|     6 | 70.7µs  |
|     7 | 34.7ms  |
|     8 | 2.96ms  |
|     9 | 8.41ms  |
|    10 | 1.46ms  |
|    11 | 6.7ms   |
|    12 | 383ms   |
|    13 | 1.72ms  |
|    14 | 1.18ms  |
|    15 | 1.62s   |
|    16 | 344µs   |
|    17 | 2.22ms  |
|    18 | 4.77s   |
|    19 | 25.6s   |
|    20 | 644ms   |
| Total | 33.1s   |
