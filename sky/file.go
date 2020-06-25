package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	targetFolder string
	targetFile   string
	searchResult []string
)

func findFile(path string, fileInfo os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return nil
	}

	absolute, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if fileInfo.IsDir() {
		fmt.Println("Searching directory ...", absolute)
		testDir, err := os.Open(absolute)

		if err != nil {
			if os.IsPermission(err) {
				fmt.Println("No permission to scan ...", absolute)
				fmt.Println(err)
			}
		}
		testDir.Close()
		return nil
	} else {
		matched, err := filepath.Match(targetFile, fileInfo.Name())
		//matched, err := filepath.Glob(targetFile)
		if err != nil {
			fmt.Println(err)
		}
		if matched {
			add := "Found : " + absolute
			searchResult = append(searchResult, add)
		}
	}
	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("USAGE : %s <target_directory> <target_file> \n", os.Args[0])
		os.Exit(0)
	}

	targetFolder = os.Args[1]
	targetFile = os.Args[2]

	fmt.Println("Searching for [", targetFolder, "]")

	testFile, err := os.Open(targetFolder)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer testFile.Close()

	testFileInfo, _ := testFile.Stat()
	if !testFileInfo.IsDir() {
		fmt.Println(targetFolder, " is not a directory!")
		os.Exit(-1)
	}

	err = filepath.Walk(targetFolder, findFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println("\n\nFound ", len(searchResult), " hits!")
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@")

	for _, v := range searchResult {
		fmt.Println(v)
	}
}
