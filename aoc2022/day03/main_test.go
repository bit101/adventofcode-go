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

func TestParse(t *testing.T) {
	lines := getLines("test.txt")
	for _, line := range lines {
		rs := parseRucksack(line)
		size1 := len(rs.c1)
		size2 := len(rs.c1)

		if size1 != size2 {
			t.Errorf("expected %d and %d to be equal.", size1, size2)
		}
	}
}

func TestDiff(t *testing.T) {
	lines := getLines("test.txt")
	exp := []string{"p", "L", "P", "v", "t", "s", "v", "l", "V"}
	for i, line := range lines {
		rs := parseRucksack(line)
		diff := findDiff(rs.c1, rs.c2)
		if diff != exp[i] {
			t.Errorf("expected %q, got %q.\n", exp[i], diff)
		}
	}
}

func TestScore(t *testing.T) {
	lines := getLines("test.txt")
	exp := []int{16, 38, 42, 22, 20, 19, 22, 12, 48}
	for i, line := range lines {
		rs := parseRucksack(line)
		diff := findDiff(rs.c1, rs.c2)
		score := getScore(diff)
		if score != exp[i] {
			t.Errorf("expected %d, got %d.\n", exp[i], score)
		}
	}
}

func TestGetGroups(t *testing.T) {
	lines := getLines("test.txt")
	groups := getGroups(lines)
	size := len(groups)
	exp := 3
	if size != exp {
		t.Errorf("expected %d, got %d.\n", exp, size)
	}
}

func TestBadge(t *testing.T) {
	lines := getLines("test.txt")
	exp := []string{"r", "Z", "M"}
	groups := getGroups(lines)
	for i, g := range groups {
		badge := getBadge(g)
		if badge != exp[i] {
			t.Errorf("expected %q, got %q.\n", exp[i], badge)
		}
	}
}

func TestBadgeScores(t *testing.T) {
	lines := getLines("test.txt")
	exp := []int{18, 52, 39}
	groups := getGroups(lines)
	for i, g := range groups {
		badge := getBadge(g)
		score := getScore(badge)
		if score != exp[i] {
			t.Errorf("expected %d, got %d.\n", exp[i], score)
		}
	}
}
