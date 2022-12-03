package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	lines := getLines("input.txt")
	total := 0
	for _, line := range lines {
		rs := parseRucksack(line)
		diff := findDiff(rs.c1, rs.c2)
		score := getScore(diff)
		total += score
	}
	fmt.Println("score: ", total)

	total = 0
	groups := getGroups(lines)
	for _, g := range groups {
		badge := getBadge(g)
		score := getScore(badge)
		total += score
	}
	fmt.Println("badge score: ", total)
	// 2853 too high
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
	length := len(lines)
	if lines[length-1] == "" {
		lines = lines[:length-1]
	}
	return lines
}

type rucksack struct {
	c1 string
	c2 string
}

func parseRucksack(line string) rucksack {
	size := len(line)
	rs := rucksack{
		line[0 : size/2],
		line[size/2:],
	}
	return rs
}

func findDiff(c1, c2 string) string {
	for _, a := range c1 {
		for _, b := range c2 {
			if a == b {
				return string(a)
			}
		}
	}
	return ""
}

func getScore(s string) int {
	c := s[0]
	if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 1)
	}
	if c >= 'A' && c <= 'Z' {
		return int(c - 'A' + 27)
	}
	return 0

}

type group struct {
	rs0 string
	rs1 string
	rs2 string
}

func getGroups(lines []string) []group {
	groups := []group{}
	for i := 0; i < len(lines)-2; i += 3 {
		rs0 := lines[i+0]
		rs1 := lines[i+1]
		rs2 := lines[i+2]
		g := group{rs0, rs1, rs2}
		groups = append(groups, g)
	}
	return groups
}

func getBadge(g group) string {
	common := []rune{}
	for _, a := range g.rs0 {
		for _, b := range g.rs1 {
			if a == b {
				common = append(common, a)
			}
		}
	}
	for _, a := range common {
		for _, b := range g.rs2 {
			if a == b {
				return string(a)
			}
		}
	}
	return ""
}
