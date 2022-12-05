// Package main is the executable
package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := getData("input.txt")
	// lines := getData("test.txt")
	stacks := getStacks(lines)
	moves := getMoves(lines)
	for _, m := range moves {
		makeMove(stacks, m)
	}
	top := getTop(stacks)
	fmt.Println("top: ", top)

	stacks = getStacks(lines)
	moves = getMoves(lines)
	for _, m := range moves {
		makeMove2(stacks, m)
	}
	top = getTop(stacks)
	fmt.Println("top 2: ", top)
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

type stack struct {
	list []string
	size int
}

func newStack() *stack {
	return &stack{
		[]string{},
		0,
	}
}

func (s *stack) push(value string) {
	s.list = append(s.list, value)
	s.size++
}

func (s *stack) pop() (string, error) {
	if s.size < 1 {
		return "", errors.New("nothing to pop")
	}
	value := s.list[s.size-1]
	s.size--
	s.list = s.list[:s.size]
	return value, nil
}

func (s *stack) peek() (string, error) {
	if s.size < 1 {
		return "", errors.New("nothing to pop")
	}
	value := s.list[s.size-1]
	return value, nil
}

func (s *stack) String() string {
	return strings.Join(s.list, "+")
}

func getStacks(data []string) []*stack {
	stackLines := []string{}
	for _, line := range data {
		if line == "" {
			break
		}
		stackLines = append(stackLines, line)
	}
	last := strings.Trim(stackLines[len(stackLines)-1], " ")
	indices := strings.Split(last, " ")
	lastIndex := indices[len(indices)-1]
	size, err := strconv.Atoi(lastIndex)
	if err != nil {
		log.Fatal(err)
	}
	stacks := make([]*stack, size)
	for i := 0; i < size; i++ {
		stacks[i] = newStack()
	}

	// parse lines in reverse order
	for i := len(stackLines) - 2; i >= 0; i-- {
		line := stackLines[i]
		// get char from each column (might be empty or past end of line)
		for j := 0; j < size; j++ {
			// position of char
			k := j*4 + 1
			// is is beyond end of line?
			if k >= len(line) {
				break
			}
			value := string(line[k])
			// is it empty
			if value != " " {
				// found valid char. push it onto current column's stack.
				stacks[j].push(value)
			}
		}
	}
	return stacks
}

type move struct {
	from, to, count int
}

func getMoves(data []string) []move {
	var index int
	for index = 0; index < len(data); index++ {
		line := strings.Trim(data[index], " ")
		if line == "" {
			break
		}
	}
	index++

	moves := []move{}
	for i := index; i < len(data); i++ {
		var from, to, count int
		fmt.Sscanf(data[i], "move %d from %d to %d", &count, &from, &to)
		moves = append(moves, move{from - 1, to - 1, count})
	}
	return moves
}

func makeMove(stacks []*stack, aMove move) {
	for i := 0; i < aMove.count; i++ {
		val, err := stacks[aMove.from].pop()
		if err != nil {
			log.Fatal(err)
		}
		stacks[aMove.to].push(val)
	}
}

func getTop(stacks []*stack) string {
	top := ""
	for _, s := range stacks {
		val, err := s.peek()
		if err != nil {
			log.Fatal(err)
		}
		top += val
	}
	return top
}

func makeMove2(stacks []*stack, aMove move) {
	temp := newStack()
	for i := 0; i < aMove.count; i++ {
		val, err := stacks[aMove.from].pop()
		if err != nil {
			log.Fatal(err)
		}
		temp.push(val)
	}
	for temp.size > 0 {
		val, err := temp.pop()
		if err != nil {
			log.Fatal(err)
		}
		stacks[aMove.to].push(val)
	}
}
