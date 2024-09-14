package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func readWordsFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

func writeRandomWordsToFile(outputPath string, sizeBytes int64, words []string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	var writtenBytes int64

	for writtenBytes < sizeBytes {
		word := words[rand.Intn(len(words))]
		line := word + "\n"
		n, err := writer.WriteString(line)
		if err != nil {
			return err
		}
		writtenBytes += int64(n)
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <size> <MB|GB>")
		os.Exit(1)
	}

	sizeStr := os.Args[1]
	unit := os.Args[2]

	var sizeBytes int64
	var err error

	switch unit {
	case "MB":
		sizeMB, err := strconv.Atoi(sizeStr)
		if err != nil {
			fmt.Printf("Invalid size: %v\n", err)
			os.Exit(1)
		}
		sizeBytes = int64(sizeMB) * 1024 * 1024
	case "GB":
		sizeGB, err := strconv.Atoi(sizeStr)
		if err != nil {
			fmt.Printf("Invalid size: %v\n", err)
			os.Exit(1)
		}
		sizeBytes = int64(sizeGB) * 1024 * 1024 * 1024
	default:
		fmt.Println("Unit must be MB or GB")
		os.Exit(1)
	}

	// Add your source destination
	dictFile := "/usr/share/dict/words/one.txt"
	outputFile := fmt.Sprintf("random_words_%s_%s.txt", sizeStr, unit)

	words, err := readWordsFromFile(dictFile)
	if err != nil {
		fmt.Printf("Error reading dictionary file: %v\n", err)
		os.Exit(1)
	}

	if len(words) == 0 {
		fmt.Println("Dictionary file is empty")
		os.Exit(1)
	}

	if err := writeRandomWordsToFile(outputFile, sizeBytes, words); err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated file of size %d bytes, saved to %s\n", sizeBytes, outputFile)
}
