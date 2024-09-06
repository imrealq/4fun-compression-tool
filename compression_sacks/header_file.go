package compression_sacks

import (
	"encoding/binary"
	"io"
)

// HeaderInfo lưu trữ thông tin cần thiết cho header
type HeaderInfo struct {
	SymbolCount uint32
	Symbols     []rune
	Codes       []string
}

// WriteHeader ghi thông tin header vào writer
func WriteHeader(w io.Writer, info HeaderInfo) error {
	// Ghi số lượng ký tự
	if err := binary.Write(w, binary.BigEndian, info.SymbolCount); err != nil {
		return err
	}

	// Ghi từng ký tự và mã Huffman của nó
	for i, symbol := range info.Symbols {
		if err := binary.Write(w, binary.BigEndian, uint32(symbol)); err != nil {
			return err
		}

		// Ghi mã Huffman
		if _, err := w.Write([]byte(info.Codes[i])); err != nil {
			return err
		}
	}

	return nil
}

// CreateHeaderInfo tạo HeaderInfo từ bảng mã Huffman
func CreateHeaderInfo(codes map[rune]string) HeaderInfo {
	info := HeaderInfo{
		SymbolCount: uint32(len(codes)),
		Symbols:     make([]rune, 0, len(codes)),
		Codes:       make([]string, 0, len(codes)),
	}

	for symbol, code := range codes {
		info.Symbols = append(info.Symbols, symbol)
		info.Codes = append(info.Codes, code)
	}

	return info
}
