package main

import (
	"fmt"
	"os"
	"strings"
)

type coordinate struct {
	x int
	y int
}

func main() {
	file, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(file), "\r\n")
	
	pipemap := make([]string, 0, len(lines[0]))

	var start coordinate;
	firststeps := make([]coordinate, 0, 2)

	// make the map, and find the starting point
	for i, line := range lines {
		pipemap = append(pipemap, line)
		for j, char := range line {
			if char == 'S' {
				start.x = i 
				start.y = j
			}
		}
	}

	// Find the two pipes adjacent to it
	var stuff []coordinate = []coordinate {{start.x - 1, start.y}, {start.x + 1, start.y}, {start.x, start.y - 1}, {start.x, start.y + 1}}



	for _, coord := range stuff {
		key := pipemap[coord.x][coord.y]	
		if (key == '|' &&( pipemap[coord.x + 1][coord.y] == 'S' || pipemap[coord.x - 1][coord.y] == 'S')) {
			firststeps = append(firststeps, coordinate{coord.x, coord.y})
		} else if (key == '-' &&( pipemap[coord.x][coord.y-1] == 'S' || pipemap[coord.x][coord.y + 1] == 'S')) {
			firststeps = append(firststeps, coordinate{coord.x, coord.y})
		} else if (key == 'L' &&( pipemap[coord.x][coord.y + 1] == 'S' || pipemap[coord.x - 1][coord.y] == 'S')) {
			firststeps = append(firststeps, coordinate{coord.x, coord.y})
		} else if (key == 'J' &&( pipemap[coord.x][coord.y - 1] == 'S' || pipemap[coord.x - 1][coord.y] == 'S')) {
			firststeps = append(firststeps, coordinate{coord.x, coord.y})
		} else if (key == '7' &&( pipemap[coord.x][coord.y - 1] == 'S' || pipemap[coord.x + 1][coord.y] == 'S')) {
			firststeps = append(firststeps, coordinate{coord.x, coord.y})
		} else if (key == 'F' &&( pipemap[coord.x][coord.y + 1] == 'S' || pipemap[coord.x + 1][coord.y] == 'S')) {
			firststeps = append(firststeps, coordinate{coord.x, coord.y})
		}
	}

	O := make([][]int, 0, len(lines))
	for i := 0 ; i < len(lines); i++ {
		O = append(O, make([]int, len(lines[0]), len(lines[0]))) 
	}
	O[start.x][start.y] = 1
	fmt.Println(part1(pipemap, start, firststeps[0], firststeps[1], O))

	ct := 0
	for i, line := range lines {
		inn := false
		for j, char := range line {
			if O[i][j] == 1{
				if strings.ContainsRune("|JL", char) || (pipemap[i][j] == 'S') {
					inn = !inn
				}
			} else {
				if inn {
					ct += 1
				}
			}
		}
	}

	fmt.Println(ct)

}

func part1(pipemap []string, start coordinate, curr1 coordinate, curr2 coordinate, part2map [][]int) int {

	steps := 1
	prev1 := coordinate{start.x, start.y}
	prev2 := coordinate{start.x, start.y}

	for (curr1 != curr2) {

		new1 := findnext(prev1, curr1, pipemap[curr1.x][curr1.y])
		new2 := findnext(prev2, curr2, pipemap[curr2.x][curr2.y])

		part2map[curr1.x][curr1.y] = 1
		part2map[curr2.x][curr2.y] = 1

		prev1 = coordinate{curr1.x, curr1.y}
		prev2 = coordinate{curr2.x, curr2.y}
		curr1 = coordinate{new1.x, new1.y}
		curr2 = coordinate{new2.x, new2.y}

		steps++
	}

	part2map[curr1.x][curr1.y] = 1
	part2map[curr2.x][curr2.y] = 1

	return steps
}


func findnext(prev coordinate, curr coordinate, pipe byte) coordinate {
	if (pipe == '|') {
		if curr.x + 1 == prev.x && curr.y == prev.y {
			return coordinate{curr.x - 1, curr.y}
		} else {
			return coordinate{curr.x + 1, curr.y}
		}
	} else if (pipe == '-') {
		if curr.x == prev.x && curr.y - 1 == prev.y {
			return coordinate{curr.x, curr.y + 1}
		} else {
			return coordinate{curr.x, curr.y - 1}
		}
	} else if (pipe == 'L') {
		if curr.x == prev.x && curr.y + 1 == prev.y {
			return coordinate{curr.x - 1, curr.y}
		} else {
			return coordinate{curr.x , curr.y + 1}
		}
	} else if (pipe == 'J') {
				if curr.x == prev.x && curr.y - 1 == prev.y {
			return coordinate{curr.x - 1, curr.y}
		} else {
			return coordinate{curr.x , curr.y - 1}
		}
	} else if (pipe == '7') {
				if curr.x == prev.x && curr.y - 1 == prev.y {
			return coordinate{curr.x + 1, curr.y}
		} else {
			return coordinate{curr.x, curr.y - 1}
		}
	} else {
		if curr.x == prev.x && curr.y + 1 == prev.y {
			return coordinate{curr.x + 1, curr.y}
		} else {
			return coordinate{curr.x, curr.y + 1}
		}
	}
}