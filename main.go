package main

import "strings"

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

func main() {

}