// Package main is the executable
package main

import (
	"testing"
)

func TestGetLines(t *testing.T) {
	lines := getData("test.txt")
	count := len(lines)
	exp := 25

	if count != exp {
		t.Errorf("expected %d, got %d\n", exp, count)
	}
}

func TestGetNums(t *testing.T) {
	lines := getData("test.txt")
	nums := getNums(lines)
	count := len(nums)
	exp := 25

	if count != 25 {
		t.Errorf("expected %d, got %d\n", exp, count)
	}
	if nums[0] != 20 {
		t.Errorf("expected %d, got %d\n", 25, nums[0])
	}
	if nums[21] != 19 {
		t.Errorf("expected %d, got %d\n", 19, nums[21])
	}
}

func TestValid(t *testing.T) {
	lines := getData("test.txt")
	nums := getNums(lines)
	valid := isValid(26, nums, 25)
	if valid != true {
		t.Errorf("expected %t, got %t\n", true, valid)
	}
	valid = isValid(49, nums, 25)
	if valid != true {
		t.Errorf("expected %t, got %t\n", true, valid)
	}
	valid = isValid(100, nums, 25)
	if valid != false {
		t.Errorf("expected %t, got %t\n", false, valid)
	}
	valid = isValid(50, nums, 25)
	if valid != false {
		t.Errorf("expected %t, got %t\n", false, valid)
	}

	nums = append(nums, 45)
	valid = isValid(26, nums, 25)
	if valid != true {
		t.Errorf("expected %t, got %t\n", true, valid)
	}
	valid = isValid(65, nums, 25)
	if valid != false {
		t.Errorf("expected %t, got %t\n", false, valid)
	}
	valid = isValid(64, nums, 25)
	if valid != true {
		t.Errorf("expected %t, got %t\n", true, valid)
	}
	valid = isValid(66, nums, 25)
	if valid != true {
		t.Errorf("expected %t, got %t\n", true, valid)
	}
}

func TestGetFirstError(t *testing.T) {
	lines := getData("test2.txt")
	nums := getNums(lines)

	num := getFirstError(nums, 5)
	exp := 127
	if num != exp {
		t.Errorf("expected %d, got %d\n", exp, num)
	}

}

func TestContig(t *testing.T) {
	lines := getData("test2.txt")
	nums := getNums(lines)
	start, end := findContinguous(nums, 127)
	expStart := 2
	expEnd := 5

	if start != expStart {
		t.Errorf("expected %d, got %d\n", expStart, start)
	}
	if end != expEnd {
		t.Errorf("expected %d, got %d\n", expEnd, start)
	}
}

func TestGetWeakness(t *testing.T) {
	lines := getData("test2.txt")
	nums := getNums(lines)
	start, end := findContinguous(nums, 127)
	weakness := getWeakness(nums, start, end)
	exp := 62
	if weakness != exp {
		t.Errorf("expected %d, got %d\n", exp, weakness)
	}

}
