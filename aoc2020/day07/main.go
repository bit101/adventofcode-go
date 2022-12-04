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
	bags := parseBags(lines)
	matching := getAllBagsFor(bags, "shiny gold")
	fmt.Println("count: ", len(matching))

	count := getBagCount(bags, "shiny gold")
	fmt.Println("count: ", count)
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

type bag struct {
	container string
	contents  []bagItem
}

type bagItem struct {
	count   int
	bagType string
}

func parseBags(lines []string) []bag {
	bags := []bag{}
	for _, line := range lines {
		parts := strings.Split(line, "contain")
		container := strings.Trim(parts[0], " ")
		contents := strings.Trim(parts[1], " ")
		b := bag{
			container[:len(container)-5],
			parseContents(contents),
		}
		bags = append(bags, b)
	}
	return bags
}

func getBagsFor(bags []bag, bagType string) []bag {
	list := []bag{}
	for _, b := range bags {
		for _, bc := range b.contents {
			if bagType == bc.bagType {
				list = append(list, b)
			}
		}
	}
	return list
}

func getAllBagsFor(bags []bag, bagType string) []bag {
	list := getBagsFor(bags, bagType)
	for _, b := range list {
		sub := getAllBagsFor(bags, b.container)
		list = addToList(list, sub)
	}
	return list
}

func addToList(list, sub []bag) []bag {
	for _, a := range sub {
		found := false
		for _, b := range list {
			if a.container == b.container {
				found = true
				break
			}
		}
		if !found {
			list = append(list, a)
		}
	}
	return list
}

func parseContents(contents string) []bagItem {
	bagItems := []bagItem{}

	if contents == "no other bags." {
		return bagItems
	}
	items := strings.Split(contents, ",")
	for _, item := range items {
		item := strings.Trim(strings.Split(item, " bag")[0], " ")
		parts := strings.Split(item, " ")
		count, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		desc := strings.Join(parts[1:], " ")
		bi := bagItem{count, desc}
		bagItems = append(bagItems, bi)
	}
	return bagItems
}

func getBagCount(bags []bag, bagType string) int {
	count := 0
	b, err := getBagDef(bags, bagType)
	if err != nil {
		log.Fatal(err)
	}
	for _, bc := range b.contents {
		count += bc.count
		count += bc.count * getBagCount(bags, bc.bagType)

	}

	return count
}

func getBagDef(bags []bag, bagType string) (bag, error) {
	for _, b := range bags {
		if b.container == bagType {
			return b, nil
		}
	}
	return bag{}, errors.New("no bag type " + bagType)

}
