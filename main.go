package main

import (
	"fmt"
	"goreutils/cmd"
	"os"
	"path"
)

var cmds = map[string]func([]string){
	"cat":  cmd.Cat,
	"echo": cmd.Echo,
}

func main() {
	args := os.Args
	prog := path.Base(args[0])

	var cmdName string
	var cmdArgs []string

	switch {
	case len(args) > 1 && isKnownCommand(args[1]):
		// Called like: goreutils cat file.txt
		cmdName = args[1]
		cmdArgs = args[2:]
	case isKnownCommand(prog):
		// Called via symlink: cat file.txt
		cmdName = prog
		cmdArgs = args[1:]
	default:
		printUsage(prog)
		os.Exit(1)
	}

	runCommand(cmdName, cmdArgs)
}

func isKnownCommand(name string) bool {
	_, ok := cmds[name]
	return ok
}

func runCommand(name string, args []string) {
	if fn, ok := cmds[name]; ok {
		fn(args)
	} else {
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", name)
		os.Exit(1)
	}
}

func printUsage(prog string) {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "  %s <command> [args...]\n", prog)
}
