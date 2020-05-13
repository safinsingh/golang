package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func readFile(fileName string) (string, error) {
	fileContent, err := ioutil.ReadFile(fileName)
	return string(fileContent), err
}

func FileContainsRegex(fileName string, expressionString string) (bool, error) {
	fileContent, _ := readFile(fileName)
	matched, err := regexp.Match(expressionString, []byte(fileContent))
	return matched, err
}

func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func DirContainsRegex(dirName string, expressionString string) (bool, error) {
	files, err := FilePathWalkDir(dirName)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		trueDat, err := FileContainsRegex(file, expressionString)
		if err != nil {
			return false, err
		}
		return trueDat, nil
	}
	return true, nil
}

func main() {
	test, err := DirContainsRegex(`C:\Users\Safin\Documents\golang\dir-contains-regex\test`, `bleh`)
	if err != nil {
		panic(err)
	}
	fmt.Println(strconv.FormatBool(test))
}
