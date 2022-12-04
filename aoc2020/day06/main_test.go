// Package main is the executable
package main

import (
	"testing"
)

func TestGetLines(t *testing.T) {
	lines := getLines("test.txt")
	count := len(lines)
	exp := 15

	if count != exp {
		t.Errorf("expected %d, got %d\n", exp, count)
	}
}

func TestGetGroups(t *testing.T) {
	lines := getLines("test.txt")
	groups := getGroups(lines)
	count := len(groups)
	exp := 5
	if count != exp {
		t.Errorf("expected %d, got %d\n", exp, count)
	}
}

func TestGetAnswers(t *testing.T) {
	expCount := []int{3, 3, 3, 1, 1}
	expAnswers := [][]string{
		{"a", "b", "c"},
		{"a", "b", "c"},
		{"a", "b", "c"},
		{"a"},
		{"b"},
	}
	lines := getLines("test.txt")
	groups := getGroups(lines)
	for i, g := range groups {
		answers := getAnswers(g)
		if len(answers) != expCount[i] {
			t.Errorf("expected %d, got %d\n", expCount[i], len(answers))
		}

		for j, a := range answers {
			if a != expAnswers[i][j] {
				t.Errorf("expected %s, got %s\n", expAnswers[i][j], a)
			}
		}
	}
}

func TestScore(t *testing.T) {
	exp := []int{3, 3, 3, 1, 1}
	lines := getLines("test.txt")
	groups := getGroups(lines)
	for i, g := range groups {
		a := getScore(g)
		if a != exp[i] {
			t.Errorf("expected %d, got %d\n", exp[i], a)
		}
	}
}

func TestAllAnswerCount(t *testing.T) {
	exp := []int{3, 0, 1, 1, 1}
	lines := getLines("test.txt")
	groups := getGroups(lines)
	for i, g := range groups {
		c := getAllAnswersCount(g)
		if c != exp[i] {
			t.Errorf("expected %d, got %d\n", exp[i], c)
		}
	}

}
