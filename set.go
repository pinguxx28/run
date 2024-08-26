package main

import (
	"os"
	"fmt"
)

func set(path string, command string) {
	// print the command after the current working directory
	// inside of the .local/share/run file

	dir, err := os.Getwd()
	checkErr(err, "couldnt get working directory")

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0755) 
	checkErr(err, "cannot append file")

	line := fmt.Sprintf("%s %s\n", dir, command)
	_, err = file.WriteString(line)
	checkErr(err, "cannot write to file")

	err = file.Close()
	checkErr(err, "cannot close file")

	// verbose output
	fmt.Printf("Set command: [%s] for directory [%s]\n", command, dir)
}
