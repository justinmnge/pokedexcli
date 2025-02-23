package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	cleaned := strings.TrimSpace(text)
	cleaned = strings.ToLower(cleaned)
	return strings.Fields(cleaned)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Pokedex > ")
		scanner.Scan()          // this waits for user input
		input := scanner.Text() // this gets the text they typed
		cleanedInput := cleanInput(input)
		if len(cleanedInput) > 0 {
			fmt.Printf("Your command was: %s\n", cleanedInput[0])
		}
	}
}
