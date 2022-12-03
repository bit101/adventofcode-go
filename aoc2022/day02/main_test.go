package main

import (
	"testing"
)

func TestGetLines(t *testing.T) {
	lines := getLines("test.txt")
	count := len(lines)

	if count != 8 {
		t.Errorf("expected %d, got %d\n", 14, count)
	}
}

func TestGetPlays(t *testing.T) {
	lines := getLines("test.txt")
	turn := getPlays(lines[0])
	if turn.one != rock {
		t.Errorf("expected %v, got %v\n", rock, turn.one)
	}
	if turn.two != paper {
		t.Errorf("expected %v, got %v\n", paper, turn.two)
	}

	turn = getPlays(lines[1])
	if turn.one != paper {
		t.Errorf("expected %v, got %v\n", paper, turn.one)
	}
	if turn.two != rock {
		t.Errorf("expected %v, got %v\n", rock, turn.two)
	}

	turn = getPlays(lines[2])
	if turn.one != scissors {
		t.Errorf("expected %v, got %v\n", scissors, turn.one)
	}
	if turn.two != scissors {
		t.Errorf("expected %v, got %v\n", scissors, turn.two)
	}

	turn = getPlays(lines[3])
	if turn.one != rock {
		t.Errorf("expected %v, got %v\n", rock, turn.one)
	}
	if turn.two != rock {
		t.Errorf("expected %v, got %v\n", rock, turn.two)
	}

}

func TestScore(t *testing.T) {
	lines := getLines("test.txt")
	turn := getPlays(lines[0])
	score := getScore(turn)
	if score != 8 {
		t.Errorf("expected %d, got %d\n", 8, score)
	}

	turn = getPlays(lines[1])
	score = getScore(turn)
	if score != 1 {
		t.Errorf("expected %d, got %d\n", 1, score)
	}

	turn = getPlays(lines[2])
	score = getScore(turn)
	if score != 6 {
		t.Errorf("expected %d, got %d\n", 6, score)
	}

	turn = getPlays(lines[3])
	score = getScore(turn)
	if score != 4 {
		t.Errorf("expected %d, got %d\n", 4, score)
	}
}

func TestGetStrategy(t *testing.T) {
	lines := getLines("test.txt")

	turn := getStrategy(lines[0])
	if turn.one != rock {
		t.Errorf("expected %v, got %v\n", rock, turn.one)
	}
	if turn.result != draw {
		t.Errorf("expected %v, got %v\n", draw, turn.result)
	}

	turn = getStrategy(lines[1])
	if turn.one != paper {
		t.Errorf("expected %v, got %v\n", paper, turn.one)
	}
	if turn.result != lose {
		t.Errorf("expected %v, got %v\n", lose, turn.result)
	}

	turn = getStrategy(lines[2])
	if turn.one != scissors {
		t.Errorf("expected %v, got %v\n", scissors, turn.one)
	}
	if turn.result != win {
		t.Errorf("expected %v, got %v\n", win, turn.result)
	}

	turn = getStrategy(lines[3])
	if turn.one != rock {
		t.Errorf("expected %v, got %v\n", rock, turn.one)
	}
	if turn.result != lose {
		t.Errorf("expected %v, got %v\n", lose, turn.result)
	}

}

func TestGetMoveToMake(t *testing.T) {
	lines := getLines("test.txt")

	strat := getStrategy(lines[0])
	next := getMoveToMake(strat)
	if next != rock {
		t.Errorf("expected %v, got %v\n", rock, next)
	}

	strat = getStrategy(lines[1])
	next = getMoveToMake(strat)
	if next != rock {
		t.Errorf("expected %v, got %v\n", rock, next)
	}

	strat = getStrategy(lines[2])
	next = getMoveToMake(strat)
	if next != rock {
		t.Errorf("expected %v, got %v\n", rock, next)
	}

	strat = getStrategy(lines[3])
	next = getMoveToMake(strat)
	if next != scissors {
		t.Errorf("expected %v, got %v\n", scissors, next)
	}
}

func TestGetScore2(t *testing.T) {
	lines := getLines("test.txt")

	strat := getStrategy(lines[0])
	score := getScore2(strat)
	if score != 4 {
		t.Errorf("expected %d, got %d\n", 4, score)
	}

	strat = getStrategy(lines[1])
	score = getScore2(strat)
	if score != 1 {
		t.Errorf("expected %d, got %d\n", 1, score)
	}

	strat = getStrategy(lines[2])
	score = getScore2(strat)
	if score != 7 {
		t.Errorf("expected %d, got %d\n", 7, score)
	}

	strat = getStrategy(lines[3])
	score = getScore2(strat)
	if score != 3 {
		t.Errorf("expected %d, got %d\n", 3, score)
	}

	strat = getStrategy(lines[4])
	score = getScore2(strat)
	if score != 8 {
		t.Errorf("expected %d, got %d\n", 8, score)
	}

	strat = getStrategy(lines[5])
	score = getScore2(strat)
	if score != 5 {
		t.Errorf("expected %d, got %d\n", 5, score)
	}

	strat = getStrategy(lines[6])
	score = getScore2(strat)
	if score != 9 {
		t.Errorf("expected %d, got %d\n", 9, score)
	}

	strat = getStrategy(lines[7])
	score = getScore2(strat)
	if score != 2 {
		t.Errorf("expected %d, got %d\n", 2, score)
	}
}
