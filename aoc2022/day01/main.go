package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := getLines("input.txt")
	elves := getElves(lines)
	top := getHighest(elves, 1)
	top3 := getHighest(elves, 3)

	fmt.Println("top: ", top)
	fmt.Println("top 3: ", top3)
}

func getLines(filename string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")
	if err != nil {
		log.Fatal(err)
	}
	return lines
}

func getElves(lines []string) []int {
	elves := []int{}

	total := 0
	for _, line := range lines {
		if line == "" {
			if total > 0 {
				elves = append(elves, total)
			}
			total = 0
		} else {
			num, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			total += num
		}
	}
	return elves
}

func getHighest(elves []int, count int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	total := 0
	for i := 0; i < count; i++ {
		total += elves[i]
	}
	return total
}
