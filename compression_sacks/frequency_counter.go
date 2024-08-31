package compression_sacks

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const ChunkSize = 32 * 1024 // 32 KB

// CountFrequencies để có thể export, dùng trong file khác cùng package
func CountFrequencies(filePath string) (map[byte]int, error) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("Error open file: %v\n", err)
		return nil, err
	}

	defer file.Close()

	frequencies := make(map[byte]int)

	// đối với file lớn. không nên đọc toàn bộ file 1 lúc
	// chia nhỏ thành ChunkSize
	reader := bufio.NewReaderSize(file, ChunkSize)
	buffer := make([]byte, ChunkSize)

	// for forever
	for {
		// n số lượng byte đã đọc vào buffer
		// nếu buffer là "ABACA" -> n = 5 (byte)
		n, err := reader.Read(buffer)

		// TODO: khi ký tự hơn 1 byte? Ví dụ "世" = 3 byte

		if err != nil && err != io.EOF {
			fmt.Printf("Error read file: %v\n", err)
			return nil, err
		}

		// empty file
		if n == 0 {
			break
		}

		for i := 0; i < n; i++ {
			frequencies[buffer[i]]++
		}

		// khi read hết file
		if err == io.EOF {
			break
		}
	}
	return frequencies, nil
}
