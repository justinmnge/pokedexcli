package pokedexcli

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input: "hello  world",
			expected: []string{"hello", "world"}, 
		},
		{
			input: " hello world",
			expected: []string{"hello", "world"},
		},
	}
	
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use the t.Errorf to print an error message
		// and fail the test
		if len(c.expected) != len(actual) {
			t.Errorf("Got length %v, wanted length %v", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Got word %v, wanted word %v", word, expectedWord)
			}
		}
	}
}