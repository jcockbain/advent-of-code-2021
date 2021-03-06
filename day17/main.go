package main

import (
	_ "embed"
	"regexp"
	"strconv"

	"fmt"
)

var (
	benchmark = false
	re        = regexp.MustCompile(`target area: x=(-?\d+)..(-?\d+), y=(-?\d+)..(-?\d+)`)
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

type vel struct{ x, y int }

type pos struct{ x, y int }

func (p pos) isBetween(min pos, max pos) bool {
	return (p.x >= min.x) && (p.x <= max.x) && (p.y >= min.y) && (p.y <= max.y)
}

type area struct{ min, max pos }

func fireGun(v vel, target area) (bool, int) {
	p := pos{0, 0}
	highestY, enteredT := 0, false
	for p.x <= target.max.x && p.y >= target.min.y {
		p.x += v.x
		p.y += v.y
		if v.x > 0 {
			v.x--
		} else if v.x < 0 {
			v.x++
		}
		v.y--
		if p.y > highestY {
			highestY = p.y
		}
		if p.isBetween(target.min, target.max) {
			enteredT = true
		}
	}
	return enteredT, highestY
}

func part1() int {
	parts := re.FindStringSubmatch(input)
	minX, maxX, minY, maxY := toInt(parts[1]), toInt(parts[2]), toInt(parts[3]), toInt(parts[4])
	target := area{
		pos{minX, minY},
		pos{maxX, maxY},
	}
	maxH := 0
	for vx := 0; vx < maxX; vx++ {
		for vy := -50; vy < 200; vy++ {
			entersT, h := fireGun(vel{vx, vy}, target)
			if entersT {
				if h > maxH {
					maxH = h
				}
			}
		}
	}
	return maxH
}

func part2() int {
	parts := re.FindStringSubmatch(input)
	minX, maxX, minY, maxY := toInt(parts[1]), toInt(parts[2]), toInt(parts[3]), toInt(parts[4])
	target := area{
		pos{minX, minY},
		pos{maxX, maxY},
	}
	numberOfEntries := 0
	for vx := 0; vx <= maxX; vx++ {
		for vy := -200; vy < 200; vy++ {
			entersT, _ := fireGun(vel{vx, vy}, target)
			if entersT {
				numberOfEntries++
			}
		}
	}
	return numberOfEntries
}

func toInt(x string) int {
	i, err := strconv.Atoi(x)
	check(err)
	return i
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
