package main

import (
	"fmt"
	"io"
	"os"
	// "strings"
)

func main() {

	// read
	// r := strings.NewReader("Hello World!")		// read from a string
	file, fileOpenErr := os.Open("readFile.txt") // read from a file
	if fileOpenErr != nil {
		fmt.Println("failed to open file:", fileOpenErr)
		return
	}
	defer file.Close()

	count, err := countBytes(file)

	if err != nil {
		fmt.Println("error occured while counting:", err)
	}
	fmt.Println("number of bytes from the string:", count)

	// write
	wrtFile, fileCreateErr := os.Create("writeFile.txt")
	if fileCreateErr != nil {
		fmt.Println("error while creating a file:", fileCreateErr)
	}
	defer file.Close()
	msg, writeErr := writeFile(wrtFile,"Hi Anil!!")
	if writeErr != nil {
		fmt.Println("error writing file:", writeErr)
		return
	}

	fmt.Println("file written:", msg)
}

func countBytes(r io.Reader) (int, error) {
	buf := make([]byte, 1024)
	total := 0

	for {
		n, err := r.Read(buf)
		total += n

		if err == io.EOF {
			return total, nil
		}

		if err != nil {
			return total, err
		}
	}
}

func writeFile(w io.Writer, data string) (string, error) {
	if _, writeErr := w.Write([]byte(data)); writeErr != nil {
		return "failed", writeErr
	}

	return "Success", nil
}
