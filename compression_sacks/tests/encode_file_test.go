package compression_sacks_test

import (
	"bytes"
	"compression_sacks"
	"os"
	"testing"
)

func TestEncodeFile(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		codes       map[rune]string
		expected    []byte
		expectError bool
	}{
		{
			name:  "Basic encoding",
			input: "AABBBCCCC",
			codes: map[rune]string{
				'A': "0",
				'B': "10",
				'C': "11",
			},
			expected:    []byte{0b00101010, 0b11111111},
			expectError: false,
		},
		{
			name:        "Empty input",
			input:       "",
			codes:       map[rune]string{},
			expected:    []byte{},
			expectError: false,
		},
		{
			name:  "Unicode characters",
			input: "こんにちは",
			codes: map[rune]string{
				'こ': "00",
				'ん': "01",
				'に': "10",
				'ち': "110",
				'は': "111",
			},
			expected:    []byte{0b00011011, 0b01110000},
			expectError: false,
		},
		{
			name:  "Missing code",
			input: "ABC",
			codes: map[rune]string{
				'A': "0",
				'B': "1",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary input and output files
			inputFile, err := os.CreateTemp("", "input")
			if err != nil {
				t.Fatalf("Failed to create temp input file: %v", err)
			}
			defer os.Remove(inputFile.Name())

			outputFile, err := os.CreateTemp("", "output")
			if err != nil {
				t.Fatalf("Failed to create temp output file: %v", err)
			}
			defer os.Remove(outputFile.Name())

			// Write input to file
			_, err = inputFile.Write([]byte(tt.input))
			if err != nil {
				t.Fatalf("Failed to write to input file: %v", err)
			}
			inputFile.Close()

			// Call EncodeFile
			err = compression_sacks.EncodeFile(inputFile.Name(), outputFile.Name(), tt.codes)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected an error, but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			// Read output file
			output, err := os.ReadFile(outputFile.Name())
			if err != nil {
				t.Fatalf("Failed to read output file: %v", err)
			}

			// Compare output with expected
			if !bytes.Equal(output, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, output)
			}
		})
	}
}
