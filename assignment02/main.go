package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type logWriter struct{}

func main() {
	filename := os.Args[1]

	fmt.Println("File name:", filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	lw := logWriter{}

	io.Copy(lw, file)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
