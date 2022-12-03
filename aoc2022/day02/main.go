package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	lines := getLines("input.txt")
	score := 0
	for _, line := range lines {
		turn := getPlays(line)
		turnScore := getScore(turn)
		score += turnScore
	}
	fmt.Println("score: ", score)

	score = 0
	for _, line := range lines {
		strat := getStrategy(line)
		score += getScore2(strat)
	}
	fmt.Println("score: ", score)

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

type move int

const (
	rock move = iota
	paper
	scissors
)

type play struct {
	one move
	two move
}

func getPlays(line string) play {
	var one move
	var two move
	switch string(line[0]) {
	case "A":
		one = rock
	case "B":
		one = paper
	case "C":
		one = scissors
	}

	switch string(line[2]) {
	case "X":
		two = rock
	case "Y":
		two = paper
	case "Z":
		two = scissors
	}
	return play{one, two}
}

func getScore(turn play) int {
	score := turn.two + 1

	diff := int(turn.two - turn.one)
	if diff == 1 || diff == -2 {
		score += 6
	} else if diff == 0 {
		score += 3
	}
	return int(score)
}

type outcome int

const (
	lose outcome = iota
	draw
	win
)

type strategy struct {
	one    move
	result outcome
}

func getStrategy(line string) strategy {
	var one move
	var result outcome
	switch string(line[0]) {
	case "A":
		one = rock
	case "B":
		one = paper
	case "C":
		one = scissors
	}

	switch string(line[2]) {
	case "X":
		result = lose
	case "Y":
		result = draw
	case "Z":
		result = win
	}
	return strategy{one, result}
}

func getMoveToMake(turn strategy) move {
	switch turn.one {
	case rock:
		switch turn.result {
		case win:
			return paper
		case lose:
			return scissors
		}
	case paper:
		switch turn.result {
		case win:
			return scissors
		case lose:
			return rock
		}
	case scissors:
		switch turn.result {
		case win:
			return rock
		case lose:
			return paper
		}
	}
	return turn.one
}

func getScore2(strat strategy) int {
	next := getMoveToMake(strat)

	score := next + 1
	switch strat.result {
	case win:
		score += 6
	case draw:
		score += 3
	}

	return int(score)
}
