package main

import (
	"testing"
)

// func TestMain(t *testing.T) {
// 	cases := []struct {
// 		input    string
// 		expected string
// 	}{
// 		{
// 			input:    "hello world",
// 			expected: "hello",
// 		},
// 		{
// 			input:    "HeLlO WoRlD",
// 			expected: "hello",
// 		},
// 		{
// 			input:    "Sisters is the best army",
// 			expected: "sisters",
// 		},
// 	}
// }

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Bulbasaur, Rhyhorn, pikachu",
			expected: []string{"bulbasaur", "rhyhorn", "pikachu"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(c.expected) != len(actual) {
			t.Errorf("slice length mismatch, input length %v, expected %v", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word mismatch, got %s, expected %s", word, expectedWord)
			}
		}

	}
}
