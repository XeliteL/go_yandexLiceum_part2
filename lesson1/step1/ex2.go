package step1

import (
	"io"
)

func ReadString(r io.Reader) (string, error) {
	var result []byte
	data := make([]byte, 1024)

	for {
		bytesRead, err := r.Read(data)
		if bytesRead > 0 {
			result = append(result, data[:bytesRead]...)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
	}

	return string(result), nil
}
