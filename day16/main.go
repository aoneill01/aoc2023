package main

import (
	"aoc2023/util"
	"fmt"
)

const (
	north = 1
	east  = 2
	south = 4
	west  = 8
)

type location struct {
	row int
	col int
}

func (l location) move(dir byte) location {
	switch dir {
	case north:
		return location{l.row - 1, l.col}
	case east:
		return location{l.row, l.col + 1}
	case south:
		return location{l.row + 1, l.col}
	case west:
		return location{l.row, l.col - 1}
	default:
		return l
	}
}

func main() {
	input := util.ReadDay(16)

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	seen := make(map[location]byte)
	processLight(location{0, -1}, east, input, seen)
	return len(seen)
}

func part2(input []string) int {
	max := -1
	var seen map[location]byte
	for row := range input {
		seen = make(map[location]byte)
		processLight(location{row, -1}, east, input, seen)
		if len(seen) > max {
			max = len(seen)
		}
		seen = make(map[location]byte)
		processLight(location{row, len(input[0])}, west, input, seen)
		if len(seen) > max {
			max = len(seen)
		}
	}
	for col := range input[0] {
		seen = make(map[location]byte)
		processLight(location{-1, col}, south, input, seen)
		if len(seen) > max {
			max = len(seen)
		}
		seen = make(map[location]byte)
		processLight(location{len(input), col}, north, input, seen)
		if len(seen) > max {
			max = len(seen)
		}
	}
	return max
}

func processLight(l location, dir byte, tiles []string, seen map[location]byte) {
	next := l.move(dir)
	// fmt.Println(next)
	// Off the edge
	if next.row < 0 || next.col < 0 || next.row >= len(tiles) || next.col >= len(tiles[0]) {
		return
	}
	// Already been to this location moving in this direction
	if seen[next]&dir != 0 {
		return
	}
	seen[next] |= dir

	switch tiles[next.row][next.col] {
	case '.':
		processLight(next, dir, tiles, seen)
	case '|':
		if dir == north || dir == south {
			processLight(next, dir, tiles, seen)
		} else {
			processLight(next, north, tiles, seen)
			processLight(next, south, tiles, seen)
		}
	case '-':
		if dir == north || dir == south {
			processLight(next, east, tiles, seen)
			processLight(next, west, tiles, seen)
		} else {
			processLight(next, dir, tiles, seen)
		}
	case '/':
		switch dir {
		case north:
			processLight(next, east, tiles, seen)
		case east:
			processLight(next, north, tiles, seen)
		case south:
			processLight(next, west, tiles, seen)
		case west:
			processLight(next, south, tiles, seen)
		}
	case '\\':
		switch dir {
		case north:
			processLight(next, west, tiles, seen)
		case east:
			processLight(next, south, tiles, seen)
		case south:
			processLight(next, east, tiles, seen)
		case west:
			processLight(next, north, tiles, seen)
		}
	}
}
