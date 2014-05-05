package main

// Original grep type code came from here:
// https://github.com/StefanSchroeder/Golang-Regex-Tutorial/blob/master/01-chapter3.markdown

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

type Match struct {
	file  string
	match string
}

func grep(re, filename string) (matches []string) {
	regex, err := regexp.Compile(re)
	if err != nil {
		return // there was a problem with the regular expression.
	}

	fh, err := os.Open(filename)
	f := bufio.NewReader(fh)

	if err != nil {
		return // there was a problem opening the file.
	}
	defer fh.Close()

	buf := make([]byte, 1024)
	for {
		buf, _, err = f.ReadLine()
		if err != nil {
			return
		}

		s := string(buf)
		if regex.MatchString(s) {
			matches = append(matches, string(buf))
		}
	}
	return matches
}

// like ls recursive
func walk(path string) (allFiles []string) {
	filepath.Walk(path, func(file string, info os.FileInfo, err error) error {
		if file != path {
			allFiles = append(allFiles, file)
		}
		return nil
	})
	return allFiles
}

// like ls
func list(path string) (allFiles []string) {
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {

		fullPath := fmt.Sprint(path + "/" + file.Name())
		allFiles = append(allFiles, fullPath)
	}
	return allFiles
}

func exists(path string) (err error) {
	_, err = os.Stat(path)
	return err
}

func run(args []string, recursive *bool) (matches []Match) {

	if len(args) != 2 {
		fmt.Println("Usage: gerp [-r] directory pattern")
	} else {
		pattern := args[1]
		dir := args[0]
		var allFiles []string
		var grepMatches []string

		if *recursive == true {
			allFiles = walk(dir)
		} else {
			allFiles = list(dir)
		}
		for _, file := range allFiles {
			grepMatches = grep(pattern, file)
			for _, match := range grepMatches {
				fileMatch := Match{file, match}
				matches = append(matches, fileMatch)
			}
		}
		return matches
	}
	return nil
}

func main() {

	recursive := flag.Bool("r", false, "recursive")
	flag.Parse()
	args := flag.Args()

	allMatches := run(args, recursive)
	for _, match := range allMatches {
		//fmt.Println(file + ": " + match)
		fmt.Printf(match.file + ": ")
		fmt.Println(match.match)
	}
}
