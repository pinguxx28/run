# Run

When you run: `run set "go run main.go"`, the command `go run main.go`,
will be stored in .local/share/run/wd-cmds.txt. Then when you run: `run`,
run will search in .local/share/run/wd-cmds.txt for the current working directory,
and then run the command for the specified directory.
