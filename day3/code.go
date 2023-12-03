package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	locations := [8][2]int {{-1, -1},{-1, 0},{-1, 1},{0, -1},{0, 1},{1, -1},{1, 0},	{1, 1}}

	file, _ := os.ReadFile("input.txt")
	rows := strings.Split(string(file), "\n")
	eng := make([][][]int, len(rows), len(rows))

	sum1 := 0
	sum2 := 0
	// Scan the lines for the number's positions
	for i, line := range rows {

		re := regexp.MustCompile(`\d+`)
		numbers := re.FindAllStringIndex(line, -1)
		for j, item := range numbers {
			number, _ := strconv.ParseInt(line[item[0]:item[1]], 0, 64)
			numbers[j] = append(item, int(number))
		}
		eng[i] = numbers
	}

	// Scan the lines again for the symbol's location.s
	for i, line := range rows {

		for j, char := range line {
			if !unicode.IsDigit(char) && char != '.' && unicode.IsPrint(char) {
				temparray := make([]int, 0, 8)
				for _, locate := range locations {
					newx, newy := i + locate[0], j + locate[1]

					if (newx >= 0 && newx < len(rows) && newy >= 0 && newy < len(line)) {
						for k, array := range eng[newx] {
							if newy >= array[0] && newy < array[1] {
								// Found a match
								eng[newx][k] = []int{-1, -1, 0}
								sum1 += array[2]
								temparray = append(temparray, array[2])
							}
						}

					}
				}
				if (len(temparray) == 2) {
					sum2 += temparray[0] * temparray[1]
				}

			}
		}
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}