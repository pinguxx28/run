package main

import (
	"os"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func run(file string) {
	dir, err := os.Getwd()
	checkErr(err, "cannot get working directory")

	line := getLineContaining(file, dir)
	if line == "" {
		log.Fatalln("This directory doesn't have a command set")
	}

	index := strings.Index(line, " ")
	if index == -1 {
		log.Fatalf("Line is weird: [%v]\n", line)
	}

	command := line[index+1:]
	commandArgs := strings.Fields(command)
	
	var e *exec.ExitError
	output, e := exec.Command(commandArgs[0], commandArgs[1:]...).Output()
	checkErr(e, e.Stderr)
	/*
	cmd := exec.Command("sleep", "-u")
	err := cmd.Run()
	var exerr *exec.ExitError
	if errors.As(err, &exerr) {
		fmt.Printf("the command exited unsuccessfully: %d\n", exerr.ExitCode())
	}
	*/
	fmt.Printf("+ %v\n", command)
	fmt.Print(string(output))
}
