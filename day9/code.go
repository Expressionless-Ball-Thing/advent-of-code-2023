package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(file), "\r\n")

	predictSum, predictSum2 := 0, 0
	for _, line := range lines {
		stringarray := strings.Split(line, " ")
		sequence := make([]int, 0, len(stringarray))
		for _, item := range stringarray{
			num, _ := strconv.Atoi(item)
			sequence = append(sequence, num)
		}
		predictSum += predict(sequence, false)
		predictSum2 += predict(sequence, true)
	}

	var result = predictSum
	var result2 = predictSum2
	fmt.Println(result)
	fmt.Println(result2)
}

func predict(sequence []int, part2 bool) int {
	differences := make([]int, len(sequence)-1)
	allZeroes := true
	for i := 0; i < len(sequence)-1; i++ {
		differences[i] = sequence[i+1] - sequence[i]
		if differences[i] != 0 {
			allZeroes = false
		}
	}
	nextVal := 0
	if !allZeroes {
		nextVal = predict(differences, part2)
	}
	if part2 {
		return sequence[0] - nextVal
	}
	return sequence[len(sequence)-1] + nextVal
}