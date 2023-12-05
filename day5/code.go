package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.ReadFile("input.txt")
	fmt.Println(part1(string(file)))
	// fmt.Println(part2(string(file)))

	// aochelper.AocMain(aochelper.Solution{Day: 5, Part1: part1, Part2: part2})
}

type mapData struct {
	src uint64
	dst uint64
	len uint64
}

type mapper struct {
	data []mapData
}

type multiMapper struct {
	mappers []mapper
}

func (m *mapper) mapFrom(src uint64) uint64 {
	for _, d := range m.data {
		if src >= d.src && src < d.src + d.len {
			return d.dst + (src - d.src)
		}
	}
	return src
}

func (m *multiMapper) mapFrom(src uint64) uint64 {
	for _, mapper := range m.mappers {
		src = mapper.mapFrom(src)
	}
	return src
}

func parseInput(input string) ([]uint64, multiMapper) {
	
	lines := strings.Split(input, "\r\n\r\n")
	// parse the Seed numbers
	seedStrings := strings.Split(lines[0], " ")
	seeds := make([]uint64, 0, len(seedStrings)-1)
	for _, seed := range seedStrings[1:] {
		s, _ := strconv.ParseUint(seed, 10, 64)
		seeds = append(seeds, s)
	}

	// Parse the mappers.
	mappers := make([]mapper, 7, 7)
	for _, mapping := range lines[1:] {
		splitted := strings.Split(mapping, "\n")[1:]
		for _, line := range splitted {
			parts := strings.Split(line, " ")
			dst, _ := strconv.ParseUint(parts[0], 10, 64)
			src, _ := strconv.ParseUint(parts[1], 10, 64)
			mlen, _ := strconv.ParseUint(parts[2], 10, 64)
			mappers[len(mappers)-1].data = append(mappers[len(mappers)-1].data, mapData{src, dst, mlen})			
		}
	}


	return seeds, multiMapper{mappers}
}

func part1(input string) any {
	seeds, mappers := parseInput(input)
	var minLocation uint64 = math.MaxUint64
	for _, seed := range seeds {
		minLocation = min(minLocation, mappers.mapFrom(seed))
	}
	return minLocation
}

func part2(input string) any {
	seeds, mappers := parseInput(input)
	var minLocation uint64 = math.MaxUint64
	for i := 0; i < len(seeds); i += 2 {
		end := seeds[i] + seeds[i+1]
		for s := seeds[i]; s < end; s++ {
			minLocation = min(minLocation, mappers.mapFrom(s))
		}
	}
	return minLocation
}