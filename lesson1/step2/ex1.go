package step2

import (
	"io"
	"os"
)

func ReadContent(filename string) string {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	if err != nil {
		return ""
	}
	defer file.Close()

	var result []byte
	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			return ""
		}

		result = append(result, buffer[:n]...)
	}

	return string(result)
}
