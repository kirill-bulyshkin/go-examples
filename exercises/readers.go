package exercises

import (
	"log"
	"os"

	"golang.org/x/tour/reader"
)

// Exercise: Readers

// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

// TODO: Add a Read([]byte) (int, error) method to MyReader.

type MyReader struct{}

func (MyReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 'A'
	}
	return len(p), nil
}

func TestRead() {
	myReader := MyReader{}

	reader.Validate(myReader)

	buf := make([]byte, 20)
	_, err := myReader.Read(buf)

	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(buf)
}
