package step8

import (
	"math"
	"time"
)

func GeneratePrimeNumbers(stop chan struct{}, prime_nums chan int, N int) {
	go func() {
		defer close(prime_nums)

		if N > 2 {
			select {
			case prime_nums <- 2:
			case <-stop:
				return
			}
		}

		for n := 3; n < N; n += 2 {
			isPrime := true
			limit := int(math.Sqrt(float64(n))) + 1
			for i := 3; i < limit; i++ {
				if n%i == 0 {
					isPrime = false
					break
				}
			}

			if isPrime {
				select {
				case prime_nums <- n:
				case <-stop:
					return
				}
			}

			select {
			case <-stop:
				return
			default:
			}
		}
	}()

	time.AfterFunc(time.Millisecond*100, func() {
		close(stop)
	})
}
