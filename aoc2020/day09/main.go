// Package main is the executable
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := getData("input.txt")
	// lines := getData("test.txt")
	nums := getNums(lines)
	num := getFirstError(nums, 25)
	fmt.Println("answer: ", num)

	start, end := findContinguous(nums, num)
	weakness := getWeakness(nums, start, end)
	fmt.Println("weakness: ", weakness)
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

func getNums(list []string) []int {
	ints := []int{}
	for _, line := range list {
		line := strings.Trim(line, " ")
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, num)
	}
	return ints
}

func isValid(num int, list []int, preamble int) bool {
	size := len(list)
	list = list[size-preamble:]
	for i := 0; i < len(list)-1; i++ {
		a := list[i]
		for j := i + 1; j < len(list); j++ {
			b := list[j]
			if a+b == num {
				return true
			}
		}
	}
	return false
}

func getFirstError(nums []int, preamble int) int {
	for i := 0; i < len(nums)-preamble; i++ {
		partial := nums[0 : i+preamble]
		tester := nums[i+preamble]
		if !isValid(tester, partial, preamble) {
			return tester
		}
	}

	return -1
}

func findContinguous(nums []int, sum int) (int, int) {
	start := 0
	end := 0
	for start = 0; start < len(nums); start++ {
		end = 0
		total := 0
		for total <= sum {
			total = contiguousSum(nums, start, end)
			if total == sum {
				return start, end
			}
			end++
		}

	}
	return start, end
}

func contiguousSum(nums []int, start, end int) int {
	sum := 0
	for i := start; i <= end; i++ {
		sum += nums[i]
	}
	return sum
}

func getWeakness(nums []int, start, end int) int {
	nums = nums[start : end+1]
	sort.Ints(nums)
	return nums[0] + nums[len(nums)-1]
}
