// Package main is the executable
package main

import (
	"testing"
)

func TestGetLines(t *testing.T) {
	lines := getLines("test.txt")
	count := len(lines)
	exp := 9

	if count != exp {
		t.Errorf("expected %d, got %d\n", exp, count)
	}
}

func TestParseBags(t *testing.T) {
	exp := []string{
		"light red",
		"dark orange",
		"bright white",
		"muted yellow",
		"shiny gold",
		"dark olive",
		"vibrant plum",
		"faded blue",
		"dotted black",
	}
	exp2 := []int{2, 2, 1, 2, 2, 2, 2, 0, 0}
	exp3 := []int{3, 7, 1, 11, 3, 7, 11, 0, 0}
	lines := getLines("test.txt")
	bags := parseBags(lines)
	for i, b := range bags {
		if b.container != exp[i] {
			t.Errorf("expected %q, got %q\n", exp[i], b.container)
		}
		if len(b.contents) != exp2[i] {
			t.Errorf("expected %q, got %q\n", exp2[i], b.contents)
		}
		totalCount := 0
		for _, bi := range b.contents {
			totalCount += bi.count
		}
		if totalCount != exp3[i] {
			t.Errorf("expected %q, got %q\n", exp2[i], b.contents)
		}
	}
}

func TestGetBagFor(t *testing.T) {
	lines := getLines("test.txt")
	bags := parseBags(lines)
	matching := getBagsFor(bags, "shiny gold")
	size := len(matching)
	exp := 2
	if size != exp {
		t.Errorf("expected %d, got %d\n", exp, size)
	}
}

func TestGetAllBagsFor(t *testing.T) {
	lines := getLines("test.txt")
	bags := parseBags(lines)
	matching := getAllBagsFor(bags, "shiny gold")
	size := len(matching)
	exp := 4
	if size != exp {
		t.Errorf("expected %d, got %d\n", exp, size)
	}

	exp2 := []string{
		"bright white",
		"muted yellow",
		"light red",
		"dark orange",
	}
	for i, b := range matching {
		if b.container != exp2[i] {
			t.Errorf("expected %q, got %q\n", exp2[i], b.contents)
		}
	}
}

func TestGetBagCount(t *testing.T) {
	lines := getLines("test.txt")
	bags := parseBags(lines)

	count := getBagCount(bags, "shiny gold")
	exp := 32
	if count != exp {
		t.Errorf("expected %d, got %d\n", exp, count)
	}

}
