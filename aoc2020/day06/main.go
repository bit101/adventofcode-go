// Package main is the executable
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	total := 0
	lines := getLines("input.txt")
	groups := getGroups(lines)
	for _, g := range groups {
		a := getScore(g)
		total += a
	}
	fmt.Println("total: ", total)

	total = 0
	for _, g := range groups {
		c := getAllAnswersCount(g)
		total += c
	}
	fmt.Println("total all answered: ", total)
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

type group struct {
	lines []string
}

func newGroup() group {
	return group{
		[]string{},
	}
}

func getGroups(lines []string) []group {
	groups := []group{}
	g := newGroup()
	for _, line := range lines {
		if line == "" {
			groups = append(groups, g)
			g = newGroup()
		} else {
			g.lines = append(g.lines, line)
		}
	}
	groups = append(groups, g)
	return groups
}

func getAnswers(g group) []string {
	answers := make(map[string]bool)
	for _, line := range g.lines {
		for _, c := range line {
			answers[string(c)] = true
		}
	}
	result := []string{}
	for k := range answers {
		result = append(result, k)
	}

	sort.Strings(result)
	return result
}

func getScore(g group) int {
	answers := getAnswers(g)
	return len(answers)
}

func getAllAnswersCount(g group) int {
	count := 0
	for _, c := range g.lines[0] {
		if checkOthersInGroup(c, g) {
			count++
		}
	}
	return count
}

func checkOthersInGroup(c rune, g group) bool {
	for i := 1; i < len(g.lines); i++ {
		line := g.lines[i]
		found := false
		for _, d := range line {
			if c == d {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}
