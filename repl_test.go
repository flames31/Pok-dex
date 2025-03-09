package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  Hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " TESTING FUNC",
			expected: []string{"testing", "func"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf(`---------------------------------
input:	%v
want:  	%v
got:    %v
`, c.input, c.expected, actual)
		}

		for i := range actual {
			got := actual[i]
			want := c.expected[i]

			if got != want {
				t.Errorf(`---------------------------------
input:	%v
want:  	%v
got:    %v
`, c.input, want, got)

			}
		}
		fmt.Printf(`---------------------------------
input:	%v
want:   %v
got:    %v
`, c.input, c.expected, actual)
	}

}
