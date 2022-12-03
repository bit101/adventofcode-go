package main

import "testing"

func TestGetMap(t *testing.T) {
	mymap := getMap("test.txt")
	count := len(mymap)

	if count != 11 {
		t.Errorf("expected %d, got %d\n", 11, count)
	}
}

func TestHasTree(t *testing.T) {
	mymap := getMap("test.txt")

	// top left
	b := hasTree(mymap, 0, 0)
	if b {
		t.Errorf("expected %t, got %t\n", false, b)
	}

	// has tree
	b = hasTree(mymap, 2, 0)
	if !b {
		t.Errorf("expected %t, got %t\n", true, b)
	}

	// wrap
	b = hasTree(mymap, 14, 0)
	if !b {
		t.Errorf("expected %t, got %t\n", true, b)
	}

	// wrap
	b = hasTree(mymap, 12, 0)
	if b {
		t.Errorf("expected %t, got %t\n", false, b)
	}

	// lower row
	b = hasTree(mymap, 7, 8)
	if !b {
		t.Errorf("expected %t, got %t\n", true, b)
	}

	// lower row wrapped
	b = hasTree(mymap, 12, 10)
	if !b {
		t.Errorf("expected %t, got %t\n", true, b)
	}

	// lower row wrapped
	b = hasTree(mymap, 13, 10)
	if b {
		t.Errorf("expected %t, got %t\n", false, b)
	}

}

func TestTraverse(t *testing.T) {
	mymap := getMap("test.txt")
	count := traverse(0, 0, 1, 1, mymap)
	if count != 2 {
		t.Errorf("expected %d, got %d\n", 2, count)
	}

	count = traverse(0, 0, 3, 1, mymap)
	if count != 7 {
		t.Errorf("expected %d, got %d\n", 11, count)
	}

	count = traverse(0, 0, 5, 1, mymap)
	if count != 3 {
		t.Errorf("expected %d, got %d\n", 3, count)
	}

	count = traverse(0, 0, 7, 1, mymap)
	if count != 4 {
		t.Errorf("expected %d, got %d\n", 4, count)
	}

	count = traverse(0, 0, 1, 2, mymap)
	if count != 2 {
		t.Errorf("expected %d, got %d\n", 2, count)
	}

}
