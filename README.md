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

// defaults to using input.txt in the current dir
go run main.go

// run with custom input (I normally use the test input given with the problem)
go run main.go test1.txt

```

## Runtimes

The results are found using a `BenchmarkMain` benchmark in each solution. This table is generated by running the `benchmark.sh` script.

|  DAY  | RUNTIME |
|-------|---------|
|     1 | 132µs   |
|     2 | 707µs   |
|     3 | 560µs   |
|     4 | 14.1ms  |
|     5 | 41.3ms  |
|     6 | 72.2µs  |
|     7 | 320ms   |
|     8 | 3.14ms  |
|     9 | 7.69ms  |
|    10 | 1.54ms  |
|    11 | 6.9ms   |
|    12 | 2.32s   |
| Total | 2.71s   |
