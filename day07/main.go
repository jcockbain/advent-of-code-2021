package main

import (
	_ "embed"
	"strconv"
	"strings"

	"fmt"
	"math"
)

const maxInt = math.MaxInt32

var (
	benchmark = false
)

//go:embed input.txt
var input string

func main() {
	p1 := part1()
	p2 := part2()

	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	nums := mapToInts(strings.Split(input, ","))
	ans := maxInt
	occupied := make(map[int]int, len(nums))
	for _, n := range nums {
		occupied[n]++
	}

	for pos := minSlice(nums); pos <= maxSlice(nums); pos++ {
		total := 0
		for start, count := range occupied {
			total += (abs(pos-start) * count)
		}
		ans = min(ans, total)
	}
	return ans
}

func part2() int {
	nums := mapToInts(strings.Split(input, ","))
	ans := maxInt
	occupied := make(map[int]int, len(nums))
	for _, n := range nums {
		occupied[n]++
	}
	for pos := minSlice(nums); pos <= maxSlice(nums); pos++ {
		total := 0
		for start, count := range occupied {
			total += getP2Distance(start, pos) * count
		}
		ans = min(ans, total)
	}
	return ans
}

func getP2Distance(start, end int) int {
	diff := abs(start - end)
	return diff * (diff + 1) / 2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minSlice(s []int) int {
	m := maxInt
	for _, x := range s {
		m = min(m, x)
	}
	return m
}

func maxSlice(s []int) int {
	m := 0
	for _, x := range s {
		if x > m {
			m = x
		}
	}
	return m
}

func mapToInts(l []string) []int {
	a := make([]int, len(l))
	for i, s := range l {
		a[i] = stringToInt(s)
	}
	return a
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
