package compression_sacks_test

import (
	"bytes"
	"compression_sacks"
	"encoding/binary"
	"reflect"
	"testing"
)

func TestCreateHeaderInfo(t *testing.T) {
	codes := map[rune]string{
		'A': "0",
		'B': "10",
		'C': "110",
		'D': "111",
	}

	expected := compression_sacks.HeaderInfo{
		SymbolCount: 4,
		Symbols:     []rune{'A', 'B', 'C', 'D'},
		Codes:       []string{"0", "10", "110", "111"},
	}

	result := compression_sacks.CreateHeaderInfo(codes)

	if result.SymbolCount != expected.SymbolCount {
		t.Errorf("Expected SymbolCount %d, got %d", expected.SymbolCount, result.SymbolCount)
	}

	if !reflect.DeepEqual(result.Symbols, expected.Symbols) {
		t.Errorf("Expected Symbols %v, got %v", expected.Symbols, result.Symbols)
	}

	if !reflect.DeepEqual(result.Codes, expected.Codes) {
		t.Errorf("Expected Codes %v, got %v", expected.Codes, result.Codes)
	}
}

func TestWriteHeader(t *testing.T) {
	info := compression_sacks.HeaderInfo{
		SymbolCount: 3,
		Symbols:     []rune{'A', 'B', 'C'},
		Codes:       []string{"0", "10", "110"},
	}

	buf := new(bytes.Buffer)

	err := compression_sacks.WriteHeader(buf, info)
	if err != nil {
		t.Fatalf("WriteHeader returned an error: %v", err)
	}

	// Verify the written data
	var symbolCount uint32
	err = binary.Read(buf, binary.BigEndian, &symbolCount)
	if err != nil {
		t.Fatalf("Failed to read SymbolCount: %v", err)
	}
	if symbolCount != info.SymbolCount {
		t.Errorf("Expected SymbolCount %d, got %d", info.SymbolCount, symbolCount)
	}

	for i, expectedSymbol := range info.Symbols {
		var symbol uint32
		err = binary.Read(buf, binary.BigEndian, &symbol)
		if err != nil {
			t.Fatalf("Failed to read symbol: %v", err)
		}
		if rune(symbol) != expectedSymbol {
			t.Errorf("Expected symbol %c, got %c", expectedSymbol, rune(symbol))
		}

		code := make([]byte, uint8(len(info.Codes[i])))
		_, err = buf.Read(code)
		if err != nil {
			t.Fatalf("Failed to read code: %v", err)
		}
		if string(code) != info.Codes[i] {
			t.Errorf("Expected code %s, got %s", info.Codes[i], string(code))
		}
	}

	if buf.Len() != 0 {
		t.Errorf("Buffer not fully read, %d bytes remaining", buf.Len())
	}
}
