// Package main is the executable
package main

import (
	"log"
	"testing"
)

func TestGetLines(t *testing.T) {
	lines := getData("test.txt")
	count := len(lines)
	exp := 5

	if count != exp {
		t.Errorf("expected %d, got %d\n", exp, count)
	}
}

func TestUnique(t *testing.T) {
	x := unique("abcd")
	if x != true {
		t.Errorf("expected %t, got %t\n", true, x)
	}

	x = unique("abca")
	if x != false {
		t.Errorf("expected %t, got %t\n", false, x)
	}

	x = unique("abcc")
	if x != false {
		t.Errorf("expected %t, got %t\n", false, x)
	}

	x = unique("aaxz")
	if x != false {
		t.Errorf("expected %t, got %t\n", false, x)
	}
}

func TestFindMarker(t *testing.T) {
	lines := getData("test.txt")
	marker, err := findMarker(lines[0])
	if err != nil {
		log.Fatal(err)
	}
	exp := 7
	if marker != exp {
		t.Errorf("expected %d, got %d\n", exp, marker)
	}

	marker, err = findMarker(lines[1])
	if err != nil {
		log.Fatal(err)
	}
	exp = 5
	if marker != exp {
		t.Errorf("expected %d, got %d\n", exp, marker)
	}

	marker, err = findMarker(lines[2])
	if err != nil {
		log.Fatal(err)
	}
	exp = 6
	if marker != exp {
		t.Errorf("expected %d, got %d\n", exp, marker)
	}

	marker, err = findMarker(lines[3])
	if err != nil {
		log.Fatal(err)
	}
	exp = 10
	if marker != exp {
		t.Errorf("expected %d, got %d\n", exp, marker)
	}

	marker, err = findMarker(lines[4])
	if err != nil {
		log.Fatal(err)
	}
	exp = 11
	if marker != exp {
		t.Errorf("expected %d, got %d\n", exp, marker)
	}

}

func TestFindMessage(t *testing.T) {
	lines := getData("test.txt")
	marker, err := findMessage(lines[0])
	if err != nil {
		log.Fatal(err)
	}
	exp := 19
	if marker != exp {
		t.Errorf("expected %d, got %d\n", exp, marker)
	}

	marker, err = findMessage(lines[1])
	if err != nil {
		log.Fatal(err)
	}
	exp = 23
	if marker != exp {
		t.Errorf("expected %d, got %d\n", exp, marker)
	}

	marker, err = findMessage(lines[2])
	if err != nil {
		log.Fatal(err)
	}
	exp = 23
	if marker != exp {
		t.Errorf("expected %d, got %d\n", exp, marker)
	}

	marker, err = findMessage(lines[3])
	if err != nil {
		log.Fatal(err)
	}
	exp = 29
	if marker != exp {
		t.Errorf("expected %d, got %d\n", exp, marker)
	}

	marker, err = findMessage(lines[4])
	if err != nil {
		log.Fatal(err)
	}
	exp = 26
	if marker != exp {
		t.Errorf("expected %d, got %d\n", exp, marker)
	}

}
