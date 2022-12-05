// Package main is the executable
package main

import (
	"testing"
)

func TestGetLines(t *testing.T) {
	lines := getData("test.txt")
	count := len(lines)
	exp := 6

	if count != exp {
		t.Errorf("expected %d, got %d\n", exp, count)
	}
}

func TestParsePair(t *testing.T) {
	exp := [][]int{
		{2, 4, 6, 8},
		{2, 3, 4, 5},
		{5, 7, 7, 9},
		{2, 8, 3, 7},
		{6, 6, 4, 6},
		{2, 6, 4, 8},
	}

	lines := getData("test.txt")
	for i, line := range lines {
		aPair := parsePair(line)
		first := aPair.first
		second := aPair.second
		if first.start != exp[i][0] {
			t.Errorf("expected %d, got %d\n", exp[i][0], first.start)
		}
		if first.end != exp[i][1] {
			t.Errorf("expected %d, got %d\n", exp[i][1], first.end)
		}
		if second.start != exp[i][2] {
			t.Errorf("expected %d, got %d\n", exp[i][2], second.start)
		}
		if second.end != exp[i][3] {
			t.Errorf("expected %d, got %d\n", exp[i][3], second.end)
		}
	}
}

func TestEnclosed(t *testing.T) {
	exp := []bool{false, false, false, true, true, false}

	lines := getData("test.txt")
	for i, line := range lines {
		aPair := parsePair(line)
		enclosed := isEnclosed(aPair)
		if enclosed != exp[i] {
			t.Errorf("expected %t, got %t\n", exp[i], enclosed)
		}
	}
}

func TestOverlap(t *testing.T) {

	exp := []bool{false, false, true, true, true, true}

	lines := getData("test.txt")
	for i, line := range lines {
		aPair := parsePair(line)
		enclosed := doOverlap(aPair)
		if enclosed != exp[i] {
			t.Errorf("expected %t, got %t\n", exp[i], enclosed)
		}
	}
}
