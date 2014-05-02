package main

import (
	"testing"
)


func Test_Grep_1(t *testing.T) {
	test_matches := []string{"hi"}
	matches := grep("hi", "./test/test_grep.txt")

	if test_matches[0] != matches[0] {
		t.Error("grep did not pass test 1")
	} else {
		t.Log("grep test 1 passed")
	}
}

func Test_Grep_2(t *testing.T) {
	matches := grep("adfasfd", "./test/test_grep.txt")

	// Expect this to fail, ie. not find a match
	if matches != nil {
		t.Error("grep did not pass test 2")
	} else {
		t.Log("grep test 2 passed")
	}
}