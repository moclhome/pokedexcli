package internal

import "testing"

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{
		"leadingTrailingSpaces": {
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		"simple": {
			input:    "My name is Lucy",
			expected: []string{"my", "name", "is", "lucy"},
		},
		"onlyOne": {
			input:    "OneWord",
			expected: []string{"oneword"},
		},
		"empty": {
			input:    "",
			expected: []string{""},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			actual := CleanInput(c.input)
			if len(actual) != len(c.expected) {
				t.Fatalf("expected slice length: %d, got: %d", len(c.expected), len(actual))
			}
			for i := range actual {
				word := actual[i]
				expectedWord := c.expected[i]
				if word != expectedWord {
					t.Fatalf("expected: %s, got: %s", c.expected, actual)
				}
			}
		})
	}
}
