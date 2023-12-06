package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(file), "\n")
	times := regexp.MustCompile(`\d+`).FindAllString(lines[0], -1)
	distances := regexp.MustCompile(`\d+`).FindAllString(lines[1], -1)
	
	product := 1.0

	for i := 0; i < len(times); i++ {
		t0, _ := strconv.ParseFloat(times[i], 64)
		d, _ := strconv.ParseFloat(distances[i], 64)
		min, max := math.Floor((-t0 + math.Sqrt(math.Pow(t0, 2) - 4*d)) / -2), math.Ceil((-t0 - math.Sqrt(math.Pow(t0, 2) - 4*d)) / -2)
		// fmt.Println(max - min + 1)
		fmt.Println((-t0 + math.Sqrt(math.Pow(t0, 2) - 4*d)) / -2, (-t0 - math.Sqrt(math.Pow(t0, 2) - 4*d)) / -2)
		product *= (max - min - 1)
	}

	fmt.Println(product)

	t0, _ := strconv.ParseFloat(strings.Join(times, ""), 64)
	d, _ := strconv.ParseFloat(strings.Join(distances, ""), 64)
	min, max := int64(math.Floor((-t0 + math.Sqrt(math.Pow(t0, 2) - 4*d)) / -2)), int64(math.Ceil((-t0 - math.Sqrt(math.Pow(t0, 2) - 4*d)) / -2))
	fmt.Println(max - min - 1)

}