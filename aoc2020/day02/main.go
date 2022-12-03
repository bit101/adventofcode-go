package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	passwords := getPasswords("input.txt")
	count := 0
	for _, p := range passwords {
		if validatePassword(p) {
			count++
		}
	}
	fmt.Println("count 1: ", count)

	count = 0
	for _, p := range passwords {
		if validatePassword2(p) {
			count++
		}
	}
	fmt.Println("count 1: ", count)
}

type password struct {
	min  int
	max  int
	char byte
	pw   string
}

func getPasswords(filename string) []password {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	passwords := []password{}
	for _, line := range lines {
		passwords = append(passwords, parsePassword(line))
	}
	return passwords
}

func parsePassword(line string) password {
	var min, max int
	var char byte
	var pw string

	fmt.Sscanf(line, "%d-%d %c: %s", &min, &max, &char, &pw)
	return password{min, max, char, pw}
}

func validatePassword(pw password) bool {

	count := 0
	for i := 0; i < len(pw.pw); i++ {
		c := pw.pw[i]
		if c == pw.char {
			count++
		}
	}
	return count >= pw.min && count <= pw.max
}

func validatePassword2(pw password) bool {

	x := pw.pw[pw.min-1] == pw.char
	y := pw.pw[pw.max-1] == pw.char
	if x != y {
		return true
	}
	return false
}
