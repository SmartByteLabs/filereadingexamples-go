package bytesbuffer_read

import (
	"bytes"
	"io"
)

func BytesBufferRead(input io.Reader) (string, error) {
	var buf bytes.Buffer

	if _, err := io.Copy(&buf, input); err != nil {
		return "", err
	}
	return buf.String(), nil
}
