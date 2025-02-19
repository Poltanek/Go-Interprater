package main

import (
	"fmt"
	"os"
)

func main() {
	argCount := len(os.Args)

	if argCount > 2 {
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	} else if argCount == 2 {
		runScript(os.Args[1])
	} else {
		startREPL()
	}
}

func runScript(script string) {
	fmt.Println("Executing Script: ", script)
}

func startREPL() {
	fmt.Println("Starting REPL")
}

fmt.Println("Hello World")