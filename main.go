package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var pokedexCommands = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"exit": {
		name:        "exit",
		description: "Exits the pokedex",
		callback:    commandExit,
	},
}

func commandHelp() error {
	return nil
}

func commandExit() error {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("pokedex > ")

		if !scanner.Scan() {
			fmt.Println("Error, not able to scan user input. Exiting program.")
			os.Exit(1)
		}
		userInput := scanner.Text()
		if command, exists := pokedexCommands[userInput]; exists {
			command.callback()
		} else {
			fmt.Println("Invalid command. See Usage")
		}

	}
}
