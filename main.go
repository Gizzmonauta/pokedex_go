package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Gizzmonauta/pokedex_go/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

var registry map[string]cliCommand

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, comm := range registry {
		fmt.Printf("%s: %s\n", comm.name, comm.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	// Create a new request using http.NewRequest
	req, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = req.Next
	cfg.prevLocationsURL = req.Previous

	for _, area := range req.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	return nil

}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	// Create a new request using http.NewRequest
	req, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = req.Next
	cfg.prevLocationsURL = req.Previous

	for _, area := range req.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	return nil

}

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
	}

	registry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokemon> ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)

		if comm, ok := registry[words[0]]; ok {
			err := comm.callback(cfg)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command:", words[0])
		}
	}
}
