package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var mapDigits map[string]int = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9, "zero": 0}

func main() {
	var sum, sum2 int64 = 0,0

	// Reading the file
	file, _ := os.Open("input.txt")
	
	// Making a new scanner and split the thing line by line
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	
	// Looping through the lines
	for fileScanner.Scan() {

		line := fileScanner.Text()
	
		// Get all the digits
		stringslice, stringslice2 := make([]string, len(line)), make([]string, len(line)) // A slice of strings

		for index, char := range line {
			if !unicode.IsLetter(char) {
				stringslice[index], stringslice2[index] = string(char), string(char)
			}
		}
	
		// Get all the digit words
		for stringDigit, intDigit := range mapDigits {
			start := strings.Index(line, stringDigit)
			last := strings.LastIndex(line, stringDigit)
			if start != -1 {
				stringslice2[start] = strconv.Itoa(intDigit)
			}
			if last != -1 {
				stringslice2[last] = strconv.Itoa(intDigit)
			}
		}
	
		// Concatting the strings, turn it into an int.
		numString := strings.Join(stringslice, "")
		numString = string(numString[0]) + string(numString[len(numString) - 1])
		num, _ := strconv.ParseInt(numString, 0, 64)
		sum += num

		numString = strings.Join(stringslice2, "")
		numString = string(numString[0]) + string(numString[len(numString) - 1])
		num, _ = strconv.ParseInt(numString, 0, 64)
		sum2 += num
	
	}

	fmt.Println(sum, sum2)
	file.Close()

}