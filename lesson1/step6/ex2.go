package step6

import (
	"sync"
)

func Average(names []string) (int, error) {
	if len(names) == 0 {
		return 0, nil
	}

	total := 0
	var mu sync.Mutex
	var firstError error
	wg := &sync.WaitGroup{}

	wg.Add(len(names))
	for _, name := range names {
		go func(studentName string) {
			defer wg.Done()

			grade, err := getGrade(studentName)

			mu.Lock()
			defer mu.Unlock()

			if firstError != nil {
				return
			}
			if err != nil {
				firstError = err
				return
			}

			total += grade
		}(name)
	}
	wg.Wait()

	mu.Lock()
	defer mu.Unlock()

	if firstError != nil {
		return 0, firstError
	}

	return total / len(names), nil
}
