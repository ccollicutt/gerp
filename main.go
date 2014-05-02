package main

// Original grep type code came from here: 
// https://github.com/StefanSchroeder/Golang-Regex-Tutorial/blob/master/01-chapter3.markdown

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
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
func walk(re, path string) (files []string) {

	filepath.Walk(path, func(file string, info os.FileInfo, err error) error {
		if file != path {
			matches := grep(re, file)

			for _, match := range matches {
				fmt.Printf(file + ": ")
				fmt.Println(match)
			}
		}
		return nil
	})
	return nil
}

// like ls
func list(re, path string) (fileArray []os.FileInfo) {

	files, _ := ioutil.ReadDir(path)
	for _, file := range files {

		full_path := fmt.Sprint(path + "/" + file.Name())
		matches := grep(re, full_path)

		for _, match := range matches {
			fmt.Printf(full_path + ": ")
			fmt.Println(match)
		}
	}
	return files
}

func exists(path string) (err error) {
	_, err = os.Stat(path)
	return err
}

func main() {

	app := cli.NewApp()
	app.Name = "gerp"
	app.Usage = "gerp regex file"
	app.Version = "0.1"

	app.Flags = []cli.Flag{
		cli.BoolFlag{"recursive, r", "walk the directory"},
	}

	app.Action = func(c *cli.Context) {

		if len(c.Args()) != 2 {
			fmt.Println("Usage: gerp [-r|--recursive] pattern file|directory")
			os.Exit(1)
		} else {

			pattern := c.Args()[0]
			path := c.Args()[1]

			err := exists(path)
			if err != nil {
				fmt.Println(path + " does not exist")
				os.Exit(1)
			}

			if c.Bool("recursive") == true {
				walk(pattern, path)
			} else {
				list(pattern, path)
			}
		}

	}

	app.Run(os.Args)

}
