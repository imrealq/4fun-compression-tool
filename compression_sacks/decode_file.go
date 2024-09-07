package compression_sacks

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func DecodeFile(inputPath, outputPath string) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("error opening input file: %v", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer outputFile.Close()

	// Đọc header
	header, err := ReadHeader(inputFile)
	if err != nil {
		return fmt.Errorf("error reading header: %v", err)
	}

	// Tạo bảng mã Huffman ngược
	decodeTable := make(map[string]rune)
	for i, symbol := range header.Symbols {
		decodeTable[header.Codes[i]] = symbol
	}

	// Đọc và giải mã dữ liệu
	buffer := make([]byte, 1)
	var currentCode string
	for {
		_, err := inputFile.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading input file: %v", err)
		}

		for i := 7; i >= 0; i-- {
			bit := (buffer[0] >> i) & 1
			currentCode += fmt.Sprintf("%d", bit)

			if symbol, ok := decodeTable[currentCode]; ok {
				outputFile.WriteString(string(symbol))
				currentCode = ""
			}
		}
	}

	return nil
}

// ReadHeader đọc thông tin header từ file đã nén
func ReadHeader(r io.Reader) (HeaderInfo, error) {
	var info HeaderInfo

	// Đọc số lượng ký tự
	if err := binary.Read(r, binary.BigEndian, &info.SymbolCount); err != nil {
		return info, fmt.Errorf("error reading symbol count: %v", err)
	}

	// Đọc từng ký tự và mã Huffman của nó
	for i := uint32(0); i < info.SymbolCount; i++ {
		var symbol uint32
		if err := binary.Read(r, binary.BigEndian, &symbol); err != nil {
			return info, fmt.Errorf("error reading symbol: %v", err)
		}
		info.Symbols = append(info.Symbols, rune(symbol))

		var code string
		for {
			var b byte
			if err := binary.Read(r, binary.BigEndian, &b); err != nil {
				return info, fmt.Errorf("error reading code byte: %v", err)
			}
			if b == 0 {
				break
			}
			code += string(b)
		}
		info.Codes = append(info.Codes, code)
	}

	return info, nil
}
