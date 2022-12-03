package main

import "testing"

func TestGetExpenses(t *testing.T) {
	expenses := getExpenses("test.txt")
	count := len(expenses)

	if count != 6 {
		t.Errorf("expected %d, got %d\n", 6, count)
	}

	val := expenses[0]
	if val != 1721 {
		t.Errorf("expected %d, got %d\n", 1721, val)
	}
	val = expenses[1]
	if val != 979 {
		t.Errorf("expected %d, got %d\n", 979, val)
	}
	val = expenses[2]
	if val != 366 {
		t.Errorf("expected %d, got %d\n", 366, val)
	}
	val = expenses[3]
	if val != 299 {
		t.Errorf("expected %d, got %d\n", 299, val)
	}
	val = expenses[4]
	if val != 675 {
		t.Errorf("expected %d, got %d\n", 675, val)
	}
	val = expenses[5]
	if val != 1456 {
		t.Errorf("expected %d, got %d\n", 1456, val)
	}
}

func TestGet2020(t *testing.T) {
	expenses := getExpenses("test.txt")
	a, b, err := get2020(expenses)

	if err != nil {
		t.Error(err)
	}

	if a+b != 2020 {
		t.Errorf("expected %d and %d, got %d and %d\n", 1721, 299, a, b)
	}

	if a*b != 514579 {
		t.Errorf("expected %d, got %d\n", 514579, a*b)
	}
}

func TestGet2020x3(t *testing.T) {
	expenses := getExpenses("test.txt")
	a, b, c, err := get2020x3(expenses)
	if err != nil {
		t.Error(err)
	}

	if a+b+c != 2020 {
		t.Errorf("expected %d and %d, got %d and %d\n", 1721, 299, a, b)
	}

	if a*b*c != 241861950 {
		t.Errorf("expected %d, got %d\n", 241861950, a*b)
	}

}
