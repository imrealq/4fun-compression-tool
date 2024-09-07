package compression_sacks

import (
	"bytes"
	"encoding/binary"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestReadHeader(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		want    HeaderInfo
		wantErr bool
	}{
		{
			name: "Basic ASCII characters",
			input: func() []byte {
				buf := new(bytes.Buffer)
				binary.Write(buf, binary.BigEndian, uint32(3))
				binary.Write(buf, binary.BigEndian, uint32('A'))
				buf.Write([]byte("0\x00"))
				binary.Write(buf, binary.BigEndian, uint32('B'))
				buf.Write([]byte("10\x00"))
				binary.Write(buf, binary.BigEndian, uint32('C'))
				buf.Write([]byte("11\x00"))
				return buf.Bytes()
			}(),
			want: HeaderInfo{
				SymbolCount: 3,
				Symbols:     []rune{'A', 'B', 'C'},
				Codes:       []string{"0", "10", "11"},
			},
			wantErr: false,
		},
		{
			name: "Unicode characters",
			input: func() []byte {
				buf := new(bytes.Buffer)
				binary.Write(buf, binary.BigEndian, uint32(3))
				binary.Write(buf, binary.BigEndian, uint32('世'))
				buf.Write([]byte("0\x00"))
				binary.Write(buf, binary.BigEndian, uint32('界'))
				buf.Write([]byte("10\x00"))
				binary.Write(buf, binary.BigEndian, uint32('🌍'))
				buf.Write([]byte("11\x00"))
				return buf.Bytes()
			}(),
			want: HeaderInfo{
				SymbolCount: 3,
				Symbols:     []rune{'世', '界', '🌍'},
				Codes:       []string{"0", "10", "11"},
			},
			wantErr: false,
		},
		{
			name: "Empty input",
			input: func() []byte {
				buf := new(bytes.Buffer)
				binary.Write(buf, binary.BigEndian, uint32(0))
				return buf.Bytes()
			}(),
			want: HeaderInfo{
				SymbolCount: 0,
				Symbols:     nil,
				Codes:       nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := bytes.NewReader(tt.input)
			got, err := ReadHeader(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeFile(t *testing.T) {
	// Tạo temporary directory cho test files
	tempDir, err := os.MkdirTemp("", "huffman_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Tạo mock encoded file
	inputPath := filepath.Join(tempDir, "encoded.bin")
	outputPath := filepath.Join(tempDir, "decoded.txt")

	encodedContent := []byte{
		0, 0, 0, 3, // SymbolCount
		0, 0, 0, 65, '0', 0, // 'A' and its code
		0, 0, 0, 66, '1', '0', 0, // 'B' and its code
		0, 0, 0, 67, '1', '1', 0, // 'C' and its code
		0b00101010,
		0b11111111, // Encoded data: AABBBCCCC
	}

	if err := os.WriteFile(inputPath, encodedContent, 0644); err != nil {
		t.Fatalf("Failed to write mock encoded file: %v", err)
	}

	// Run DecodeFile
	err = DecodeFile(inputPath, outputPath)
	if err != nil {
		t.Fatalf("DecodeFile failed: %v", err)
	}

	// Read and check the decoded content
	decodedContent, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read decoded file: %v", err)
	}

	expectedContent := "AABBBCCCC"
	if string(decodedContent) != expectedContent {
		t.Errorf("Decoded content mismatch. Got %s, want %s", string(decodedContent), expectedContent)
	}
}
