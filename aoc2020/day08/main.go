// Package main is the executable
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := getLines("input.txt")
	// lines := getLines("test.txt")
	ins := getInstructions(lines)
	acc := findLoop(ins)
	fmt.Println("acc: ", acc)

	acc = findEnd(ins)
	fmt.Println("acc end: ", acc)
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

type instruction struct {
	name  string
	value int
}

func newInstruction(line string) instruction {
	parts := strings.Split(line, " ")
	name := strings.Trim(parts[0], " ")
	value, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	return instruction{name, value}
}

func getInstructions(lines []string) []instruction {
	instructions := []instruction{}
	for _, line := range lines {
		ins := newInstruction(line)
		instructions = append(instructions, ins)
	}
	return instructions
}

func findLoop(instructions []instruction) int {
	acc := 0
	visited := make([]bool, len(instructions))

	index := 0
	visited[index] = true
	for true {
		switch instructions[index].name {
		case "nop":
			index++
		case "jmp":
			index += instructions[index].value
		case "acc":
			acc += instructions[index].value
			index++
		}
		if visited[index] {
			break
		}
		visited[index] = true
	}

	return acc
}

func findLoopOrEnd(instructions []instruction) (bool, int) {
	acc := 0
	visited := make([]bool, len(instructions))

	finished := false
	index := 0
	visited[index] = true
	for true {
		switch instructions[index].name {
		case "nop":
			index++
		case "jmp":
			index += instructions[index].value
		case "acc":
			acc += instructions[index].value
			index++
		}
		if index == len(instructions) {
			finished = true
			break
		}
		if visited[index] {
			break
		}
		visited[index] = true
	}

	return finished, acc
}

func findEnd(ins []instruction) int {
	found := false
	acc := 0
	changeIndex := -1
	var err error

	for !found {
		found, acc = findLoopOrEnd(ins)
		if !found {
			// revert last change
			changeInstruction(changeIndex, ins)
			changeIndex, err = findNextChange(changeIndex, ins)
			if err != nil {
				log.Fatal(err)
			}
			changeInstruction(changeIndex, ins)
		}
	}

	return acc
}

func changeInstruction(index int, ins []instruction) {
	if index == -1 {
		// there was no last change
		return
	}
	if ins[index].name == "nop" {
		ins[index].name = "jmp"
	} else if ins[index].name == "jmp" {
		ins[index].name = "nop"
	}
}

func findNextChange(index int, ins []instruction) (int, error) {
	for i := index + 1; i < len(ins); i++ {
		if ins[i].name == "nop" || ins[i].name == "jmp" {
			return i, nil
		}
	}
	return -1, errors.New("couldn't find next change")
}
