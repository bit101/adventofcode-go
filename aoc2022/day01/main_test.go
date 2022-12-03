package main

import (
	"testing"
)

func TestGetLines(t *testing.T) {
	lines := getLines("test.txt")
	count := len(lines)

	if count != 15 {
		t.Errorf("expected %d, got %d\n", 14, count)
	}
}

func TestGetElves(t *testing.T) {
	lines := getLines("test.txt")
	elves := getElves(lines)

	count := len(elves)

	if count != 5 {
		t.Errorf("expected %d, got %d\n", 5, count)
	}

	e0 := elves[0]
	if e0 != 6000 {
		t.Errorf("expected %d, got %d\n", 6000, e0)
	}

	e1 := elves[1]
	if e1 != 4000 {
		t.Errorf("expected %d, got %d\n", 4000, e1)
	}

	e2 := elves[2]
	if e2 != 11000 {
		t.Errorf("expected %d, got %d\n", 11000, e2)
	}

	e3 := elves[3]
	if e3 != 24000 {
		t.Errorf("expected %d, got %d\n", 24000, e3)
	}

	e4 := elves[4]
	if e4 != 10000 {
		t.Errorf("expected %d, got %d\n", 10000, e4)
	}
}

func TestHighest(t *testing.T) {
	lines := getLines("test.txt")
	elves := getElves(lines)

	top := getHighest(elves, 1)
	if top != 24000 {
		t.Errorf("expected %d, got %d\n", 24000, top)
	}

	top = getHighest(elves, 3)
	if top != 45000 {
		t.Errorf("expected %d, got %d\n", 45000, top)
	}

}
