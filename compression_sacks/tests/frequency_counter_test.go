package compression_sacks_test

import (
	"compression_sacks"
	"os"
	"reflect"
	"testing"
)

func TestCountFrequencies(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected map[rune]int
	}{
		{
			name:     "Empty file",
			content:  "",
			expected: map[rune]int{},
		},
		{
			name:    "Simple ASCII content",
			content: "ABACABA",
			expected: map[rune]int{
				'A': 4,
				'B': 2,
				'C': 1,
			},
		},
		{
			name:    "Unicode content",
			content: "Hello, 世界！",
			expected: map[rune]int{
				'H': 1,
				'e': 1,
				'l': 2,
				'o': 1,
				',': 1,
				' ': 1,
				'世': 1,
				'界': 1,
				'！': 1,
			},
		},
		{
			name:    "Mixed content",
			content: "АБВ abc 123 あいう",
			expected: map[rune]int{
				'А': 1, 'Б': 1, 'В': 1,
				'a': 1, 'b': 1, 'c': 1,
				'1': 1, '2': 1, '3': 1,
				' ': 3,
				'あ': 1, 'い': 1, 'う': 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpfile, err := os.CreateTemp("", "test")
			if err != nil {
				t.Fatalf("Could not create temporary file: %v", err)
			}
			defer os.Remove(tmpfile.Name())

			if _, err := tmpfile.WriteString(tt.content); err != nil {
				t.Fatalf("Could not write to temporary file: %v", err)
			}
			if err := tmpfile.Close(); err != nil {
				t.Fatalf("Could not close temporary file: %v", err)
			}

			got, err := compression_sacks.CountFrequencies(tmpfile.Name())
			if err != nil {
				t.Fatalf("compression.CountFrequencies() error = %v", err)
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("compression.CountFrequencies() = %v, want %v", got, tt.expected)
			}
		})
	}
}
