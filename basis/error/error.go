package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func readFile(path string) ([]byte, error) {
	parentPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	fullPath := filepath.Join(parentPath, path)
	file, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return read(file)
}

func main() {
	path := "error.go"
	ba, err := readFile(path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("The content of '%s':\n%s\n", path, ba)
}

func read(r io.Reader) ([]byte, error) {
	br := bufio.NewReader(r)
	var buf bytes.Buffer
	for {
		ba, isPrefix, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		buf.Write(ba)
		if !isPrefix {
			buf.WriteByte('\n')
		}
	}
	return buf.Bytes(), nil
}
