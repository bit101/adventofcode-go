// Package main is the executable
package main

import (
	"testing"
)

func TestGetLines(t *testing.T) {
	lines := getLines("test.txt")
	count := len(lines)
	exp := 9

	if count != exp {
		t.Errorf("expected %d, got %d\n", exp, count)
	}
}

func TestInstructions(t *testing.T) {
	expNames := []string{"nop", "acc", "jmp", "acc", "jmp", "acc", "acc", "jmp", "acc"}
	expVals := []int{0, 1, 4, 3, -3, -99, 1, -4, 6}
	lines := getLines("test.txt")
	for i, line := range lines {
		ins := newInstruction(line)
		if ins.name != expNames[i] {
			t.Errorf("expected %q, got %q\n", expNames[i], ins.name)
		}
		if ins.value != expVals[i] {
			t.Errorf("expected %d, got %d\n", expVals[i], ins.value)
		}
	}
}

func TestGetInstructions(t *testing.T) {
	lines := getLines("test.txt")
	ins := getInstructions(lines)
	size := len(ins)
	exp := 9
	if size != exp {
		t.Errorf("expected %d, got %d\n", exp, size)
	}
}

func TestFindLoop(t *testing.T) {
	lines := getLines("test.txt")
	ins := getInstructions(lines)
	acc := findLoop(ins)
	exp := 5
	if acc != exp {
		t.Errorf("expected %d, got %d\n", exp, acc)
	}
}

func TestFindLoopOrEnd(t *testing.T) {
	lines := getLines("test.txt")
	ins := getInstructions(lines)

	ins[0].name = "jmp"
	end, acc := findLoopOrEnd(ins)
	expEnd := false
	expAcc := 0
	if end != expEnd || acc != expAcc {
		t.Errorf("expected end: %t and acc: %d. Got end: %t and acc: %d\n", expEnd, expAcc, end, acc)
	}

	ins = getInstructions(lines)
	ins[2].name = "nop"
	end, _ = findLoopOrEnd(ins)
	expEnd = false
	if end != expEnd {
		t.Errorf("expected end: %t. Got end: %t \n", expEnd, end)
	}

	ins = getInstructions(lines)
	ins[4].name = "nop"
	end, _ = findLoopOrEnd(ins)
	expEnd = false
	if end != expEnd {
		t.Errorf("expected end: %t. Got end: %t \n", expEnd, end)
	}

	ins = getInstructions(lines)
	ins[7].name = "nop"
	end, acc = findLoopOrEnd(ins)
	expEnd = true
	expAcc = 8
	if end != expEnd || acc != expAcc {
		t.Errorf("expected end: %t and acc: %d. Got end: %t and acc: %d\n", expEnd, expAcc, end, acc)
	}
}

func TestChangeInstruction(t *testing.T) {
	lines := getLines("test.txt")
	ins := getInstructions(lines)

	index := 0
	changeInstruction(index, ins)
	newIns := ins[index]
	exp := "jmp"
	if newIns.name != exp {
		t.Errorf("expected %q, got %q\n", exp, newIns.name)
	}
	changeInstruction(index, ins)
	newIns = ins[index]
	exp = "nop"
	if newIns.name != exp {
		t.Errorf("expected %q, got %q\n", exp, newIns.name)
	}

	index = 2
	changeInstruction(index, ins)
	newIns = ins[index]
	exp = "nop"
	if newIns.name != exp {
		t.Errorf("expected %q, got %q\n", exp, newIns.name)
	}

	index = 4
	changeInstruction(index, ins)
	newIns = ins[index]
	exp = "nop"
	if newIns.name != exp {
		t.Errorf("expected %q, got %q\n", exp, newIns.name)
	}

	index = 7
	changeInstruction(index, ins)
	newIns = ins[index]
	exp = "nop"
	if newIns.name != exp {
		t.Errorf("expected %q, got %q\n", exp, newIns.name)
	}
}

func TestFindNextChange(t *testing.T) {
	lines := getLines("test.txt")
	ins := getInstructions(lines)
	index := -1

	exp := 0
	index, err := findNextChange(index, ins)
	if err != nil {
		t.Error(err)
	}
	if index != exp {
		t.Errorf("expected %d, got %d\n", exp, index)
	}

	exp = 2
	index, err = findNextChange(index, ins)
	if err != nil {
		t.Error(err)
	}
	if index != exp {
		t.Errorf("expected %d, got %d\n", exp, index)
	}

	exp = 4
	index, err = findNextChange(index, ins)
	if err != nil {
		t.Error(err)
	}
	if index != exp {
		t.Errorf("expected %d, got %d\n", exp, index)
	}

	exp = 7
	index, err = findNextChange(index, ins)
	if err != nil {
		t.Error(err)
	}
	if index != exp {
		t.Errorf("expected %d, got %d\n", exp, index)
	}

	index, err = findNextChange(index, ins)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestFindEnd(t *testing.T) {
	lines := getLines("test.txt")
	ins := getInstructions(lines)
	acc := findEnd(ins)
	exp := 8
	if acc != exp {
		t.Errorf("expected %d, got %d\n", exp, acc)
	}

}
