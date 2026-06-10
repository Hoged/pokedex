package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		rawInput := scanner.Text()
		commands := strings.Fields(strings.ToLower(rawInput))
		fmt.Printf("Your command was: %v\n", commands[0])
	}
}

