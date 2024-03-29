package main

import (
	"testing"
	//	"fmt"
)

func Test_Run_1(t *testing.T) {
	args := []string{"./test", "hi"}
	// for some reason the bool arg from recursive := flag.Bool("r", false, "recursive")
	// in main is a pointer...

	// https://groups.google.com/forum/#!topic/golang-nuts/OELe3zXuPV4
    var tBool bool
    recursive :=&tBool
    *recursive = true

	matches := run(args, recursive)
	for _, match := range matches {
		if match.file != "test/test_grep.txt" {
			t.Error("run did not return the right file")
		} else {
			t.Log("run did return the right file")
		}
	}
}

func Test_Walk_1(t *testing.T) {
	allFiles := walk("./test")
	for _, file := range allFiles {
		t.Log(file)
		if file != "test/test_grep.txt" {
			t.Error("walk did not fine the right file")
		} else {
			t.Log("walk found the right file")
		}
	}
}

func Test_List_1(t *testing.T) {
	allFiles := list("./test")
	for _, file := range allFiles {
		t.Log(file)
		if file != "./test/test_grep.txt" {
			t.Log(file)
			t.Error("list did not find the right file")
		} else {
			t.Log("list found the right file")
		}
	}
}

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

func Test_Exists_1(t *testing.T) {
	err := exists("test/test_grep.txt")
	if err != nil {
		t.Error("exists did not pass test 1")
	} else {
		t.Log("exists test 1 passed")
	}
}

func Test_Exists_2(t *testing.T) {
	err := exists("test/doesnotexist.txt")
	if err == nil {
		t.Error("exists did not pass test 2")
	} else {
		t.Log("exists test 2 passed")
	}
}
