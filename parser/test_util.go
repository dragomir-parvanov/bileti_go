package parser

import (
	"bileti_go/utils"
	"io"
	"os"
)

func GetTestFile(filePath string) io.Reader {
	file := utils.Must(os.Open(filePath))

	return file
}
