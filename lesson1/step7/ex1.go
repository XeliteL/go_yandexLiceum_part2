package step7

import (
	"errors"
	"time"
)

func TimeoutFibonacci(n int, timeout time.Duration) (int, error) {
	result := make(chan int, 1)

	switch {
	case n < 0:
		return 0, errors.New("n must be non-negative")
	case n == 0:
		result <- 0
	case n == 1:
		result <- 1
	default:
		go func(h int) {
			fib := []int{0, 1}
			for len(fib) <= n {
				fib = append(fib, fib[len(fib)-1]+fib[len(fib)-2])
			}

			result <- fib[n]
		}(n)
	}

	select {
	case res := <-result:
		return res, nil
	case <-time.After(timeout):
		return 0, errors.New("timeout")
	}
}
