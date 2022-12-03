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
	expenses := getExpenses("input.txt")
	a, b, err := get2020(expenses)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("product of 2 : %d\n", a*b)

	a, b, c, err := get2020x3(expenses)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("product of 3: %d\n", a*b*c)
}

func getExpenses(filename string) []int {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(file), "\n")
	expenses := []int{}
	for _, line := range lines {
		expense, err := strconv.Atoi(line)
		if err == nil {
			expenses = append(expenses, expense)
		}
	}
	return expenses
}

func get2020(expenses []int) (int, int, error) {
	for i := 0; i < len(expenses)-1; i++ {
		a := expenses[i]
		for j := i + 1; j < len(expenses); j++ {
			b := expenses[j]
			if a+b == 2020 {
				return a, b, nil
			}
		}
	}
	return 0, 0, errors.New("no digits add up to 2020")
}

func get2020x3(expenses []int) (int, int, int, error) {
	for i := 0; i < len(expenses)-2; i++ {
		a := expenses[i]
		for j := i + 1; j < len(expenses)-1; j++ {
			b := expenses[j]
			for k := j + 1; k < len(expenses); k++ {
				c := expenses[k]
				if a+b+c == 2020 {
					return a, b, c, nil
				}
			}
		}
	}
	return 0, 0, 0, errors.New("no digits add up to 2020")
}
