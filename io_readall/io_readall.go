package io_readall

import (
	"io"
	"log"
)

func IoReadAll(input io.Reader) (string, error) {
	readData, err := io.ReadAll(input)
	if err != nil {
		log.Fatal(err)
		return "", err // Return empty string and error in case of error
	}
	return string(readData), nil 
}
