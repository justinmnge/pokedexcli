package main

import (
	"bufio" // For reading user input
	"encoding/json"
	"fmt" // For printing output
	"net/http"
	"os"      // For systems operations like exit
	"strings" // For string manipulation
)

func cleanInput(text string) []string {
	cleaned := strings.TrimSpace(text) // Removes whitespace
	cleaned = strings.ToLower(cleaned) // Converts to lowercase
	return strings.Fields(cleaned)     // Splits into words
}

type cliCommand struct { // Blueprint for our commands
	name        string              // The command word users will type
	description string              // Help text explaining what the command does
	callback    func(*config) error // Now accepts a pointer to a config struct
}

type config struct {
	Next     *string // URL for the next set of locations
	Previous *string // URL for the previous set of locations
}

var commands map[string]cliCommand // A map where keys are strings(the command words) and values are the cliCommand structs

func main() {
	scanner := bufio.NewScanner(os.Stdin) // Creates a new scanner to read user input from standard input (keyboard)

	conf := &config{} // Initialize an empty config struct

	commands = map[string]cliCommand{ // Initializes our commands map with two commands: exit and help
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    func(conf *config) error { return commandExit(conf) },
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 locations",
			callback:    commandMapb,
		},
	}

	// Infinite loop for our REPL (read-eval-print loop)
	for {
		fmt.Println("Pokedex > ")         // Prints the prompt
		scanner.Scan()                    // Waits for and reads the next line of input
		input := scanner.Text()           // Gets the text that way typed
		cleanedInput := cleanInput(input) // Cleans the input (removes spaces, converts to lowercase, splits into words)
		if len(cleanedInput) == 0 {       // If nothing was typed, start the loop over
			continue
		}

		if cmd, ok := commands[cleanedInput[0]]; ok { // Tries to find the command in the map
			err := cmd.callback(conf) // If found, exectute the command's callback
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command") // If command not found, print error
		}
	}
}

// Automatically picks up new commands because it iterates over the map
func commandHelp(conf *config) error {
	fmt.Println("Welcome to the Pokedex!") // Prints the welcome message
	fmt.Println("Usage:")                  // Prints the usuage header
	fmt.Println()                          // Prints a blank line for formatting

	// Loops through each command in our commands map
	for _, cmd := range commands { // For each command, prints its name and description
		fmt.Printf("%s: %s\n", cmd.name, cmd.description) // %s are placesholders for strings that get replaced with cmd
	}
	return nil // Returns nil because no errors occurred
}

func commandMap(conf *config) error {
	// Check if conf.Next is nil and initialize it with the default URL
	url := "https://pokeapi.co/api/v2/location-area"
	if conf == nil {
		return fmt.Errorf("configuration is nil")
	}

	if conf.Next != nil {
		url = *conf.Next
	}

	fmt.Println("You're on the last page of locations.")

	// Make HTTP request to the `Next` URL
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch locations: %v", err)
	}
	defer resp.Body.Close()

	// Decode response into struct
	var apiResponse struct {
		Results []struct {
			Name string `json:"name"`
		}
		Next     *string `json:"next"`
		Previous *string `json:"previous"`
	}
	decoder := json.NewDecoder(resp.Body) // Use decoder for response
	if err := decoder.Decode(&apiResponse); err != nil {
		return fmt.Errorf("failed to decode response: %v", err)
	}

	// Print location names
	for _, location := range apiResponse.Results {
		fmt.Println(location.Name)
	}

	// Update the config struct
	conf.Next = apiResponse.Next
	conf.Previous = apiResponse.Previous

	return nil
}

func commandMapb(conf *config) error {
	if conf.Previous == nil {
		fmt.Println("You are at the first page")
		return nil
	}

	resp, err := http.Get(*conf.Previous)
	if err != nil {
		return fmt.Errorf("failed to fetch locations: %v", err)
	}
	defer resp.Body.Close()

	// Decode response into struct
	var apiResponse struct {
		Results []struct {
			Name string `json:"name"`
		}
		Next     *string `json:"next"`
		Previous *string `json:"previous"`
	}
	decoder := json.NewDecoder(resp.Body) // Use decoder for response
	if err := decoder.Decode(&apiResponse); err != nil {
		return fmt.Errorf("failed to decode response: %v", err)
	}

	// Print location names
	for _, location := range apiResponse.Results {
		fmt.Println(location.Name)
	}

	// Update the config struct
	conf.Next = apiResponse.Next
	conf.Previous = apiResponse.Previous

	return nil
}

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!") // Prints goodbye message
	os.Exit(0)                                     // Exits the program with status code 0 (success)
	return nil                                     // This return statement never reach because os.Exit() terminates the program
}
