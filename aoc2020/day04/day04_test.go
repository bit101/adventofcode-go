package main

import "testing"

func TestGetMap(t *testing.T) {
	list := getList("test.txt")
	count := len(list)

	if count != 13 {
		t.Errorf("expected %d, got %d\n", 14, count)
	}
}

func TestGetPassports(t *testing.T) {
	list := getList("test.txt")
	passports := getPassports(list)
	count := len(passports)

	if count != 3 {
		t.Errorf("expected %d, got %d\n", 3, count)
	}
}

func TestValidate(t *testing.T) {
	list := getList("test.txt")
	passports := getPassports(list)
	count := 0
	for _, pp := range passports {
		valid := validatePassport(pp)
		if valid {
			count++
		}
	}

	if count != 2 {
		t.Errorf("expected %d, got %d\n", 2, count)
	}
}

func TestValidate2Valid(t *testing.T) {
	list := getList("valid.txt")
	passports := getPassports(list)
	count := 0
	for _, pp := range passports {
		valid := validatePassport2(pp)
		if valid {
			count++
		}
	}

	if count != 4 {
		t.Errorf("expected %d, got %d\n", 4, count)
	}
}

func TestValidate2Invalid(t *testing.T) {
	list := getList("invalid.txt")
	passports := getPassports(list)
	for _, pp := range passports {
		valid := validatePassport2(pp)
		if valid {
			t.Error("expected no valid passports")
		}
	}
}
