package main

import(
	"testing"
)

func TestCleanInput(t *testing.T){
	cases := []struct{
		input string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input: "Hi thEre    Everybody how are yOu doiNg?  ",
			expected: []string{"hi", "there", "everybody", "how", "are", "you", "doing?"},
		},
		{
			input: "    ???????????????????",
			expected: []string{"???????????????????"},
		},
		{
			input: "       ",
			expected: []string{},
		},
	}

	for k, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		if len(actual) != len(c.expected){
			t.Errorf("case %d: string length is not match. Expected: %v. Got: %v", k+1, len(c.expected), len(actual))
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
		
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			if word != expectedWord{
				t.Errorf("case %d: string word does not match. Expected: %v. Got: %v", k+1, expectedWord, word)
			}
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			}
	}

}

