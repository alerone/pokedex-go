package main

import "testing"

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
			input:    " hi good engineer ",
			expected: []string{"hi", "good", "engineer"},
		},
		{
			input: "    ",
			expected: []string{},
		},
	}
	
	for _,c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len( c.expected ){
			t.Errorf("length of \ncleanInput(%s): %d, \n\tdoes not match expected: %d", 
				c.input, len(actual), len(c.expected))
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("word expected: %s \n\tactual word: %s", expectedWord, word)
				return
			}
		}
	}
}
