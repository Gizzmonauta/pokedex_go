package main

import (
	"bufio"
	"fmt"
	"os"
	"net/http"
	"encoding/json"
)

type cliCommand struct {
	name 		string
	description string
	callback	func() error
}

var registry map[string]cliCommand

type cliLocationArea struct{
	Count int
	Next  string
	Previous string
	Results []struct{
		Name string
		Url  string
	}
}

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

func commandMap() error {
	fullURL := "https://pokeapi.co/api/v2/location-area/"

	// Create a new request using http.NewRequest
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return err
	}

	// Make the request using the http.Client's Do method.
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Decode and return the response's JSON body (which is also a User)
	var la cliLocationArea
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&la)
	if err != nil {
		return err
	}

	locationAreas := la.Results
	for _, area := range locationAreas {
		fmt.Printf("- %s\n", area.Name)
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
		"map": {
			name: 		 "map",
			description: "Displays the map",
			callback: 	 commandMap,
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

