package main

import (
	"aoc2023/util"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type AlminacMap []AlminacTranslation

type AlminacTranslation struct {
	destStart int
	srcStart  int
	rng       int
}

func (am AlminacMap) mapInput(id int) int {
	for _, at := range am {
		if at.hasTranslation(id) {
			return at.translate(id)
		}
	}
	return id
}

func (at AlminacTranslation) hasTranslation(id int) bool {
	return id >= at.srcStart && id < at.srcStart+at.rng
}

func (at AlminacTranslation) translate(id int) int {
	return id + (at.destStart - at.srcStart)
}

func main() {
	input := util.ReadDay(5)

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	seeds, maps := parseInput(input)
	min := -1

	// fmt.Printf("%+v %+v\n", seeds, maps)
	for _, seed := range seeds {
		id := seed
		for _, am := range maps {
			id = am.mapInput(id)
		}
		if min == -1 || id < min {
			min = id
		}
	}

	return min
}

func part2(input []string) int {
	seeds, maps := parseInput(input)
	var mutex *sync.Mutex = new(sync.Mutex)
	var wg *sync.WaitGroup = new(sync.WaitGroup)

	min := -1

	for i := 0; i < len(seeds); i += 2 {
		wg.Add(1)
		go func(seedStart int, seedRange int, maps []AlminacMap, mutex *sync.Mutex, wg *sync.WaitGroup) {
			defer wg.Done()
			localMin := -1
			// fmt.Printf("%d - %d\n", seedStart, seedStart+seedRange-1)
			for j := 0; j < seedRange; j++ {
				id := seedStart + j
				for _, am := range maps {
					id = am.mapInput(id)
				}
				if localMin == -1 || id < localMin {
					localMin = id
				}
			}
			mutex.Lock()
			if min == -1 || localMin < min {
				min = localMin
			}
			mutex.Unlock()
		}(seeds[i], seeds[i+1], maps, mutex, wg)
	}

	wg.Wait()

	return min
}

func parseInput(input []string) (seeds []int, maps []AlminacMap) {
	after, _ := strings.CutPrefix(input[0], "seeds: ")
	seeds = parseNumbers(after)

	currMap := AlminacMap{}
	for i := 3; i < len(input); i++ {
		if input[i] == "" {
			i++
			maps = append(maps, currMap)
			currMap = AlminacMap{}
			continue
		}
		nums := parseNumbers(input[i])
		currMap = append(currMap, AlminacTranslation{nums[0], nums[1], nums[2]})
	}
	maps = append(maps, currMap)

	return
}

func parseNumbers(list string) []int {
	result := []int{}

	for _, n := range strings.Split(list, " ") {
		if n == "" {
			continue
		}
		v, _ := strconv.Atoi(n)
		result = append(result, v)
	}

	return result
}
