package compression_sacks_test

import (
	"os"
	"reflect"
	"testing"

	"compression_sacks"
)

func TestCountFrequencies(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected map[byte]int
	}{
		{
			name:     "Empty file",
			content:  "",
			expected: map[byte]int{},
		},
		{
			name:    "Simple ASCII content",
			content: "ABACABA",
			expected: map[byte]int{
				'A': 4,
				'B': 2,
				'C': 1,
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

			if _, err := tmpfile.Write([]byte(tt.content)); err != nil {
				t.Fatalf("Could not write to temporary file: %v", err)
			}
			if err := tmpfile.Close(); err != nil {
				t.Fatalf("Could not close temporary file: %v", err)
			}

			got, err := compression_sacks.CountFrequencies(tmpfile.Name())
			if err != nil {
				t.Fatalf("compression_sacks.CountFrequencies() error = %v", err)
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("compression_sacks.CountFrequencies() = %v, want %v", got, tt.expected)
			}
		})
	}
}
