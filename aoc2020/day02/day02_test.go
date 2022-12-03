package main

import "testing"

func TestGetPasswords(t *testing.T) {
	passwords := getPasswords("test.txt")
	count := len(passwords)

	if count != 3 {
		t.Errorf("expected %d, got %d\n", 3, count)
	}
}

func TestValidate(t *testing.T) {
	passwords := getPasswords("test.txt")

	val := validatePassword(passwords[0])
	if !val {
		t.Errorf("expected %t, got %t\n", true, val)
	}

	val = validatePassword(passwords[1])
	if val {
		t.Errorf("expected %t, got %t\n", false, val)
	}

	val = validatePassword(passwords[2])
	if !val {
		t.Errorf("expected %t, got %t\n", true, val)
	}
}

func TestParsePassword(t *testing.T) {
	passwords := getPasswords("test.txt")

	pw := passwords[0]
	if pw.min != 1 && pw.max != 3 && pw.char != 'a' && pw.pw != "abcde" {
		t.Errorf("expected %d, %d, %c, %s, got %d, %d, %c, %s\n", 1, 3, 'a', "abcde", pw.min, pw.max, pw.char, pw.pw)
	}

	pw = passwords[1]
	if pw.min != 1 && pw.max != 3 && pw.char != 'b' && pw.pw != "cdefg" {
		t.Errorf("expected %d, %d, %c, %s, got %d, %d, %c, %s\n", 1, 3, 'b', "cdefg", pw.min, pw.max, pw.char, pw.pw)
	}

	pw = passwords[2]
	if pw.min != 2 && pw.max != 9 && pw.char != 'c' && pw.pw != "ccccccccc" {
		t.Errorf("expected %d, %d, %c, %s, got %d, %d, %c, %s\n", 1, 3, 'c', "ccccccccc", pw.min, pw.max, pw.char, pw.pw)
	}
}

func TestValidate2(t *testing.T) {
	passwords := getPasswords("test.txt")

	val := validatePassword2(passwords[0])
	if !val {
		t.Errorf("expected %t, got %t\n", true, val)
	}

	val = validatePassword2(passwords[1])
	if val {
		t.Errorf("expected %t, got %t\n", false, val)
	}

	val = validatePassword2(passwords[2])
	if val {
		t.Errorf("expected %t, got %t\n", false, val)
	}
}
