package main

import (
	"fmt"
	"image"
	"math"
	"os"
	"strings"
)


func main() {
	file, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(file), "\r\n")

	var galaxies []image.Point;
	var emptyColumns []int 
	var emptyRows []int 

	for i, line := range lines {
		
		// Find empty columns
		if (!strings.Contains(line, "#")) {
			emptyRows = append(emptyRows, i)
			continue
		}
		
		// Get all the galaxies
		for j, char := range line {
			if char == '#' {
				galaxies = append(galaxies, image.Point{i, j})
			}
		}
	}

	// Find all empty columns
	for j := range lines[0] {
		found := false
		for i := range lines {
			if lines[i][j] == '#' {
				found = true
			}
		}
		if found { continue }
		emptyColumns = append(emptyColumns, j)
	}

	getstuff(galaxies, 2, emptyRows, emptyColumns)
	getstuff(galaxies, 1000000, emptyRows, emptyColumns)
}


func getstuff(galaxies []image.Point, distanceMult int, emptyRows []int, emptyColumns []int) {
	galaxiesCopy := make([]image.Point, len(galaxies))
	copy(galaxiesCopy, galaxies)

	sum := 0

	for i, galaxy := range galaxiesCopy {
		x, y := galaxy.X, galaxy.Y
		dx, dy := 0, 0
		for _, stuff := range emptyRows {
			if (stuff < x) {
				dx += distanceMult - 1
			}
		}
		for _, stuff := range emptyColumns {
			if (stuff < y) {
				dy += distanceMult - 1
			}
		}

		galaxiesCopy[i].X += dx 
		galaxiesCopy[i].Y += dy
	}

	for i, galaxy1 := range galaxiesCopy {
		for _, galaxy2 := range galaxiesCopy[i + 1:] {
			sum += int(math.Abs(float64(galaxy1.X - galaxy2.X))) + int(math.Abs(float64(galaxy1.Y - galaxy2.Y)))
		}
	}

	fmt.Println(sum)
}