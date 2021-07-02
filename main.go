package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func dirTree(output io.Writer, path string, printFiles bool) error {
	firstPath, _ := os.Getwd()

	if err := os.Chdir(path); err != nil {
		return fmt.Errorf("error - directory does not exist")
	}

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	currPath, _ := os.Getwd()
	tempFile := ""
	var tempSize int64
	for _, f := range files {
		if tempFile != "" {
			printDirFiles(output, tempFile, "├───", tempSize)
			os.Chdir(currPath)
		}

		tempFile = f.Name()
		tempSize = f.Size()
	}
	printDirFiles(output, tempFile, "└───", tempSize)
	os.Chdir(firstPath)

	return nil
}

func printDirFiles(output io.Writer, dir string, tab string, size int64) {
	if err := os.Chdir(dir); err != nil {
		if size == 0 {
			fmt.Fprintln(output, tab+dir+" (empty)")
		} else {
			fmt.Fprintln(output, tab+dir+" ("+strconv.FormatInt(size, 10)+"b)")
		}
		return
	}
	fmt.Fprintln(output, tab+dir)

	if tab[len(tab)-len("└───"):] == "└───" && len(tab)-len("└───") != 0 {
		tab = tab[:len(tab)-len("└───")] + "\t" + "├───"
	} else if tab[len(tab)-len("└───"):] != "└───" {
		tab = "│\t" + tab
	}

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	currPath, _ := os.Getwd()
	tempFile := ""
	var tempSize int64
	for _, f := range files {
		if tempFile != "" {
			printDirFiles(output, tempFile, tab, tempSize)
			os.Chdir(currPath)
		}

		tempFile = f.Name()
		tempSize = f.Size()
	}

	printDirFiles(output, tempFile, tab[:len(tab)-len("└───")]+"└───", tempSize)
}

func main() {
	out := os.Stdout

	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
