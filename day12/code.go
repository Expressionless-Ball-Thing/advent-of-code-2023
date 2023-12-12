package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


type state struct {
	pattern string
	numbers string
}

var cache = make(map[state]int)


func main() {

	file, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n")

	var patterns []string 
	var numbers [][]uint8

	for _, line := range lines {
		parts := strings.Split(line, " ")
		patterns = append(patterns, parts[0])
		var group []uint8
		for _, num := range strings.Split(parts[1], ",") {
			num, _ := strconv.Atoi(num)
			group = append(group, uint8(num))
		}
		numbers = append(numbers, group)
	}

	fmt.Println(solve(patterns, numbers, false))
	fmt.Println(solve(patterns, numbers, true))
}

func setCache(pattern string, numbers []uint8, value int) int {
	cache[state{pattern, string(numbers)}] = value
	return value
}

func count(pattern string, numbers []uint8) int {
	if len(pattern) == 0 && len(numbers) == 0 {
		return 1
	}

	if len(pattern) == 0 {
		return 0
	}

	if value, ok := cache[state{pattern, string(numbers)}]; ok {
		return value
	}

	if pattern[0] == '.' {
		res := count(pattern[1:], numbers)
		return setCache(pattern, numbers, res)
	}

	// cut branches
	var sum int
	for _, n := range numbers {
		sum += int(n)
	}
	if len(pattern) < sum {
		res := 0
		return setCache(pattern, numbers, res)
	}

	if pattern[0] == '?' {
		res := count(pattern[1:], numbers) + count("#"+pattern[1:], numbers)
		return setCache(pattern, numbers, res)
	}

	if pattern[0] == '#' {
		if len(numbers) == 0 {
			res := 0
			return setCache(pattern, numbers, res)
		}

		n := numbers[0]
		indexDot := strings.Index(pattern, ".")
		if indexDot == -1 {
			indexDot = len(pattern)
		}
		if indexDot < int(n) {
			res := 0
			return setCache(pattern, numbers, res)
		}
		remaining := pattern[n:]
		if len(remaining) == 0 {
			res := count(remaining, numbers[1:])
			return setCache(pattern, numbers, res)
		}
		if remaining[0] == '#' {
			// fail
			res := 0
			return setCache(pattern, numbers, res)
		}
		res := count(remaining[1:], numbers[1:])
		return setCache(pattern, numbers, res)
	}
	panic("unreachable")
}

func unfoldPattern(pattern string) string {
	var res = pattern
	for i := 0; i < 4; i++ {
		res = res + "?" + pattern
	}
	return res
}

func unfoldNumbers(numbers []uint8) []uint8 {
	var res []uint8
	for i := 0; i < 5; i++ {
		res = append(res, numbers...)
	}
	return res
}

func solve(patterns []string, numbers [][]uint8, unfold bool ) int {

	var res int
	for i := 0; i < len(patterns); i++ {
		pattern, numbers := patterns[i], numbers[i]
		if unfold {
			pattern = unfoldPattern(pattern)
			numbers = unfoldNumbers(numbers)
		}
		res += count(pattern, numbers)
	}

	return res
}