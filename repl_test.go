package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  Hello World   ",
			expected: []string{"hello", "world"},
		},
		{
			input: "  Hi My name IS OnoLAX           ",
			expected: []string{"hi", "my", "name", "is", "onolax"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Wrong output")
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Wrong output")
				continue
			}
		}
	}
}