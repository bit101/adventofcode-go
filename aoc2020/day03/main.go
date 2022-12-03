package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	mymap := getMap("input.txt")
	a := traverse(0, 0, 1, 1, mymap)
	b := traverse(0, 0, 3, 1, mymap)
	c := traverse(0, 0, 5, 1, mymap)
	d := traverse(0, 0, 7, 1, mymap)
	e := traverse(0, 0, 1, 2, mymap)
	fmt.Println("count: ", a*b*c*d*e)
}

func getMap(filename string) []string {
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

func hasTree(m []string, x, y int) bool {
	length := len(m[0])
	x %= length
	return m[y][x] == '#'
}

func traverse(x, y, dx, dy int, mymap []string) int {
	count := 0
	rows := len(mymap)
	for ; y < rows; y += dy {
		if hasTree(mymap, x, y) {
			count++
		}
		x += dx
	}
	return count
}
