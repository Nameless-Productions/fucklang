package main

import (
	"fmt"
	"log"
	"strings"
)

type Command struct {
	Name string   `json:"name"`
	Args []string `json:"args"`
}

func tokenize(line string) []string {
	var tokens []string

	tokens = strings.Split(line, " ")
	return tokens;
}

func parse(source string) []Command {
	var commands []Command

	for _, line := range strings.Split(source, "\n") {
		line = strings.TrimSpace(line)
		
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}

		tokens := tokenize(line)
		if len(tokens) == 0{
			continue
		}

		cmd := Command{
			Name: tokens[0],
			Args: tokens[1:],
		}
		commands = append(commands, cmd)
	}

	return commands
}

func interpret(commands []Command) {
	vars := map[string]string{}

	for _, cmd := range commands{
		switch cmd.Name {
		case "print":
			fmt.Println(strings.Join(cmd.Args, " "))
		case "set":
			if len(cmd.Args) == 0 || len(cmd.Args) == 1 {
				log.Fatal("set requires at least 2 arguments")
			}
			vars[cmd.Args[0]] = strings.Join(cmd.Args[1:], " ")
		case "get":
			if len(cmd.Args) != 1 {
				log.Fatal("get requires 1 argument")
			}

			val, ok := vars[cmd.Args[0]]

			if ok {
				fmt.Println(val)
			} else {
				log.Fatalf("No variable under name %s", cmd.Args[0])
			}
		default:
			log.Fatalf("Unknows command: %s", cmd.Name)
		}
	}
}