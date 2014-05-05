package main

// Original grep type code came from here:
// https://github.com/StefanSchroeder/Golang-Regex-Tutorial/blob/master/01-chapter3.markdown

import (
	"bufio"
	"fmt"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

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

		full_path := fmt.Sprint(path + "/" + file.Name())
		allFiles = append(allFiles, full_path)
	}
	return allFiles
}

func exists(path string) (err error) {
	_, err = os.Stat(path)
	return err
}

func main() {

	recursive := flag.Bool("r", false, "recursive")
	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		fmt.Println("Usage: gerp [-r] pattern directory")
	} else {
		pattern := args[0]
		dir := args[1]
		var allFiles []string
		var matches []string

		if *recursive == true {
			allFiles = walk(dir)
		} else {
			allFiles = list(dir)
		}
		for _, file := range allFiles {
			matches = grep(pattern, file)
			for _, match := range matches {
				fmt.Println(file + ": " + match)
			}
		}
	}
}
