package main

import (
	"aoc2023/util"
	"fmt"
	"hash/fnv"
)

func main() {
	input := util.ReadDay(14)

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	result := 0

	for col := range input[0] {
		slideTo := 0
		for row := range input {
			switch input[row][col] {
			case 'O':
				result += len(input) - slideTo
				slideTo++
			case '#':
				slideTo = row + 1
			}
		}
	}

	return result
}

type hashLoad struct {
	hash uint32
	load int
}

const stabilizeCount = 200

func part2(input []string) int {
	rocks := make([][]byte, 0, len(input))
	for _, line := range input {
		rocks = append(rocks, []byte(line))
	}

	for i := 0; i < stabilizeCount; i++ {
		rotation(rocks)
	}

	repeated := []hashLoad{}
	for true {
		rotation(rocks)
		next := hashLoad{hash(rocks), totalLoad(rocks)}
		if len(repeated) > 0 && repeated[0].hash == next.hash {
			break
		}
		repeated = append(repeated, next)
	}

	// why "- 1"?
	i := (1000000000 - stabilizeCount - 1) % len(repeated)
	return repeated[i].load
}

func rotation(rocks [][]byte) {
	slideNorth(rocks)
	slideWest(rocks)
	slideSouth(rocks)
	slideEast(rocks)
}

func slideNorth(rocks [][]byte) {
	for col := range rocks[0] {
		slideTo := 0
		for row := range rocks {
			switch rocks[row][col] {
			case 'O':
				if slideTo != row {
					rocks[slideTo][col] = 'O'
					rocks[row][col] = '.'
				}
				slideTo++
			case '#':
				slideTo = row + 1
			}
		}
	}
}

func slideSouth(rocks [][]byte) {
	for col := range rocks[0] {
		slideTo := len(rocks) - 1
		for row := len(rocks) - 1; row >= 0; row-- {
			switch rocks[row][col] {
			case 'O':
				if slideTo != row {
					rocks[slideTo][col] = 'O'
					rocks[row][col] = '.'
				}
				slideTo--
			case '#':
				slideTo = row - 1
			}
		}
	}
}

func slideWest(rocks [][]byte) {
	for row := range rocks {
		slideTo := 0
		for col := range rocks[0] {
			switch rocks[row][col] {
			case 'O':
				if slideTo != col {
					rocks[row][slideTo] = 'O'
					rocks[row][col] = '.'
				}
				slideTo++
			case '#':
				slideTo = col + 1
			}
		}
	}
}

func slideEast(rocks [][]byte) {
	for row := range rocks {
		slideTo := len(rocks[0]) - 1
		for col := len(rocks[0]) - 1; col >= 0; col-- {
			switch rocks[row][col] {
			case 'O':
				if slideTo != col {
					rocks[row][slideTo] = 'O'
					rocks[row][col] = '.'
				}
				slideTo--
			case '#':
				slideTo = col - 1
			}
		}
	}
}

func totalLoad(rocks [][]byte) int {
	result := 0
	for row := range rocks {
		for col := range rocks[row] {
			if rocks[row][col] == 'O' {
				result += len(rocks) - row
			}
		}
	}
	return result
}

func hash(rocks [][]byte) uint32 {
	h := fnv.New32a()
	for _, r := range rocks {
		h.Write(r)
	}
	return h.Sum32()
}
