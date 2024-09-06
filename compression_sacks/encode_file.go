package compression_sacks

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

func EncodeFile(inputPath, outputPath string, codes map[rune]string) error {
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

	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	var buffer byte
	var bitsUsed int
	chunk := make([]byte, ChunkSize)

	for {
		n, err := reader.Read(chunk)
		if err != nil && err != io.EOF {
			return fmt.Errorf("error reading input file: %v", err)
		}

		// Xử lý từng byte trong chunk
		for i := 0; i < n; {
			// đoạn này xử lý những rune là 2 byte trở lên
			// lấy toàn bộ byte còn lại xem có 1 phải là 1 rune không.
			// nếu rune đó không phải lầ rune từ 2 byte trở lên thì xử lý là rune một byte
			r, size := utf8.DecodeRune(chunk[i:])
			if r == utf8.RuneError && size == 1 {
				// Xử lý byte không hợp lệ
				r = rune(chunk[i])
				size = 1
			}

			code, ok := codes[r]
			if !ok {
				return fmt.Errorf("no Huffman code found for rune: %v", r)
			}

			// khi đã có code của ký tự trong bảng mã.
			// biến code đó thành binary
			for _, bit := range code {
				if bit == '1' {
					// buffer có kiểu byte = uint8 với 8 bit

					// a << b có nghĩa là dịch chuyển các bit của a sang trái b vị trí
					// ví dụ:
					// Nếu bitsUsed = 0: 1 << (7-0) = 1 << 7 = 10000000
					// Nếu bitsUsed = 1: 1 << (7-1) = 1 << 6 = 01000000
					// Nếu bitsUsed = 7: 1 << (7-7) = 1 << 0 = 00000001

					// Toán tử |= chỉ thêm bit 1 vào các vị trí cần thiết mà không ảnh hưởng đến các bit khác
					buffer |= 1 << (7 - bitsUsed)
				}
				bitsUsed++

				if bitsUsed == 8 {
					err = writer.WriteByte(buffer)
					if err != nil {
						return fmt.Errorf("error writing to output file: %v", err)
					}
					buffer = 0
					bitsUsed = 0
				}
			}
			i += size
		}

		if err == io.EOF {
			break
		}
	}

	// Xử lý các bit còn lại
	if bitsUsed > 0 {
		err = writer.WriteByte(buffer)
		if err != nil {
			return fmt.Errorf("error writing final byte to output file: %v", err)
		}
	}

	return nil
}
