package step8

import (
	"context"
	"os"
)

func readJSON(ctx context.Context, path string, result chan<- []byte) {
	done := make(chan []byte, 1)
	errCh := make(chan error, 1)

	go func() {
		data, err := os.ReadFile(path)
		if err != nil {
			errCh <- err
			return
		}
		done <- data
	}()

	select {
	case data := <-done:
		select {
		case result <- data:
		default:
		}
	case <-errCh:
		return
	case <-ctx.Done():
		return
	}
}
