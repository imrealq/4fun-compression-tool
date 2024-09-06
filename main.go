package main

import (
	"compression_sacks/compression_sacks"
	"fmt"
	"path/filepath"
)

func main() {
	inputPath := filepath.Join("demo", "135-0.txt")
	outputPath := filepath.Join("demo", "135-0.compressed")

	freqMap, err := compression_sacks.CountFrequencies(inputPath)

	if err != nil {
		fmt.Printf("Error counting frequencies")
		return
	}

	root := compression_sacks.BuildHuffmanTree(freqMap)

	codes := compression_sacks.GenerateHuffmanCodes(root)

	compression_sacks.EncodeFile(inputPath, outputPath, codes)
}
