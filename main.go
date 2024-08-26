package main

import (
	"os"
	"strings"
)

func main() {
	runDir, runFile := setRunDirAndFile()
	ensureDirExists(runDir)
	ensureFileExists(runFile)

	args := os.Args

	if len(args) == 1 {
		run(runFile)
	} else {
		switch args[1] {
		case "set":
			command := strings.Join(args[2:], " ")
			set(runFile, command)
		}
	}
}
