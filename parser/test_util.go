package parser

import (
	"io"
	"log"
	"os"
)

func GetTestFile(filePath string) io.Reader {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
