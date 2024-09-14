package bufio_read

import (
	"bufio"
	"io"
	"strings"
)

// BufioRead reads input from an io.Reader and returns the content as a string.
func BufioRead(input io.Reader) (string, error) {
	var output strings.Builder

	scanner := bufio.NewReader(input)

	for {
		read, err := scanner.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break // EOF, exit loop
			}
			return "", err // Return err
		}
		output.WriteString(read)
	}
	return output.String(), nil
}
