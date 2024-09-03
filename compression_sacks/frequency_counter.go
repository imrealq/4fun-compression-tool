package compression_sacks

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

const ChunkSize = 32 * 1024 // 32 KB

// CountFrequencies để có thể export, dùng trong file khác cùng package
func CountFrequencies(filePath string) (map[rune]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error open file: %v\n", err)
		return nil, err
	}
	defer file.Close()

	frequencies := make(map[rune]int)
	// đối với file lớn. không nên đọc toàn bộ file 1 lúc
	// chia nhỏ thành ChunkSize
	reader := bufio.NewReaderSize(file, ChunkSize)

	// for forever
	for {
		r, size, err := reader.ReadRune()
		if err != nil {
			// khi read hết file
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading rune: %v\n", err)
			return nil, err
		}

		if r == utf8.RuneError && size == 1 {
			fmt.Println("Warning: invalid UTF-8 encoding encountered")
			continue
		}

		frequencies[r]++
	}

	return frequencies, nil
}
