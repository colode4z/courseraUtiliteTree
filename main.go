package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
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

	mapOfDirsFiles := make(map[string]int64)

	for _, f := range files {
		if !f.IsDir() && !printFiles {
			continue
		}
		mapOfDirsFiles[f.Name()] = f.Size()
	}

	arr := make([]string, 0, len(mapOfDirsFiles))
	for k := range mapOfDirsFiles {
		arr = append(arr, k)
	}
	sort.Strings(arr)

	currPath, _ := os.Getwd()
	tempFile := ""
	var tempSize int64

	for _, name := range arr {
		if tempFile != "" {
			printDirFiles(output, tempFile, "├───", printFiles, tempSize)
			os.Chdir(currPath)
		}

		tempFile = name
		tempSize = mapOfDirsFiles[name]
	}
	if tempFile != "" {
		printDirFiles(output, tempFile, "└───", printFiles, tempSize)
		os.Chdir(firstPath)
	}

	return nil
}

func printDirFiles(output io.Writer, dir string, tab string, printFiles bool, sizeOfFile int64) {
	if err := os.Chdir(dir); err != nil {
		if sizeOfFile == 0 {
			fmt.Fprintln(output, tab+dir+" (empty)")
		} else {
			fmt.Fprintln(output, tab+dir+" ("+strconv.FormatInt(sizeOfFile, 10)+"b)")
		}
		return
	}
	fmt.Fprintln(output, tab+dir)

	//&& len(tab)-len("└───") != 0
	if tab[len(tab)-len("└───"):] == "└───" {
		tab = tab[:len(tab)-len("└───")] + "\t" + "├───"
	} else if tab[len(tab)-len("└───"):] != "└───" {
		tab = "│\t" + tab
	}

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	mapOfDirsFiles := make(map[string]int64)

	for _, f := range files {
		if !f.IsDir() && !printFiles {
			continue
		}
		mapOfDirsFiles[f.Name()] = f.Size()
	}

	arr := make([]string, 0, len(mapOfDirsFiles))
	for k := range mapOfDirsFiles {
		arr = append(arr, k)
	}
	sort.Strings(arr)

	currPath, _ := os.Getwd()
	tempFile := ""
	var tempSize int64

	for _, name := range arr {
		if tempFile != "" {
			printDirFiles(output, tempFile, tab, printFiles, tempSize)
			os.Chdir(currPath)
		}

		tempFile = name
		tempSize = mapOfDirsFiles[name]
	}

	if tempFile != "" {
		printDirFiles(output, tempFile, tab[:len(tab)-len("└───")]+"└───", printFiles, tempSize)
	}
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
