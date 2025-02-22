package pokedexcli

import "strings"

func cleanInput(text string) []string {
	// First, trime the spaces from the start and end
	cleaned := strings.TrimSpace(text)
	cleaned = strings.ToLower(cleaned)
	return strings.Fields(cleaned)
}