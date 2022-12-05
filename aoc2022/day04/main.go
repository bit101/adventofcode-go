// Package main is the executable
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lines := getData("input.txt")
	// lines := getLines("test.txt")
	count := 0
	for _, line := range lines {
		aPair := parsePair(line)
		enclosed := isEnclosed(aPair)
		if enclosed {
			count++
		}
	}
	fmt.Println("enclosed: ", count)

	count = 0
	for _, line := range lines {
		aPair := parsePair(line)
		overlap := doOverlap(aPair)
		if overlap {
			count++
		}
	}
	fmt.Println("overlap: ", count)
}

func getData(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

type area struct {
	start, end int
}

type pair struct {
	first, second area
}

func parsePair(line string) pair {
	var firstStart, firstEnd, secondStart, secondEnd int
	fmt.Sscanf(line, "%d-%d,%d-%d", &firstStart, &firstEnd, &secondStart, &secondEnd)
	return pair{
		first:  area{firstStart, firstEnd},
		second: area{secondStart, secondEnd},
	}
}

func isEnclosed(aPair pair) bool {
	if aPair.first.start <= aPair.second.start && aPair.first.end >= aPair.second.end {
		return true
	}
	if aPair.second.start <= aPair.first.start && aPair.second.end >= aPair.first.end {
		return true
	}
	return false
}

func doOverlap(aPair pair) bool {
	if aPair.first.start <= aPair.second.start && aPair.first.end >= aPair.second.start {
		return true
	}
	if aPair.first.start <= aPair.second.end && aPair.first.end >= aPair.second.end {
		return true
	}
	return isEnclosed(aPair)
}
