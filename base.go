package main

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"io/fs"
	"errors"
	"strings"
)

func checkErr(err error, desc string) {
	if err != nil {
		fmt.Println(desc)
		log.Fatal(err)
	}
}

func doesPathExist(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, fs.ErrNotExist)
}

func getLineContaining(path string, text string) string {
	file, err := os.Open(path)
	checkErr(err, "cannot open file")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), text) {
			return scanner.Text()
		}
	}

	err = scanner.Err()
	checkErr(err, "scanner err")

	err = file.Close()
	checkErr(err, "cannot close file")

	return ""
}

func setRunDirAndFile() (string, string) {
	homeDir, err := os.UserHomeDir()
	checkErr(err, "cannot get user home directory")

	runDir := homeDir + "/.local/share/run"
	runFile := runDir + "/commands.txt"

	return runDir, runFile
}

func ensureDirExists(dir string) {
	if !doesPathExist(dir) {
		err := os.MkdirAll(dir, 0755)
		checkErr(err, "cannot create dir")

		fmt.Printf("Created directory [%s]\n", dir)
	}
}

func ensureFileExists(file string) {
	if !doesPathExist(file) {
		_, err := os.OpenFile(file, os.O_CREATE, 0755)
		checkErr(err, "cannot create file")

		fmt.Printf("Created file %s\n", file)
	}
}
