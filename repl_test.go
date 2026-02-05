package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input: "Squirtle Wartortle Blastoise",
			expected: []string{"squirtle", "wartortle", "blastoise"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) returned %d words, expected %d", c.input, len(actual), len(c.expected))
		}
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%q)[%d] = %q, expected %q", c.input, i, word, expectedWord)
			}
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}
}
