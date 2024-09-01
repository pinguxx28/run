package main

import (
	"os"
	"fmt"
	"os/exec"
	"strings"
)

func run(file string) {
	dir, err := os.Getwd()
	checkErr(err, "cannot get working directory")

	originalCommand := getCommand(file, dir)
	commands := strings.Split(originalCommand, "&&")

	for _, command := range commands {
		trimmedCommand := strings.TrimSpace(command)
		commandArgs := strings.Fields(trimmedCommand)

		output, err := exec.Command(commandArgs[0], commandArgs[1:]...).Output()
		checkErr(err, "cannot exec command")

		fmt.Printf("+ %v\n", trimmedCommand)
		fmt.Print(string(output))
	}
}
