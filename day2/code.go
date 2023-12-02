package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	idSum, power, count  := 0,0,1

	// Scan the lines
	for fileScanner.Scan() {

		line := fileScanner.Text()

		// Split the sets
		slice := strings.Split(line[strings.Index(line, ":")+1:], ";")
		fucked := false

		var minballcount map[string]int = map[string]int{"blue": 0, "red": 0, "green": 0}

		// Split the balls found in each set.
		for _, s := range slice {
			set := strings.Split(s, ",")
			for _, t := range set {
				num, _ := strconv.Atoi(strings.Split(strings.TrimSpace(t), " ")[0])
				color := strings.Split(strings.TrimSpace(t), " ")[1]
				if !(color == "red" && num <= 12 || color == "green" && num <= 13 || color == "blue" && num <= 14) {
					fucked = true
				}
				if (num > minballcount[color]) {
					minballcount[color] = num
				}
			}
		}

		// multiply the min ball counts
		temppower := 1
		for _, s := range minballcount {
			temppower *= s
		}
		power += temppower

		if !fucked { idSum += count }
		count += 1
	}

	fmt.Println("part1", idSum)
	fmt.Println("part2", power)

	file.Close()
}