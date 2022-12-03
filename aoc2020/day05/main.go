package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	list := getList("input.txt")
	highest := getHighest(list)
	fmt.Println("highest: ", highest)

	missing := findMissing(list)
	fmt.Println("missing: ", missing)
}

func getList(filename string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}

func parseBoardingNumber(str string) int {
	str = strings.ReplaceAll(str, "F", "0")
	str = strings.ReplaceAll(str, "B", "1")
	str = strings.ReplaceAll(str, "L", "0")
	str = strings.ReplaceAll(str, "R", "1")

	val, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(val)
}

func getHighest(list []string) int {
	highest := 0
	for _, line := range list {
		num := parseBoardingNumber(line)
		if num > highest {
			highest = num
		}
	}

	return highest
}

func findMissing(list []string) int {
	highest := getHighest(list)
	seatMap := make([]bool, highest+1)
	for _, line := range list {
		num := parseBoardingNumber(line)
		seatMap[num] = true
	}

	for i := 1; i < highest-1; i++ {
		if !seatMap[i] && seatMap[i-1] && seatMap[i+1] {
			return i
		}
	}
	return -1
}
