package io_read

import (
	"bytes"
	"io"
	"testing"

	"github.com/0xSumeet/filereadingexamples-go/testfiles"
)

const (
	case1 = "Benchmark Test Case 1"
	case2 = "Benchmark Test Case 2"
	case3 = "Benchmark Test Case 3"
)

var (
	result string
	s      string
)

// TestIoRead tests the IoRead function
func TestIoRead(t *testing.T) {
	tests := []struct {
		name      string
		filename  string
		expectErr bool
	}{
		{
			name:      "Test File 1",
			filename:  "words_1.txt",
			expectErr: false,
		},
		{
			name:      "Test File 2",
			filename:  "words_2.txt",
			expectErr: false,
		},
		{
			name:      "Test File 3",
			filename:  "words_3.txt",
			expectErr: false,
		},
		{
			name:      "Test File 4",
			filename:  "words_4.txt",
			expectErr: false,
		},
		{
			name:      "Test File 5",
			filename:  "words_5.txt",
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the embedded test file
			fileContent, err := testfiles.TestFiles.ReadFile(tt.filename)
			if err != nil {
				t.Fatalf("Failed to read embedded file %s: %v", tt.filename, err)
			}

			// Create a reader for the file content
			reader := bytes.NewReader(fileContent)

			// Call the IoRead function
			result, err := IoRead(reader)

			if tt.expectErr && err == nil {
				t.Errorf("Expected error but got nil")
			}

			// Compare the result with the content of the test file
			if string(fileContent) != result {
				t.Errorf("Expected file content %q, but got %q", string(fileContent), result)
			}
		})
	}
}

// BenchMarks (Benchmark Test)
func BenchmarkTestCase1(b *testing.B) {
	fileName := "words_1.txt"
	fileContent, err := testfiles.TestFiles.Open(fileName)
	if err != nil {
		b.Fatalf("Failed to read embedded file %s: %v", fileName, err)
	}
	defer fileContent.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Run the IoRead Function b.N times
		s, err = IoRead(fileContent)
		if err != nil {
			b.Fatalf("IoRead Failed for %s: %v", case1, err)
		}

		// Store the result to prevent the compiler optimizations
		result = s

		// Reset the file pointer
		if _, err := fileContent.(io.Seeker).Seek(0, io.SeekStart); err != nil {
			b.Fatalf("Failed to reset %s file pointer: %v", fileName, err)
		}
	}
}

func BenchmarkTestCase2(b *testing.B) {
	fileName := "words_2.txt"
	fileContent, err := testfiles.TestFiles.Open(fileName)
	if err != nil {
		b.Fatalf("Failed to read embedded file %s: %v", fileName, err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s, err = IoRead(fileContent)
		if err != nil {
			b.Fatalf("IoRead Failed for %s: %v", case2, err)
		}
		result = s

		if _, err := fileContent.(io.Seeker).Seek(0, io.SeekStart); err != nil {
			b.Fatalf("Failed to reset %s file pointer: %v", fileName, err)
		}
	}
}

func BenchmarkTestCase3(b *testing.B) {
	fileName := "words_3.txt"
	fileContent, err := testfiles.TestFiles.Open(fileName)
	if err != nil {
		b.Fatalf("Failed to read embedded file %s: %v", fileName, err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s, err = IoRead(fileContent)
		if err != nil {
			b.Fatalf("IoRead Failed for %s: %v", case3, err)
		}
		result = s

		if _, err := fileContent.(io.Seeker).Seek(0, io.SeekStart); err != nil {
      b.Fatalf("Failed to reset %s file pointer: %v", fileName, err)
		}
	}
}
