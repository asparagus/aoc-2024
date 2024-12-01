package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const INPUT_SIZE = 1000

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput(path string) ([INPUT_SIZE]int, [INPUT_SIZE]int) {
	data, err := os.ReadFile(path)
	check(err)
	lines := strings.Split(string(data), "\n")
	var left [INPUT_SIZE]int
	var rght [INPUT_SIZE]int
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		parts := strings.Split(line, "   ")
		leftInt, err := strconv.Atoi(parts[0])
		check(err)
		rghtInt, err := strconv.Atoi(parts[1])
		check(err)
		left[i] = leftInt
		rght[i] = rghtInt
	}
	return left, rght
}

func part1(left [INPUT_SIZE]int, rght [INPUT_SIZE]int) int {
	diff := 0
	for i := 0; i < INPUT_SIZE; i++ {
		if left[i] > rght[i] {
			diff += left[i] - rght[i]
		} else {
			diff += rght[i] - left[i]
		}
	}
	return diff
}

func part2(left [INPUT_SIZE]int, rght [INPUT_SIZE]int) int {
	sim := 0
	var i int = 0
	var j int = 0
	for i < INPUT_SIZE && j < INPUT_SIZE {
		for i < INPUT_SIZE && left[i] < rght[j] {
			i++
		}
		for j < INPUT_SIZE && left[i] > rght[j] {
			j++
		}
		var acc int = 0
		var strt int = left[i]
		for j < INPUT_SIZE && left[i] == rght[j] {
			acc += left[i]
			j++
		}
		for i < INPUT_SIZE && left[i] == strt {
			sim += acc
			i++
		}
	}
	return sim
}

func main() {
	left, rght := readInput("inputs/day1.txt")
	sort.Ints(left[:])
	sort.Ints(rght[:])
	sol1 := part1(left, rght)
	sol2 := part2(left, rght)
	fmt.Println("Part 1:   ", sol1)
	fmt.Println("Part 2:   ", sol2)
}
