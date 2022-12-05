// Package main is the executable
package main

import (
	"testing"
)

func TestGetLines(t *testing.T) {
	lines := getData("test.txt")
	count := len(lines)
	exp := 9

	if count != exp {
		t.Errorf("expected %d, got %d\n", exp, count)
	}
}

func TestGetStacks(t *testing.T) {
	lines := getData("test.txt")
	stacks := getStacks(lines)
	size := len(stacks)
	exp := 3
	if size != exp {
		t.Errorf("expected %d, got %d\n", exp, size)
	}

	s := stacks[0]
	val, err := s.pop()
	if err != nil {
		t.Error(err)
	}
	if val != "N" {
		t.Errorf("expected %q, got %q\n", "N", val)
	}
	val, err = s.pop()
	if err != nil {
		t.Error(err)
	}
	if val != "Z" {
		t.Errorf("expected %q, got %q\n", "Z", val)
	}

	s = stacks[1]
	val, err = s.pop()
	if err != nil {
		t.Error(err)
	}
	if val != "D" {
		t.Errorf("expected %q, got %q\n", "D", val)
	}
	val, err = s.pop()
	if err != nil {
		t.Error(err)
	}
	if val != "C" {
		t.Errorf("expected %q, got %q\n", "C", val)
	}
	val, err = s.pop()
	if err != nil {
		t.Error(err)
	}
	if val != "M" {
		t.Errorf("expected %q, got %q\n", "M", val)
	}

	s = stacks[2]
	val, err = s.pop()
	if err != nil {
		t.Error(err)
	}
	if val != "P" {
		t.Errorf("expected %q, got %q\n", "P", val)
	}
}

func TestStacks(t *testing.T) {
	stack := newStack()
	if stack.size != 0 {
		t.Errorf("expected %d, got %d\n", 0, stack.size)
	}
	stack.push("a")
	if stack.size != 1 {
		t.Errorf("expected %d, got %d\n", 0, stack.size)
	}
	val, err := stack.peek()
	if err != nil {
		t.Error(err)
	}
	if val != "a" {
		t.Errorf("expected %q, got %q\n", "a", val)
	}
	if stack.size != 1 {
		t.Errorf("expected %d, got %d\n", 1, stack.size)
	}

	val, err = stack.pop()
	if err != nil {
		t.Error(err)
	}
	if val != "a" {
		t.Errorf("expected %q, got %q\n", "a", val)
	}
	if stack.size != 0 {
		t.Errorf("expected %d, got %d\n", 0, stack.size)
	}

	val, err = stack.pop()
	if err == nil {
		t.Errorf("popping an empty stack should error")
	}
}

func TestGetMoves(t *testing.T) {
	lines := getData("test.txt")
	moves := getMoves(lines)
	size := len(moves)
	exp := 4
	if size != exp {
		t.Errorf("expected %d, got %d\n", exp, size)
	}

	m := moves[2]
	if m.count != 2 {
		t.Errorf("expected %d, got %d\n", 2, m.count)
	}
	if m.from != 1 {
		t.Errorf("expected %d, got %d\n", 1, m.from)
	}
	if m.to != 0 {
		t.Errorf("expected %d, got %d\n", 0, m.to)
	}
}

func TestMakeMoves(t *testing.T) {
	lines := getData("test.txt")
	stacks := getStacks(lines)
	moves := getMoves(lines)
	makeMove(stacks, moves[2])

	if stacks[0].size != 4 {
		t.Errorf("expected %d, got %d\n", 3, stacks[0].size)
	}

	if stacks[1].size != 1 {
		t.Errorf("expected %d, got %d\n", 1, stacks[1].size)
	}

}

func TestMakeAllMoves(t *testing.T) {
	lines := getData("test.txt")
	stacks := getStacks(lines)
	moves := getMoves(lines)
	for _, m := range moves {
		makeMove(stacks, m)
	}

	top := getTop(stacks)
	exp := "CMZ"
	if top != exp {
		t.Errorf("expected %q, got %q\n", exp, top)
	}

}

func TestMakeAllMoves2(t *testing.T) {
	lines := getData("test.txt")
	stacks := getStacks(lines)
	moves := getMoves(lines)
	for _, m := range moves {
		makeMove2(stacks, m)
	}

	top := getTop(stacks)
	exp := "MCD"
	if top != exp {
		t.Errorf("expected %q, got %q\n", exp, top)
	}

}
