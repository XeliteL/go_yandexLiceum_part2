package step5

import (
	"bytes"
	"context"
	"errors"
	"io"
)

func Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error) {
	if len(seq) == 0 {
		return false, errors.New("")
	}

	chunk := make([]byte, 1024)
	window := make([]byte, 0, len(seq)+1024)

	for {
		select {
		case <-ctx.Done():
			return false, ctx.Err()
		default:
		}

		n, err := r.Read(chunk)
		if n > 0 {
			window = append(window, chunk[:n]...)

			if bytes.Contains(window, seq) {
				return true, nil
			}

			if len(window) > len(seq)-1 {
				window = window[len(window)-(len(seq)-1):]
			}
		}

		if err != nil {
			if err == io.EOF {
				return false, nil
			}
			return false, err
		}
	}
}
