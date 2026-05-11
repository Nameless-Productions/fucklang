package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("no filename provided")
	}

	fileName := args[0]

	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	commands := parse(string(file))
	interpret(commands)
}