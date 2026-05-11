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

func main() {

}