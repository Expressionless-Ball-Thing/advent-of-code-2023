package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)
type mapData struct {
	src int64
	dst int64
	len int64
}

type mapper struct {
	data []mapData
}

type multiMapper struct {
	mappers []mapper
}

func main() {
	
	file, _ := os.ReadFile("input.txt")
	seeds, mappers := parseInput(string(file))
	fmt.Println(part1(seeds, mappers))
	fmt.Println(part2(seeds, mappers))

}

func parseInput(input string) ([]int64, multiMapper) {
	
	lines := strings.Split(input, "\r\n\r\n")
	// parse the Seed numbers
	seedStrings := strings.Split(lines[0], " ")
	seeds := make([]int64, 0, len(seedStrings)-1)
	for _, seed := range seedStrings[1:] {
		s, _ := strconv.ParseInt(seed, 10, 64)
		seeds = append(seeds, s)
	}

	// Parse the mappers.
	mappers := make([]mapper, 0, 7)
	for _, mapping := range lines[1:] {
		mappers = append(mappers, mapper{})
		splitted := strings.Split(mapping, "\n")[1:]
		for _, line := range splitted {
			parts := regexp.MustCompile(`\d+`).FindAllString(line, -1)
			dst, _ := strconv.ParseInt(parts[0], 0, 0)
			src, _ := strconv.ParseInt(parts[1], 0, 0)
			mlen, _ := strconv.ParseInt(parts[2], 0, 0)
			mappers[len(mappers)-1].data = append(mappers[len(mappers)-1].data, mapData{src, dst, mlen})			
		}
	}
	return seeds, multiMapper{mappers}
}

func part1(seeds []int64, mappers multiMapper) any {
	vals := make([]int64, len(seeds))
	copy(vals, seeds)

	for _, mapping := range mappers.mappers {
		new_vals := make([]int64, 0, len(seeds))
		for _, val := range vals {
			newval := val 
			for _, stuff := range mapping.data {
				if (stuff.src <= val  && val < stuff.src + stuff.len) {
					newval = val - stuff.src + stuff.dst
					break
				}
			}
			new_vals = append(new_vals, newval)
		}
		vals = new_vals
	}

	min := vals[0]
	for _, v := range vals {
        if (v < min) {
            min = v
        }
	}
	return min
}

type span_range struct {
	span_start int64
	span_end int64
}

type offsetmap struct {
	offset int64 
	map_start int64 
	map_end int64
}

func part2(seeds []int64, mappers multiMapper) any {

	ranges := make([]span_range, 0, len(seeds) / 2)
	for i := 0; i < len(seeds) ; i += 2 {
		ranges = append(ranges, span_range{seeds[i], seeds[i]+seeds[i + 1] - 1})
	}

	offsets := make([][]offsetmap, 0, 7)
	for _, mapper := range mappers.mappers {
		offsetmaps := make([]offsetmap, 0, len(mappers.mappers))
		for _, data := range mapper.data {
			offsetmaps = append(offsetmaps, offsetmap{data.dst - data.src, data.src, data.src+data.len-1})
		}
		offsets = append(offsets, offsetmaps)
	} 

	for _, offset := range offsets {
		new_ranges := make([]span_range, 0)
		for _, seed_span := range ranges {
			unprocessed := []span_range {seed_span}
			for _, offsetmaps := range offset {
				new_unprocessed := make([]span_range, 0)
				for _, item := range unprocessed {
					toMap, extra := carve(item.span_start, item.span_end, offsetmaps.map_start, offsetmaps.map_end)
					new_unprocessed = append(new_unprocessed, extra...)
					for _, stuff := range toMap {
						new_ranges = append(new_ranges, span_range{stuff.span_start + offsetmaps.offset, stuff.span_end + offsetmaps.offset})
					}
				}
				unprocessed = new_unprocessed
			}
			new_ranges = append(new_ranges, unprocessed...)
		}
		ranges = new_ranges
	}

	min := ranges[0].span_start
	for _, v := range ranges {
        if (v.span_start < min) {
            min = v.span_start
        }
	}
	return min
}

//  figure out which bits overlapped in the mapping, and which didn't
func carve(span_start int64, span_end int64, map_start int64, map_end int64) ([]span_range, []span_range) {
	a,b := span_start, span_end 
    c,d := map_start, map_end  
    if a < c && d < b {
        return []span_range {{c, d}}, []span_range { {a, c-1}, {d+1, b}}
	} else if a < c && c <= b {
		return []span_range {{c, b}}, []span_range { {a, c-1}}
	} else if a <= d && d < b {
        return []span_range {{a, d}}, []span_range { {d+1, b}}
	} else if c <= a && b <= d {
		return []span_range {{a, b}}, []span_range {}
	} else {
        return []span_range {}, []span_range {{a, b}}

	}
}