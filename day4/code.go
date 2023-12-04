package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

func sumslice(array[] int) int {
    result := 0
    for _, v := range array {
        result += v
    }
    return result
}

func main() {

	file, _ := os.ReadFile("input.txt")
	rows := strings.Split(string(file), "\n")
	
	gameWins := make([]int, 0, len(rows))

	sum := 0.0

	for _, line := range rows {
		mySet1 := mapset.NewSet[string]()
		mySet2 := mapset.NewSet[string]()
		slice := strings.Split(line[strings.Index(line, ":")+1:], "|")

		for _, item := range regexp.MustCompile(`\d+`).FindAllString(slice[0], -1) {
			mySet1.Add(string(item))
		}
		for _, item := range regexp.MustCompile(`\d+`).FindAllString(slice[1], -1) {
			mySet2.Add(string(item))
		}

		wins := mySet1.Intersect(mySet2).Cardinality()
		gameWins = append(gameWins, wins)
		if wins > 0 {
			sum += math.Pow(2, float64(wins - 1))
		}
	}

	for i := len(gameWins) - 1; i >= 0; i-- {
		gameWins[i] = sumslice(gameWins[i: (i + gameWins[i] + 1)])
	}

	fmt.Println(sum)
	fmt.Println(sumslice(gameWins) + len(gameWins))
}