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

func getPokedexCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	commands := getPokedexCommands()
	fmt.Println("Welcome to the pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Printf("%s: %s\n", commands["help"].name, commands["help"].description)
	fmt.Printf("%s: %s\n", commands["exit"].name, commands["exit"].description)
	return nil
}

func commandExit() error {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}

func main() {

	pokedexCommands := getPokedexCommands()
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
			fmt.Println("Invalid command. Type 'help' for usage")
		}

	}
}
