package io_read

import (
	"bytes"
	"io"
)

func IoRead(input io.Reader) (string, error) {
	buf := make([]byte, 1024) // Create a buffer
	var output bytes.Buffer   // Create a bytes.Buffer to store the output

	for {
		n, err := input.Read(buf) // Read data into the buffer

		if err == io.EOF {
			break // End of file, break the loop
		}
		if err != nil {
			return "", err // Return the error if it's not EOF
		}

		// Write the read data from the buffer into the output
		output.Write(buf[:n])
	}

	return output.String(), nil // Return the accumulated output
}
