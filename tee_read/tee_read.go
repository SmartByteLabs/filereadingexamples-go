package tee_read

import (
	"bytes"
	"io"
)

func TeeRead(input io.Reader) (string, error) {
	// Create byte buffer to store logs
	var logs bytes.Buffer

	// Create a teeReader that read from input and writes to logs
	teeReader := io.TeeReader(input, &logs)

	// Use buffer to read data in chunks
	buffer := make([]byte, 256)
	for {
		_, err := teeReader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

	}
	return logs.String(), nil
}
