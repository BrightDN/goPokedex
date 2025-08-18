package main

import "testing"

func TestCleanInput(t *testing.T){
	cases := []struct{
		input string
		expected []string
	}{
		{
			input: "Hello world",
			expected: []string{"hello", "world"},
		},
		{
			input: "              Hello             world ",
			expected: []string{"hello", "world"},			
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Length of input does not match up")
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Word in slice does not match expected result")
				return
			}
		}
	}
}