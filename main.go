package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name 		string
	description string
	callback	func() error
}

var registry map[string]cliCommand

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, comm := range registry {
		fmt.Printf("%s: %s\n", comm.name, comm.description)
	}
	return nil
}

func main() {
	registry = map[string]cliCommand{
		"exit": {
			name: 		 "exit",
			description: "Exit the Pokedex",
			callback: 	 commandExit,
		},
		"help": {
			name: 		 "help",
			description: "Displays a help message",
			callback: 	 commandHelp,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for ;; {
		fmt.Printf("Pokemon> ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)

		if comm, ok := registry[words[0]]; ok {
			err := comm.callback()
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command:", words[0])
		}
	}
}

