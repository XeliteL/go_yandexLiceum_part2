package step1

import (
	"io"
)

func Copy(r io.Reader, w io.Writer, n uint) error {
	buf := make([]byte, n)

	bytesRead, err := r.Read(buf)

	if bytesRead > 0 {
		_, writeErr := w.Write(buf[:bytesRead])
		if writeErr != nil {
			return nil
		}
	}

	if err != nil && err != io.EOF {
		return err
	}

	return nil
}
