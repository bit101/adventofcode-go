package main

import "testing"

func TestGetMap(t *testing.T) {
	list := getList("test.txt")
	count := len(list)

	if count != 4 {
		t.Errorf("expected %d, got %d\n", 4, count)
	}
}

func TestParseNums(t *testing.T) {
	list := getList("test.txt")

	num := parseBoardingNumber(list[0])
	exp := 357
	if num != exp {
		t.Errorf("expected %d, got %d\n", exp, num)
	}

	num = parseBoardingNumber(list[1])
	exp = 567
	if num != exp {
		t.Errorf("expected %d, got %d\n", exp, num)
	}

	num = parseBoardingNumber(list[2])
	exp = 119
	if num != exp {
		t.Errorf("expected %d, got %d\n", exp, num)
	}

	num = parseBoardingNumber(list[3])
	exp = 820
	if num != exp {
		t.Errorf("expected %d, got %d\n", exp, num)
	}
}

func TestHighest(t *testing.T) {
	list := getList("test.txt")
	highest := getHighest(list)

	exp := 820
	if highest != exp {
		t.Errorf("expected %d, got %d\n", exp, highest)
	}
}
