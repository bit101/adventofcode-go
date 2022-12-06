// Package main is the executable
package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	lines := getData("input.txt")
	// lines := getData("test.txt")
	marker, err := findMarker(lines[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("marker: ", marker)

	message, err := findMessage(lines[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("message: ", message)
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

func findMarker(line string) (int, error) {
	for i := 0; i < len(line)-4; i++ {
		seq := line[i : i+4]
		if unique(seq) {
			return i + 4, nil
		}
	}
	return 0, errors.New("no marker found")
}

func unique(seq string) bool {
	for i := 0; i < len(seq)-1; i++ {
		a := seq[i]
		for j := i + 1; j < len(seq); j++ {
			b := seq[j]
			if a == b {
				return false
			}
		}
	}
	return true
}

func findMessage(line string) (int, error) {
	for i := 0; i < len(line)-14; i++ {
		seq := line[i : i+14]
		if unique(seq) {
			return i + 14, nil
		}
	}
	return 0, errors.New("no message found")
}
