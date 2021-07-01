package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unicode/utf8"
)

func dirTree(out *os.File, path string, printFiles bool) error {
	if err := os.Chdir(path); err != nil {
		return fmt.Errorf("error - directory does not exist")
	}

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	currPath, _ := os.Getwd()
	tempFile := ""
	for _, f := range files {
		if tempFile != "" {
			printDirFiles(tempFile, "├───")
			os.Chdir(currPath)
		}

		tempFile = f.Name()
	}
	printDirFiles(tempFile, "└───")

	return nil
}

func printDirFiles(dir string, tab string) {
	fmt.Println(tab + dir)

	if err := os.Chdir(dir); err != nil {
		return
	}

	if tab[utf8.RuneCountInString(tab)-utf8.RuneCountInString("└───"):] == "└───" && utf8.RuneCountInString(tab)-utf8.RuneCountInString("└───") != 0 {
		tab = tab[:utf8.RuneCountInString(tab)-utf8.RuneCountInString("└───")] + " \t" + "└───"
	} else if tab[utf8.RuneCountInString(tab)-utf8.RuneCountInString("└───"):] != "└───" {
		tab = "|\t" + tab
	}

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	currPath, _ := os.Getwd()
	tempFile := ""
	for _, f := range files {
		if tempFile != "" {
			printDirFiles(tempFile, tab)
			os.Chdir(currPath)
		}

		tempFile = f.Name()
	}
	printDirFiles(tempFile, tab[:utf8.RuneCountInString(tab)-utf8.RuneCountInString("└───")]+"└───")
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
